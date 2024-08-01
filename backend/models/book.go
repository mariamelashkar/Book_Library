package models

type Book struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	Quantity   int    `json:"quantity"`
	AccessTime int    `json:"access_time"` // in minutes
}

var Books []Book
