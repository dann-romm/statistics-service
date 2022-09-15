package service

import (
	"context"
	"statistics-service/internal/entity"
)

type (
	Auth interface {
		CreateUser(context.Context, entity.User) (int, error)
		GenerateToken(context.Context, string, string) (string, error)
		ParseToken(token string) (int, error)
	}

	Statistics interface {
		SaveData(context.Context, int, entity.StatisticalData) (int, error)
		GetData(context.Context, int) ([]entity.StatisticalData, error)
		GetDataById(context.Context, int, int) (entity.StatisticalData, error)
		UpdateData(context.Context, int, int, entity.StatisticalData) error
		DeleteData(context.Context, int, int) error
	}

	AuthRepo interface {
		CreateUser(context.Context, entity.User) (int, error)
		GetUser(context.Context, string, string) (entity.User, error)
	}

	StatisticsRepo interface {
		SaveData(context.Context, int, entity.StatisticalData) (int, error)
		GetData(context.Context, int) ([]entity.StatisticalData, error)
		GetDataById(context.Context, int, int) (entity.StatisticalData, error)
		UpdateData(context.Context, int, int, entity.StatisticalData) error
		DeleteData(context.Context, int, int) error
	}
)
