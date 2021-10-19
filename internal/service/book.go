package service

import (
	"github.com/vSterlin/bookstore/internal/entity"
	"github.com/vSterlin/bookstore/internal/repository"
)

type BookService struct {
	br *repository.BookRepo
}

func NewBookService(br *repository.BookRepo) *BookService {
	return &BookService{br: br}
}

func (bs *BookService) GetMany() []*entity.Book {
	return bs.br.GetMany()
}

func (bs *BookService) GetOne(id int) *entity.Book {
	return bs.br.GetOne(id)
}

func (bs *BookService) InsertOne(book *entity.Book) {
	bs.br.InsertOne(book)
}

func (bs *BookService) RemoveOne(id int) {
	//TODO
}
