package service

import "gitlab.com/p9359/backend-prob/febry-go/internal/repository"

type GeneralService interface {
}

type generalService struct {
	dao repository.DAO
}

func NewGeneralService(dao repository.DAO) GeneralService {
	return &generalService{dao}
}
