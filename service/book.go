package service

import (
	"github.com/vSterlin/bookstore/entity"
	"github.com/vSterlin/bookstore/repository"
)

type BookService struct {
	Br *repository.BookRepo
}

func (bs *BookService) GetMany() []*entity.Book {
	return bs.Br.GetMany()
}

func (bs *BookService) GetOne(id int) *entity.Book {
	return bs.Br.GetOne(id)
}

func (bs *BookService) InsertOne(book *entity.Book) {
	bs.Br.InsertOne(book)
}

func (bs *BookService) RemoveOne(id int) {
	//TODO
}
