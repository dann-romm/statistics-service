package repo

import (
	"statistics-service/pkg/postgres"
)

type Repository struct {
	*AuthRepo
	*StatisticsRepo
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{
		AuthRepo:       NewAuthRepo(pg),
		StatisticsRepo: NewStatisticsRepo(pg),
	}
}
