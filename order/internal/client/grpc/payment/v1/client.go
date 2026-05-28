package v1

import paymentV1 "github.com/ianagovitsyn/project/shared/pkg/proto/payment/v1"

type Client struct {
	generatedClient paymentV1.PaymentServiceClient
}

func NewClient(generatedClient paymentV1.PaymentServiceClient) *Client {
	return &Client{
		generatedClient: generatedClient,
	}
}
