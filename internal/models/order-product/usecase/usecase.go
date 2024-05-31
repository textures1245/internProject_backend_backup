package usecase

import (
	"context"
	"net/http"

	order_product "github.com/nutikuli/internProject_backend/internal/models/order-product"
	"github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
)

type order_productUsecase struct {
	order_productRepo order_product.OrderProductRepository
}

func NewOrderUsecase(order_productRepo order_product.OrderProductRepository) order_product.OrderProductUsecase {
	return &order_productUsecase{
		order_productRepo: order_productRepo,
	}
}

func (s *order_productUsecase) OnCreateOrderProduct(ctx context.Context, orders []*entities.OrderProductCreateReq) ([]*int64, int, error) {
	var createdOrderIDs = make([]*int64, 0)

	for _, order := range orders {
		newOrderID, err := s.order_productRepo.CreateOrder(ctx, order)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		createdOrderIDs = append(createdOrderIDs, newOrderID)
	}

	return createdOrderIDs, http.StatusOK, nil
}