package entity

type Book struct {
	Id          int    `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Pages       int    `json:"pages"`
	PublishDate string `json:"publishDate"`
}
