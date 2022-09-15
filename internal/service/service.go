package service

import (
	"statistics-service/internal/service/repo"
)

type Service struct {
	Auth
	Statistics
}

func New(repo *repo.Repository) *Service {
	return &Service{
		Auth:       NewAuthService(repo),
		Statistics: NewStatisticsService(repo),
	}
}
