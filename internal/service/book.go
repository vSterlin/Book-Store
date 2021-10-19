package service

import (
	"github.com/vSterlin/bookstore/internal/entity"
	"github.com/vSterlin/bookstore/internal/repository"
)

type BookService struct {
	br repository.Repo
}

func NewBookService(br repository.Repo) *BookService {
	return &BookService{br: br}
}

func (bs *BookService) GetMany() []*entity.Book {
	return bs.br.GetMany()
}

func (bs *BookService) GetOne(id int) *entity.Book {
	return bs.br.GetOne(id)
}
