package books

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bktp/booktop/server/src/config"

	_ "github.com/lib/pq"
)

type Book struct {
	ISBN         string   `json:"isbn"`
	Name         string   `json:"name"`
	OriginalName string   `json:"original,omitempty"`
	Authors      []string `json:"authors,omitempty"`
	Published    string   `json:"published,omitempty"`
	Category     string   `json:"category,omitempty"`
	Description  string   `json:"description,omitempty"`
	Cover        string   `json:"cover"`
}

func fetchCategoryID(category string) (int, error) {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var categoryID int
	err = db.QueryRow("SELECT id FROM categories WHERE name = $1", category).Scan(&categoryID)
	if err != nil {
		return 0, err
	}

	return categoryID, nil
}

func fetchBook(isbn string) (*Book, error) {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var book Book
	var categoryID int
	err = db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn).Scan(&book.ISBN, &book.Name,
		&book.OriginalName, &book.Published, &book.Description, &book.Cover, &categoryID)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT name FROM categories WHERE id = $1", categoryID).Scan(&book.Category)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT name FROM authors WHERE isbn = $1", isbn)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var author string
		err = rows.Scan(&author)
		if err != nil {
			return nil, err
		}
		book.Authors = append(book.Authors, author)
	}

	return &book, nil
}

func addBook(book *Book) error {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	categoryID, err := fetchCategoryID(book.Category)
	_, err = tx.Exec("INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7)", book.ISBN, book.Name, book.OriginalName,
		book.Published, book.Description, book.Cover, categoryID)
	if err != nil {
		return err
	}

	for _, author := range book.Authors {
		_, err = tx.Exec("INSERT INTO authors VALUES ($1, $2)", book.ISBN, author)
		if err != nil {
			return nil
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func updateBook(book *Book, isbn string) error {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	categoryID, err := fetchCategoryID(book.Category)
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE books SET isbn=$1, name=$2, original=$3, published=$4, description=$5, cover=$6, category_id=$7 WHERE isbn=$8",
		book.ISBN, book.Name, book.OriginalName, book.Published, book.Description, book.Cover, categoryID, isbn)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM authors WHERE isbn=$1", book.ISBN)
	if err != nil {
		return err
	}
	for _, author := range book.Authors {
		_, err = tx.Exec("INSERT INTO authors VALUES ($1, $2)", book.ISBN, author)
		if err != nil {
			return nil
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func deleteBook(isbn string) error {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM books WHERE isbn=$1", isbn)
	if err != nil {
		return err
	}

	return nil
}

// GetBook returns slice of books converted to JSON
func GetBook(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	book, err := fetchBook(isbn)
	if err != nil {
		config.SendError(w, http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	response, err := json.Marshal(book)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write(response)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.SendError(w, http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	var book Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	err = addBook(&book)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("Book has been added"))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.SendError(w, http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	var book Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	isbn := mux.Vars(r)["isbn"]
	err = updateBook(&book, isbn)
	if err != nil {
		config.SendError(w, http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("The book has been updated"))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	err := deleteBook(isbn)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Write([]byte(isbn + " успешно удалена"))
}
