package users

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bktp/booktop/src/books"
	"github.com/bktp/booktop/src/config"

	"github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
	Role     string `json:"role,omitempty"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	var user User
	err = json.Unmarshal(body, &user)

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	rawToken := make([]byte, 16)
	rand.Read(rawToken)
	user.Token = fmt.Sprintf("%x", rawToken)

	tx, err := db.Begin()
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	_, err = tx.Exec("INSERT INTO users(name, password, token) values($1,$2,$3)", user.Name, user.Password, user.Token)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	_, err = tx.Exec("INSERT INTO user_roles(token) VALUES ($1)", user.Token)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	err = tx.Commit()
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("User has been added"))
}

func GetUserRole(w http.ResponseWriter, r *http.Request) {
	var user User
	token, err := r.Cookie("token")
	if err != nil {
		config.SendError(w, http.StatusUnauthorized)
		log.Println(err.Error())
		return
	}
	user.Token = token.Value

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	err = db.QueryRow("SELECT role FROM user_roles WHERE token=$1", user.Token).Scan(&user.Role)
	if err != nil {
		config.SendError(w, http.StatusUnauthorized)
		log.Println(err.Error())
		return
	}

	user.Token = ""
	response, err := json.Marshal(user)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Write(response)
}

func getToken(user *User) error {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow("SELECT token FROM users WHERE name=$1 AND password=$2", user.Name, user.Password).Scan(&user.Token)
	if err != nil {
		return err
	}

	return nil
}

func Auth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	log.Printf("%v\n", user)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	err = getToken(&user)
	if err != nil {
		config.SendError(w, http.StatusUnauthorized)
		log.Println(err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: user.Token,
	})
	w.Write([]byte("Authorized"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
	w.Write(nil)
}

func CheckAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			config.SendError(w, http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		db, err := sql.Open("postgres", config.ConnectionString())
		if err != nil {
			config.SendError(w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}
		defer db.Close()

		_, err = db.Query("SELECT * FROM users WHERE token=$1", token.Value)
		if err == sql.ErrNoRows {
			config.SendError(w, http.StatusUnauthorized)
			log.Println(err.Error())
			return
		} else if err != nil {
			config.SendError(w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		handler(w, r)
	}
}

func AddToFavs(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	userToken, err := r.Cookie("token")
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO favs VALUES($1, $2)", userToken.Value, isbn)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("added"))
}

func DeleteFromFavs(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	userToken, err := r.Cookie("token")
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM favs WHERE user_token=$1 AND isbn=$2", userToken.Value, isbn)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("deleted"))
}

func CheckBook(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	userToken, err := r.Cookie("token")
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT isbn FROM favs WHERE user_token=$1 AND isbn=$2", userToken.Value, isbn).Scan(&result)
	if err != nil {
		config.SendError(w, http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("success"))
}

func GetFavs(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	var resultSet []books.Book

	rows, err := db.Query(`SELECT isbn, name, cover FROM books WHERE isbn in (
			SELECT isbn FROM favs WHERE user_token=$1)`, token.Value)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	for rows.Next() {
		var newBook books.Book
		err = rows.Scan(&newBook.ISBN, &newBook.Name, &newBook.Cover)
		if err != nil {
			config.SendError(w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		innerRows, err := db.Query("SELECT name FROM authors WHERE isbn=$1", newBook.ISBN)
		if err != nil {
			config.SendError(w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}
		for innerRows.Next() {
			var author string
			err = innerRows.Scan(&author)
			if err != nil {
				config.SendError(w, http.StatusInternalServerError)
				log.Println(err.Error())
				return
			}
			newBook.Authors = append(newBook.Authors, author)
		}

		resultSet = append(resultSet, newBook)
	}

	response, err := json.Marshal(resultSet)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write(response)
}
