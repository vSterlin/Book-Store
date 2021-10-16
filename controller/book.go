package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		fmt.Println("provide valid id")
		//TODO
	}

	book := bc.Bs.GetOne(id)

	json.NewEncoder(w).Encode(book)

}

func (bc *BookController) InsertBooks(w http.ResponseWriter, r *http.Request) {
}

func (bc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
}
