package main

import (
	"log"
	"net/http"

	"github.com/bktp/booktop/src/search"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/bktp/booktop/src/books"
	"github.com/bktp/booktop/src/categories"
	"github.com/bktp/booktop/src/pages"
	"github.com/bktp/booktop/src/users"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/book/search/popular", search.SearchPopular).Methods("GET")
	router.HandleFunc("/api/book/search", search.Search).Queries("pattern", "").Methods("GET")

	router.HandleFunc("/api/categories", categories.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories", users.CheckAuth(categories.UpdateCategory)).Methods("PUT")
	router.HandleFunc("/api/categories", users.CheckAuth(categories.AddCategory)).Queries("name", "").Methods("POST")
	router.HandleFunc("/api/categories/{id}", users.CheckAuth(categories.DeleteCategory)).Methods("DELETE")

	router.HandleFunc("/api/category", search.SearchByCategory).Queries("category", "").Methods("GET")

	router.HandleFunc("/api/book/{isbn}/pages/{pagenum}", pages.GetPage).Methods("GET")
	router.HandleFunc("/api/book/{isbn}/pages/{pagenum}", users.CheckAuth(pages.UpdatePage)).Methods("PUT")
	router.HandleFunc("/api/book/{isbn}/pages", pages.GetPages).Methods("GET")
	router.HandleFunc("/api/book/{isbn}/pages", users.CheckAuth(pages.AddPage)).Methods("POST")

	router.HandleFunc("/api/book/{isbn}", books.GetBook).Methods("GET")
	router.HandleFunc("/api/book/{isbn}", users.CheckAuth(books.DeleteBook)).Methods("DELETE")
	router.HandleFunc("/api/book/{isbn}", users.CheckAuth(books.UpdateBook)).Methods("PUT")
	router.HandleFunc("/api/book", users.CheckAuth(books.AddBook)).Methods("POST")

	router.HandleFunc("/api/auth", users.Auth).Methods("POST")
	router.HandleFunc("/api/logout", users.Logout).Methods("GET")
	router.HandleFunc("/api/users", users.AddUser).Methods("POST")

	router.HandleFunc("/api/user/favs", users.CheckAuth(users.GetFavs)).Methods("GET")
	router.HandleFunc("/api/user/favs/{isbn}", users.CheckAuth(users.CheckBook)).Methods("GET")
	router.HandleFunc("/api/user/favs/{isbn}", users.CheckAuth(users.AddToFavs)).Methods("POST")
	router.HandleFunc("/api/user/favs/{isbn}", users.CheckAuth(users.DeleteFromFavs)).Methods("DELETE")
	router.HandleFunc("/api/user/role", users.CheckAuth(users.GetUserRole)).Methods("GET")

	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./public/images"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/www")))

	log.Println("Server is starting")

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	http.ListenAndServe(":8080", handlers.CORS(headers, credentials, origins, methods)(router))
}
