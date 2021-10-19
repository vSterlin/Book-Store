package repository

import (
	"database/sql"

	"github.com/vSterlin/bookstore/internal/entity"
)

var books = []*entity.Book{
	{Id: 1, Author: "Aldous Huxley", Title: "Brave New World", Pages: 311, PublishDate: "1932"},
}

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (br *BookRepo) GetMany() []*entity.Book {

	rows, _ := br.db.Query("SELECT id, title, author, pages, publish_date  FROM books")
	books := []*entity.Book{}

	for rows.Next() {
		book := &entity.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Pages, &book.PublishDate)
		books = append(books, book)
	}

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
