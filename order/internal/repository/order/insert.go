package order

import (
	"context"
	"fmt"
	"github.com/ianagovitsyn/project/order/internal/repository/converter"
	repoModel "github.com/ianagovitsyn/project/order/internal/repository/model"

	"github.com/ianagovitsyn/project/order/internal/model"
)

func (r *Repository) Insert(ctx context.Context, o model.Order) error {
	ro := converter.OrderToRepo(o)
	_, err := r.DB.Exec(ctx,
		`INSERT INTO orders (uuid, user_uuid, part_uuids, total_price, transaction_uuid, payment_method, status)
		VALUES($1, $2, $3, $4, $5, $6, $7)`,
		ro.OrderUUID,
		ro.UserUUID,
		ro.PartUuids,
		ro.TotalPrice,
		ro.TransactionUUID,
		ro.PaymentMethod,
		ro.Status,
	)
	if err != nil {
		return fmt.Errorf("%w: %w", repoModel.ErrFailedInsert, err)
	}

	return nil
}
