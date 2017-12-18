package search

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/bktp/booktop/src/books"
	"github.com/bktp/booktop/src/config"

	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("pattern")

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	var resultBookSet []books.Book

	rows, err := db.Query(`SELECT isbn, name, cover FROM books WHERE 
		(make_tsvector(name) @@ plainto_tsquery('russian', $1)) OR
		(make_tsvector(original) @@ plainto_tsquery($1))`, pattern)
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

		resultBookSet = append(resultBookSet, newBook)
	}

	rows, err = db.Query(`SELECT isbn, name, cover FROM books WHERE isbn IN 
		(SELECT isbn FROM authors WHERE make_tsvector(name) @@ plainto_tsquery('russian', $1))`, pattern)
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

		resultBookSet = append(resultBookSet, newBook)
	}

	response, err := json.Marshal(resultBookSet)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write(response)
}

func SearchPopular(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	var resultSet []books.Book

	rows, err := db.Query(`SELECT isbn, name, cover FROM books WHERE isbn in (
		SELECT isbn FROM favs GROUP BY isbn ORDER BY count(*) DESC LIMIT 5)`)
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

func SearchByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	var resultSet []books.Book
	rows, err := db.Query("SELECT isbn, name, cover FROM books WHERE category_id = (SELECT id FROM categories WHERE name = $1)", category)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	for rows.Next() {
		var book books.Book
		err = rows.Scan(&book.ISBN, &book.Name, &book.Cover)
		if err != nil {
			config.SendError(w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		innerRows, err := db.Query("SELECT name FROM authors WHERE isbn = $1", book.ISBN)
		if err != nil {
			config.SendError(w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		var authors []string
		for innerRows.Next() {
			var author string
			err = innerRows.Scan(&author)
			if err != nil {
				config.SendError(w, http.StatusInternalServerError)
				log.Panicln(err.Error())
				return
			}

			authors = append(authors, author)
		}
		book.Authors = authors
		resultSet = append(resultSet, book)
	}

	response, err := json.Marshal(resultSet)
	if err != nil {
		config.SendError(w, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write(response)
}
