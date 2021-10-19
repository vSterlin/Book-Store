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

	rows, _ := br.db.Query("SELECT id, title, author, pages, publish_date  FROM books;")
	books := []*entity.Book{}

	for rows.Next() {
		book := &entity.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Pages, &book.PublishDate)
		books = append(books, book)
	}

	return books
}

func (br *BookRepo) GetOne(id int) *entity.Book {
	book := &entity.Book{}
	br.db.QueryRow("SELECT id, title, author, pages, publish_date FROM books WHERE id = $1;", id).Scan(&book.Id, &book.Title, &book.Author, &book.Pages, &book.PublishDate)
	return book
}
