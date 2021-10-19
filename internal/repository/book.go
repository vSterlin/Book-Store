package repository

import (
	"database/sql"

	"github.com/vSterlin/bookstore/internal/entity"
)

const (
	GetManySQL = "SELECT id, title, author, pages, publish_date  FROM books;"
	GetOneSQL  = "SELECT id, title, author, pages, publish_date FROM books WHERE id = $1;"
)

type Repo interface {
	GetMany() []*entity.Book
	GetOne(id int) *entity.Book
}
type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (br *BookRepo) GetMany() []*entity.Book {

	rows, _ := br.db.Query(GetManySQL)
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
	br.db.QueryRow(GetOneSQL, id).Scan(&book.Id, &book.Title, &book.Author, &book.Pages, &book.PublishDate)
	return book
}
