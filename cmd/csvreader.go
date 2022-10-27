package cmd

import (
	"test-go-workshop/model"

	"github.com/gocarina/gocsv"
)

func ReadCSVFromString(data string) (bookModels []*model.BookModel, err error) {
	err = gocsv.UnmarshalString(data, &bookModels)
	if err != nil {
		return bookModels, err
	}
	return bookModels, nil
}
