package categories

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"

	"github.com/bktp/booktop/server/src/config"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	var result []Category
	rows, err := db.Query("SELECT * FROM categories WHERE id > 0")
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			config.SendError(w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		result = append(result, category)
	}

	response, err := json.Marshal(result)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write(response)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	var cat Category
	err = json.Unmarshal(body, &cat)
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

	_, err = db.Exec("UPDATE categories SET name=$1 WHERE id=$2", cat.Name, cat.ID)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("success"))
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	catID := mux.Vars(r)["id"]

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM categories WHERE id=$1", catID)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("success"))
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO categories(name) VALUES($1)", name)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("success"))
}
