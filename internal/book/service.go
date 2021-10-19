package book

type BookService struct {
	br Repo
}

func NewBookService(br Repo) *BookService {
	return &BookService{br: br}
}

func (bs *BookService) GetMany() []*Book {
	return bs.br.GetMany()
}

func (bs *BookService) GetOne(id int) *Book {
	return bs.br.GetOne(id)
}
