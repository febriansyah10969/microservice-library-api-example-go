package app

import "gitlab.com/p9359/backend-prob/febry-go/internal/service"

type BookApp struct {
	service.GeneralService
}

func NewBookApp(service service.GeneralService) *BookApp {
	return &BookApp{service}
}
