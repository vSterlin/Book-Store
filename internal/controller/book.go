package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vSterlin/bookstore/internal/service"
)

type BookController struct {
	bs *service.BookService
}

func NewBookController(bs *service.BookService) *BookController {
	return &BookController{bs: bs}
}

func (bc *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books := bc.bs.GetMany()
	json.NewEncoder(w).Encode(books)
}

func (bc *BookController) GetBook(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		fmt.Println("provide valid id")
		//TODO
	}

	book := bc.bs.GetOne(id)

	json.NewEncoder(w).Encode(book)

}
