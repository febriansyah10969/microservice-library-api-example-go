package service

import "gitlab.com/p9359/backend-prob/febry-go/internal/repository"

type ReviewService interface {
}

type reviewService struct {
	dao repository.DAO
}

func NewReviewService(dao repository.DAO) ReviewService {
	return &reviewService{dao}
}
