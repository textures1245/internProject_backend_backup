package customer

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	_customerDtos "github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerUsecase interface {
	OnCreateCustomerAccount(c *fiber.Ctx, ctx context.Context, customerDatReq *_customerEntities.CustomerRegisterReq) (*_customerDtos.CustomerAccountFileRes, *_accDtos.UserToken, int, error)
	OnGetCustomerById(ctx context.Context, customerId int64) (*_customerDtos.CustomerAccountFileRes, int, error)
	OnUpdateCustomerById(ctx context.Context, userId int64, req *_customerEntities.CustomerUpdateReq) (int, error)
	OnDeletedCustomer(ctx context.Context, Id int64) (int, error)
	OnGetAllUserCustomer(ctx context.Context) ([]*_customerDtos.CustomerRes, int, error)
}
