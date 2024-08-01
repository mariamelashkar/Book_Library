package handlers
 import (
	"bookstore/models"
	"errors"
 )

func SearchForBook(id string) (*models.Book, error) { //err exists because if we don't find the book we gonna return it

	for index, value := range models.Books {
		if value.Name == id {
			return &models.Books[index], nil //we return a pointer because if we want modify it from a different function
		}
	}
	return nil, errors.New("book not found")
}