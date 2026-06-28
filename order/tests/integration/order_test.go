//go:build integration

package integration

import (
	"context"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ianagovitsyn/project/order/internal/model"
	repoModel "github.com/ianagovitsyn/project/order/internal/repository/model"
)

var _ = Describe("OrderRepository", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(suiteCtx)
	})

	AfterEach(func() {
		_, err := env.Postgres.Conn().Exec(ctx, "DELETE FROM orders")
		Expect(err).ToNot(HaveOccurred())
		cancel()
	})

	Describe("Insert", func() {
		Context("valid order", func() {
			It("should insert order without error", func() {
				order := model.Order{
					OrderUUID:  uuid.New().String(),
					UserUUID:   uuid.New().String(),
					PartUuids:  []string{uuid.New().String(), uuid.New().String()},
					TotalPrice: 199.99,
					Status:     model.StatusPendingPayment,
				}

				err := env.Repo.Insert(ctx, order)
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Describe("Get", func() {
		Context("order exists", func() {
			It("should return the order", func() {
				orderUUID := uuid.New().String()
				order := model.Order{
					OrderUUID:  orderUUID,
					UserUUID:   uuid.New().String(),
					PartUuids:  []string{uuid.New().String()},
					TotalPrice: 100.50,
					Status:     model.StatusPendingPayment,
				}

				err := env.Repo.Insert(ctx, order)
				Expect(err).ToNot(HaveOccurred())

				result, err := env.Repo.Get(ctx, orderUUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(result.OrderUUID).To(Equal(orderUUID))
				Expect(result.UserUUID).To(Equal(order.UserUUID))
				Expect(result.TotalPrice).To(Equal(order.TotalPrice))
				Expect(result.Status).To(Equal(model.StatusPendingPayment))
				Expect(result.PartUuids).To(Equal(order.PartUuids))
			})
		})

		Context("order does not exist", func() {
			It("should return error", func() {
				_, err := env.Repo.Get(ctx, uuid.New().String())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("UpdatePayment", func() {
		Context("order exists", func() {
			It("should update payment fields", func() {
				orderUUID := uuid.New().String()
				order := model.Order{
					OrderUUID:  orderUUID,
					UserUUID:   uuid.New().String(),
					PartUuids:  []string{uuid.New().String()},
					TotalPrice: 250.00,
					Status:     model.StatusPendingPayment,
				}

				err := env.Repo.Insert(ctx, order)
				Expect(err).ToNot(HaveOccurred())

				txUUID := uuid.New().String()
				pm := model.PaymentMethodCard
				order.TransactionUUID = &txUUID
				order.PaymentMethod = &pm
				order.Status = model.StatusPaid

				err = env.Repo.UpdatePayment(ctx, order)
				Expect(err).ToNot(HaveOccurred())

				result, err := env.Repo.Get(ctx, orderUUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(result.Status).To(Equal(model.StatusPaid))
				Expect(*result.TransactionUUID).To(Equal(txUUID))
				Expect(*result.PaymentMethod).To(Equal(model.PaymentMethodCard))
			})
		})

		Context("order does not exist", func() {
			It("should return error", func() {
				txUUID := uuid.New().String()
				pm := model.PaymentMethodCard
				order := model.Order{
					OrderUUID:       uuid.New().String(),
					TransactionUUID: &txUUID,
					PaymentMethod:   &pm,
					Status:          model.StatusPaid,
				}

				err := env.Repo.UpdatePayment(ctx, order)
				Expect(err).To(MatchError(repoModel.ErrOrderNotFound))
			})
		})
	})

	Describe("UpdateStatus", func() {
		Context("order exists", func() {
			It("should update status", func() {
				orderUUID := uuid.New().String()
				order := model.Order{
					OrderUUID:  orderUUID,
					UserUUID:   uuid.New().String(),
					PartUuids:  []string{uuid.New().String()},
					TotalPrice: 300.00,
					Status:     model.StatusPendingPayment,
				}

				err := env.Repo.Insert(ctx, order)
				Expect(err).ToNot(HaveOccurred())

				order.Status = model.StatusCancelled
				err = env.Repo.UpdateStatus(ctx, order)
				Expect(err).ToNot(HaveOccurred())

				result, err := env.Repo.Get(ctx, orderUUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(result.Status).To(Equal(model.StatusCancelled))
			})
		})

		Context("order does not exist", func() {
			It("should return error", func() {
				order := model.Order{
					OrderUUID: uuid.New().String(),
					Status:    model.StatusCancelled,
				}

				err := env.Repo.UpdateStatus(ctx, order)
				Expect(err).To(MatchError(repoModel.ErrOrderNotFound))
			})
		})
	})
})
