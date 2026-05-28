package model

type PaymentInfo struct {
	OrderUUID     string
	UserUUID      string
	PaymentMethod PaymentMethod
}

type PaymentMethod int

const (
	PaymentMethodUnknown PaymentMethod = iota
	PaymentMethodCard
	PaymentMethodSBP
	PaymentMethodCreditCard
	PaymentMethodInvestorMoney
)

type Transaction struct {
	TransactionUUID string
	IdempotencyKey  string // Прост чтоб структура имела смысл
}
