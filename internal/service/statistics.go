package service

import (
	"context"
	"statistics-service/internal/entity"
)

type StatisticsService struct {
	repo StatisticsRepo
}

func NewStatisticsService(repo StatisticsRepo) *StatisticsService {
	return &StatisticsService{repo: repo}
}

func (s *StatisticsService) SaveData(ctx context.Context, userId int, input entity.StatisticalData) (int, error) {
	return s.repo.SaveData(ctx, userId, input)
}

func (s *StatisticsService) GetData(ctx context.Context, userId int) ([]entity.StatisticalData, error) {
	return s.repo.GetData(ctx, userId)
}

func (s *StatisticsService) GetDataById(ctx context.Context, userId, recordId int) (entity.StatisticalData, error) {
	return s.repo.GetDataById(ctx, userId, recordId)
}

func (s *StatisticsService) DeleteData(ctx context.Context, userId, recordId int) error {
	return s.repo.DeleteData(ctx, userId, recordId)
}

func (s *StatisticsService) UpdateData(ctx context.Context, userId, recordId int, input entity.StatisticalData) error {
	return s.repo.UpdateData(ctx, userId, recordId, input)
}
