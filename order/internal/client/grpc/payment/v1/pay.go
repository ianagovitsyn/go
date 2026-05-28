package v1

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/client/converter"
	"github.com/ianagovitsyn/project/order/internal/model"
)

func (c *Client) Pay(ctx context.Context, order model.Order, paymentMethod model.PaymentMethod) (string, error) {
	payResponse, err := c.generatedClient.PayOrder(ctx, converter.PayToGrpc(order, paymentMethod))
	if err != nil {
		return "", err
	}

	return payResponse.TransactionUuid, nil
}
