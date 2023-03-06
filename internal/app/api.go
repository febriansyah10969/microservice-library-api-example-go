package app

import "gitlab.com/p9359/backend-prob/febry-go/internal/service"

type BookApp struct {
	service.BookService
}

func NewBookApp(service service.BookService) *BookApp {
	return &BookApp{service}
}
