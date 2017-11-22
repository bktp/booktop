package main

import (
	"log"
	"net/http"

	"github.com/bt/search"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/bt/books"
	"github.com/bt/pages"
	"github.com/bt/users"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/book/search", users.CheckAuth(search.Search)).Queries("pattern", "").Methods("GET")

	router.HandleFunc("/api/book/{isbn}/pages/{pagenum}", users.CheckAuth(pages.GetPage)).Methods("GET")
	router.HandleFunc("/api/book/{isbn}/pages/{pagenum}", users.CheckAuth(pages.UpdatePage)).Methods("PUT")
	router.HandleFunc("/api/book/{isbn}/pages", users.CheckAuth(pages.GetPages)).Methods("GET")
	router.HandleFunc("/api/book/{isbn}/pages", users.CheckAuth(pages.AddPage)).Methods("POST")

	router.HandleFunc("/api/book/{isbn}", users.CheckAuth(books.GetBook)).Methods("GET")
	router.HandleFunc("/api/book/{isbn}", users.CheckAuth(books.DeleteBook)).Methods("DELETE")
	router.HandleFunc("/api/book/{isbn}", users.CheckAuth(books.UpdateBook)).Methods("PUT")
	router.HandleFunc("/api/book", users.CheckAuth(books.AddBook)).Methods("POST")

	router.HandleFunc("/api/auth", users.Auth).Methods("POST")
	router.HandleFunc("/api/users", users.AddUser).Methods("POST")
	router.HandleFunc("/api/user/role", users.CheckAuth(users.GetUserRole)).Methods("GET")

	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./public/images"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/www")))

	log.Println("Server is starting")
	http.ListenAndServe(":8080", handlers.CORS()(router))
}
