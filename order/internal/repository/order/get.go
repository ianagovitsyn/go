package order

import (
	"context"
	"github.com/ianagovitsyn/project/order/internal/repository/converter"
	"github.com/jackc/pgx/v5"

	"github.com/ianagovitsyn/project/order/internal/model"
	repoModel "github.com/ianagovitsyn/project/order/internal/repository/model"
)

func (r *Repository) Get(ctx context.Context, orderUUID string) (model.Order, error) {
	rows, err := r.DB.Query(ctx,
		"SELECT uuid, user_uuid, part_uuids, total_price, transaction_uuid, payment_method, status, created_at, updated_at FROM orders WHERE uuid = $1",
		orderUUID,
	)
	if err != nil {
		return model.Order{}, err
	}

	order, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[repoModel.Order])
	if err != nil {
		return model.Order{}, err
	}

	return converter.OrderToModel(order), nil
}
