package model

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartUuids       []string
	TotalPrice      float64
	TransactionUUID string
	PaymentMethod   PaymentMethod
	Status          Status
}

type PaymentMethod string

const (
	PaymentMethodUnknown       PaymentMethod = "PAYMENT_METHOD_UNKNOWN"
	PaymentMethodCard          PaymentMethod = "PAYMENT_METHOD_CARD"
	PaymentMethodSBP           PaymentMethod = "PAYMENT_METHOD_SBP"
	PaymentMethodCreditCard    PaymentMethod = "PAYMENT_METHOD_CREDIT_CARD"
	PaymentMethodInvestorMoney PaymentMethod = "PAYMENT_METHOD_INVESTOR_MONEY"
)

type Status string

const (
	StatusPendingPayment Status = "PENDING_PAYMENT"
	StatusPaid           Status = "PAID"
	StatusCancelled      Status = "CANCELLED"
)

type CreateOrderParams struct {
	UserUUID  string
	PartUUIDs []string
}
