package model

import "time"

type Order struct {
	OrderUUID       string     `db:"uuid"`
	UserUUID        string     `db:"user_uuid"`
	PartUuids       []string   `db:"part_uuids"`
	TotalPrice      float64    `db:"total_price"`
	TransactionUUID string     `db:"transaction_uuid"`
	PaymentMethod   string     `db:"payment_method"`
	Status          string     `db:"status"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
}
