package book

import (
	"errors"
	"fmt"
	"net/http"
)

func NewClient(host string) Client {
	return Client{host, http.Client{}}
}

type Client struct {
	host string
	http http.Client
}

func (c Client) GetBook(isbn string) (Book, error) {
	uri := fmt.Sprintf("%s/book/%s", c.host, isbn)
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.Header.Add("Accept", "application/json")

	c.http.Do(req)

	return Book{}, errors.New("No book with ISBN 123-4567890123")
}


type Book struct {
	ISBN            string   `json:"isbn"`
	Title           string   `json:"title"`
	Author          string   `json:"author"`
	PublicationData JsonDate `json:"publication_data"`
	InPrint         bool     `json:"in_print"`
	NumberOfPages   int64    `json:"number_of_pages"`
}

type JsonDate struct {
}

type JsonError struct {
	Msg  string `json:"message"`
}