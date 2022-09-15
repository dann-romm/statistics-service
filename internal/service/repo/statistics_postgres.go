package repo

import (
	"context"
	"fmt"
	"statistics-service/internal/entity"
	"statistics-service/pkg/postgres"
	"time"
)

type StatisticsRepo struct {
	*postgres.Postgres
}

func NewStatisticsRepo(pg *postgres.Postgres) *StatisticsRepo {
	return &StatisticsRepo{pg}
}

func (r *StatisticsRepo) SaveData(ctx context.Context, userId int, input entity.StatisticalData) (int, error) {
	sql, args, err := r.Builder.
		Insert("statistics").
		Columns("user_id", "date", "views", "clicks", "cost").
		Values(userId, time.Time(input.Date), input.Views, input.Clicks, input.Cost).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("repo - StatisticsRepo - SaveData - r.Builder: %w", err)
	}

	var id int
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo - StatisticsRepo - SaveData - r.Pool.QueryRow: %w", err)
	}

	return id, nil
}

func (r *StatisticsRepo) GetData(ctx context.Context, userId int) ([]entity.StatisticalData, error) {
	sql, args, err := r.Builder.
		Select("id", "user_id", "date", "views", "clicks", "cost").
		From("statistics").
		Where("user_id = ?", userId).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("repo - StatisticsRepo - GetData - r.Builder: %w", err)
	}

	var data []entity.StatisticalData
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repo - StatisticsRepo - GetData - r.Pool.Query: %w", err)
	}

	for rows.Next() {
		var d entity.StatisticalData
		err = rows.Scan(&d.Id, &d.UserId, &d.Date, &d.Views, &d.Clicks, &d.Cost)
		if err != nil {
			return nil, fmt.Errorf("repo - StatisticsRepo - GetData - rows.Scan: %w", err)
		}

		data = append(data, d)
	}

	return data, nil
}

func (r *StatisticsRepo) GetDataById(ctx context.Context, userId, recordId int) (entity.StatisticalData, error) {
	sql, args, err := r.Builder.
		Select("id", "user_id", "date", "views", "clicks", "cost").
		From("statistics").
		Where("user_id = ?", userId).
		Where("id = ?", recordId).
		ToSql()

	if err != nil {
		return entity.StatisticalData{}, fmt.Errorf("repo - StatisticsRepo - GetDataById - r.Builder: %w", err)
	}

	var d entity.StatisticalData
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&d.Id, &d.UserId, &d.Date, &d.Views, &d.Clicks, &d.Cost)
	if err != nil {
		return entity.StatisticalData{}, fmt.Errorf("repo - StatisticsRepo - GetDataById - r.Pool.QueryRow: %w", err)
	}

	return d, nil
}

func (r *StatisticsRepo) UpdateData(ctx context.Context, userId, recordId int, input entity.StatisticalData) error {
	sql, args, err := r.Builder.
		Update("statistics").
		Set("date", time.Time(input.Date)).
		Set("views", input.Views).
		Set("clicks", input.Clicks).
		Set("cost", input.Cost).
		Where("user_id = ?", userId).
		Where("id = ?", recordId).
		ToSql()

	if err != nil {
		return fmt.Errorf("repo - StatisticsRepo - UpdateData - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo - StatisticsRepo - UpdateData - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *StatisticsRepo) DeleteData(ctx context.Context, userId, recordId int) error {
	sql, args, err := r.Builder.
		Delete("statistics").
		Where("user_id = ?", userId).
		Where("id = ?", recordId).
		ToSql()

	if err != nil {
		return fmt.Errorf("repo - StatisticsRepo - DeleteData - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo - StatisticsRepo - DeleteData - r.Pool.Exec: %w", err)
	}

	return nil
}
