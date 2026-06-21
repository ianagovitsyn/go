//go:build e2e

package e2e

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ = Describe("InventoryService", func() {
	var (
		ctx             context.Context
		cancel          context.CancelFunc
		inventoryClient inventoryV1.InventoryServiceClient
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(suiteCtx)

		// Создаём gRPC клиент
		conn, err := grpc.NewClient(
			env.App.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		Expect(err).ToNot(HaveOccurred(), "ожидали успешное подключение к gRPC приложению")

		inventoryClient = inventoryV1.NewInventoryServiceClient(conn)
	})

	AfterEach(func() {
		err := env.ClearInventoryCollection(ctx)
		Expect(err).ToNot(HaveOccurred(), "expected successful clearing of collection")

		cancel()
	})

	Describe("Get by UUID", func() {
		var partUUID string

		Context("Part exists", func() {
			var err error
			BeforeEach(func() {
				partUUID, err = env.InsertTestPart(ctx)
				Expect(err).ToNot(HaveOccurred(), "expected successful inserting of part into MongoDB")
			})

			It("has to successfully return part by UUID", func() {
				resp, err := inventoryClient.GetPart(ctx, &inventoryV1.GetPartRequest{
					Uuid: partUUID,
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(resp.GetPart()).ToNot(BeNil())
				Expect(resp.GetPart().GetUuid()).To(Equal(partUUID))
				Expect(resp.GetPart().GetName()).ToNot(BeEmpty())
				Expect(resp.GetPart().GetDescription()).ToNot(BeEmpty())
				Expect(resp.GetPart().GetPrice()).To(BeNumerically(">", 0))
				Expect(resp.GetPart().GetStockQuantity()).To(BeNumerically(">", 0))
				Expect(resp.GetPart().GetDimensions()).ToNot(BeNil())
				Expect(resp.GetPart().GetManufacturer()).ToNot(BeNil())
				Expect(resp.GetPart().GetManufacturer().GetName()).ToNot(BeEmpty())
				Expect(resp.GetPart().GetTags()).ToNot(BeEmpty())
			})
		})

		Context("Part doesnt exist", func() {
			It("has to return error", func() {
				resp, err := inventoryClient.GetPart(ctx, &inventoryV1.GetPartRequest{
					Uuid: gofakeit.UUID(),
				})

				Expect(err).To(HaveOccurred())
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("ListParts by filter", func() {
		Context("parts exist", func() {
			BeforeEach(func() {
				_, err := env.InsertTestPart(ctx)
				Expect(err).ToNot(HaveOccurred())
				_, err = env.InsertTestPart(ctx)
				Expect(err).ToNot(HaveOccurred())
				_, err = env.InsertTestPart(ctx)
				Expect(err).ToNot(HaveOccurred())
			})

			It("has to return list of parts", func() {
				resp, err := inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{})
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.GetParts()).To(HaveLen(3))
			})
		})

		Context("collection is empty", func() {
			It("has to return empty list", func() {
				resp, err := inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{})
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.GetParts()).To(BeEmpty())
			})
		})
	})
})
