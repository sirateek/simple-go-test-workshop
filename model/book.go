package model

type BookModel struct {
	Id     int    `csv:"id" json:"id"`
	Title  string `csv:"title" json:"title"`
	Author string `csv:"author" json:"author"`
	ISBN   string `csv:"isbn" json:"isbn"`
}
