package repository

import "github.com/vSterlin/bookstore/entity"

var books = []*entity.Book{
	{Id: 1, Author: "Aldous Huxley", Title: "Brave New World", Pages: 311, PublishDate: "1932"},
}

type BookRepo struct{}

func (br *BookRepo) GetMany() []*entity.Book {
	return books
}

func (br *BookRepo) GetOne(id int) *entity.Book {
	return books[id-1]
}

func (br *BookRepo) InsertOne(book *entity.Book) {
	books = append(books, book)
}

func (br *BookRepo) RemoveOne(id int) {
	//TODO
}
