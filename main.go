package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vSterlin/bookstore/controller"
	"github.com/vSterlin/bookstore/repository"
	"github.com/vSterlin/bookstore/service"
)

func main() {

	br := &repository.BookRepo{}
	bs := &service.BookService{Br: br}
	bc := &controller.BookController{Bs: bs}
	r := mux.NewRouter()

	r.HandleFunc("/books", bc.GetBooks)
	r.HandleFunc("/books/{id}", bc.GetBook)
	r.HandleFunc("/books", bc.InsertBooks).Methods(http.MethodPost)
	r.HandleFunc("/books", bc.DeleteBook).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", r)
}
