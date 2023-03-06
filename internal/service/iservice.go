package service

import "gitlab.com/p9359/backend-prob/febry-go/internal/repository"

type BookService interface {
}

type bookService struct {
	dao repository.DAO
}

func NewBookService(dao repository.DAO) BookService {
	return &bookService{dao}
}
