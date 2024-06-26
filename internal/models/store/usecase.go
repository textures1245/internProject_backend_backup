package store

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
	_storeEntities "github.com/nutikuli/internProject_backend/internal/models/store/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type StoreUsecase interface {
	OnCreateStoreAccount(c *fiber.Ctx, ctx context.Context, storeDatReq *_storeEntities.StoreRegisterReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_storeDtos.StoreWithFileRes, *_accDtos.UserToken, int, error)
	OnGetStoreById(ctx context.Context, storeId int64) (*_storeDtos.StoreWithFileRes, int, error)
	OnUpdateStoreById(c *fiber.Ctx, ctx context.Context, storeId int64, storeDatReq *_storeEntities.StoreUpdatedReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_storeDtos.StoreWithFileRes, int, error)
	OnDeleteStoreById(ctx context.Context, storeId int64) (int, error)
	OnGetStoreAccounts(ctx context.Context) ([]*_storeDtos.StoreWithFileRes, int, error)

	// Order
	// Product
	// Category
	// Bank

}
