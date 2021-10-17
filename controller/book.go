package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vSterlin/bookstore/service"
)

type BookController struct {
	Bs *service.BookService
}

func (bc *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books := bc.Bs.GetMany()
	json.NewEncoder(w).Encode(books)
}

func (bc *BookController) GetBook(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		fmt.Println("provide valid id")
		//TODO
	}

	book := bc.Bs.GetOne(id)

	json.NewEncoder(w).Encode(book)

}

func (bc *BookController) InsertBook(w http.ResponseWriter, r *http.Request) {
}

func (bc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
}
