package book

import (
	"database/sql"
)

const (
	GetManySQL = "SELECT id, title, author, pages, publish_date  FROM books;"
	GetOneSQL  = "SELECT id, title, author, pages, publish_date FROM books WHERE id = $1;"
)

type Repo interface {
	GetMany() []*Book
	GetOne(id int) *Book
}
type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (br *BookRepo) GetMany() []*Book {

	rows, _ := br.db.Query(GetManySQL)
	books := []*Book{}

	for rows.Next() {
		book := &Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Pages, &book.PublishDate)
		books = append(books, book)
	}

	return books
}

func (br *BookRepo) GetOne(id int) *Book {
	book := &Book{}
	br.db.QueryRow(GetOneSQL, id).Scan(&book.Id, &book.Title, &book.Author, &book.Pages, &book.PublishDate)
	return book
}
