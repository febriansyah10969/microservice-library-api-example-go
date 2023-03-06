package app

import "gitlab.com/p9359/backend-prob/febry-go/internal/service"

type ReviewApp struct {
	service.ReviewService
}

func NewReviewApp(service service.ReviewService) *ReviewApp {
	return &ReviewApp{service}
}
