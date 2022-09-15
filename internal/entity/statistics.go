package entity

import (
	"database/sql/driver"
	"time"
)

// StatisticalData - структура для хранения статистических данных
type StatisticalData struct {
	Id     int        `json:"id" db:"id"`
	UserId int        `json:"user_id" db:"user_id"`
	Date   CustomTime `json:"date" db:"date" binding:"required"`
	Views  int        `json:"views" db:"views" binding:"required"`
	Clicks int        `json:"clicks" db:"clicks" binding:"required"`
	Cost   int        `json:"cost" db:"cost" binding:"required"`
}

type CustomTime time.Time

const customTimeFormat = `"2006-01-02"`

//goland:noinspection GoMixedReceiverTypes
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*ct = CustomTime(time.Time{})
		return nil
	}
	t, err := time.Parse(customTimeFormat, string(b))
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if time.Time(ct).IsZero() {
		return []byte("null"), nil
	}
	return []byte(time.Time(ct).Format(customTimeFormat)), nil
}

//goland:noinspection GoMixedReceiverTypes
func (ct *CustomTime) Scan(src any) error {
	*ct = CustomTime(src.(time.Time))
	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (ct CustomTime) Value() (driver.Value, error) {
	return time.Time(ct), nil
}
