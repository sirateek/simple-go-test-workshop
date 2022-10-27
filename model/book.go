package model

type BookModel struct {
	Id     int    `csv:"id"`
	Title  string `csv:"title"`
	Author string `csv:"author"`
	ISBN   string `csv:"isbn"`
}
