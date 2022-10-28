package cmd_test

import (
	"test-go-workshop/cmd"
	"test-go-workshop/model"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Books", Label("unit"), func() {
	
	It("can map data from csv to book", func() {
		data := `id,title,author,isbn
		1,Book1,Grid,SB-1234`
		book := model.BookModel{
			Id:     1,
			Title:  "Book1",
			Author: "Grid",
			ISBN:   "SB-1234",
		}
		bookShelf := []*model.BookModel{&book}
		Expect(cmd.ReadCSVFromString(data)).To(Equal(bookShelf))

	})
})
