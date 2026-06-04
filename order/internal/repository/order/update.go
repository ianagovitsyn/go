package order

import (
	"context"
	"github.com/ianagovitsyn/project/order/internal/repository/converter"
	repoModel "github.com/ianagovitsyn/project/order/internal/repository/model"
	"time"

	"github.com/ianagovitsyn/project/order/internal/model"
)

func (r *Repository) UpdatePayment(ctx context.Context, o model.Order) error {
	or := converter.OrderToRepo(o)

	tag, err := r.DB.Exec(ctx, `
		UPDATE orders
		SET transaction_uuid = $1,
		    payment_method = $2,
		    status = $3,
		    updated = $4
		WHERE uuid = $5`,
		or.TransactionUUID,
		or.PaymentMethod,
		or.Status,
		time.Now(),
		or.OrderUUID,
	)
	if err != nil {
		return repoModel.ErrFailedUpdate
	}
	if tag.RowsAffected() == 0 {
		return repoModel.ErrOrderNotFound
	}

	return nil
}

func (r *Repository) UpdateStatus(ctx context.Context, o model.Order) error {
	or := converter.OrderToRepo(o)

	tag, err := r.DB.Exec(ctx, `
		UPDATE orders
		SET status = $1,
		    updated = $2
		WHERE uuid = $3`,
		or.Status,
		time.Now(),
		or.OrderUUID,
	)
	if err != nil {
		return repoModel.ErrFailedUpdate
	}
	if tag.RowsAffected() == 0 {
		return repoModel.ErrOrderNotFound
	}

	return nil
}
