package pages

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/bktp/booktop/server/src/config"
)

type Page struct {
	ISBN    string `json:"isbn"`
	Pagenum int    `json:"pagenum"`
	Text    string `json:"text"`
}

func getPage(isbn string, pagenum int) (string, error) {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return "", err
	}
	defer db.Close()

	var page string
	err = db.QueryRow("SELECT text FROM pages WHERE isbn=$1 AND pagenum=$2", isbn,
		pagenum).Scan(&page)
	if err != nil {
		return "", err
	}
	return page, nil
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pagenum, err := strconv.Atoi(vars["pagenum"])
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	page := &Page{
		ISBN:    vars["isbn"],
		Pagenum: pagenum,
	}
	page.Text, err = getPage(page.ISBN, page.Pagenum)
	if err != nil {
		config.SendError(w, http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	response, err := json.Marshal(page)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Write(response)
}

func AddPage(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	var page Page
	page.ISBN = isbn
	err = json.Unmarshal(body, &page)
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

	_, err = db.Exec("INSERT INTO pages(isbn, pagenum, text) VALUES($1,$2,$3)", page.ISBN, page.Pagenum, page.Text)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("Page has been added"))
}

func UpdatePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pagenum, err := strconv.Atoi(vars["pagenum"])
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	page := &Page{
		ISBN:    vars["isbn"],
		Pagenum: pagenum,
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	err = json.Unmarshal(body, &page)
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

	_, err = db.Exec("UPDATE pages SET text=$1 WHERE isbn=$2 AND pagenum=$3", page.Text, page.ISBN, page.Pagenum)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write([]byte("Page has been updated"))
}

func GetPages(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM pages WHERE isbn=$1", isbn).Scan(&count)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	response, err := json.Marshal(count)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write(response)
}
