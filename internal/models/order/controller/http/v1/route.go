package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_bankRepo "github.com/nutikuli/internProject_backend/internal/models/bank/repository"
	_bankUse "github.com/nutikuli/internProject_backend/internal/models/bank/usecase"
	_customerRepo "github.com/nutikuli/internProject_backend/internal/models/customer/repository"
	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	_orderProdRepo "github.com/nutikuli/internProject_backend/internal/models/order-product/repository"
	_orderProdUse "github.com/nutikuli/internProject_backend/internal/models/order-product/usecase"
	"github.com/nutikuli/internProject_backend/internal/models/order/repository"
	"github.com/nutikuli/internProject_backend/internal/models/order/usecase"
	_prodCate "github.com/nutikuli/internProject_backend/internal/models/product-category/repository"
	_prodUse "github.com/nutikuli/internProject_backend/internal/models/product/usecase"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"

	_prodRepo "github.com/nutikuli/internProject_backend/internal/models/product/repository"

	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
	_fileUse "github.com/nutikuli/internProject_backend/internal/services/file/usecase"
)

func UseOrderRoute(db *sqlx.DB, app fiber.Router) {
	orderR := app.Group("/order", func(c *fiber.Ctx) error {
		log.Infof("order : %v", c.Request().URI().String())
		return c.Next()
	})

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	fileRepo := _fileRepo.NewFileRepository(db)
	fileUse := _fileUse.NewFileUsecase(fileRepo)

	prodCateRepo := _prodCate.NewProductCategoryRepository(db)

	orderRepo := repository.NewOrderRepository(db)

	orderProdRepo := _orderProdRepo.NewOrderProductRepository(db)

	prodRepo := _prodRepo.NewproductRepository(db)
	prodUse := _prodUse.NewProductUsecase(prodRepo, fileRepo, fileUse, prodCateRepo, orderProdRepo)

	orderProdUse := _orderProdUse.NewOrderProductUsecase(orderProdRepo)

	bankRepo := _bankRepo.NewBankRepository(db)
	bankUse := _bankUse.NewBankUsecase(bankRepo, fileRepo)

	customerRepo := _customerRepo.NewCustomerRepository(db)

	orderUse := usecase.NewOrderUsecase(orderRepo, fileRepo, bankUse, prodUse, orderProdUse, customerRepo)

	orderConn := NewOrderHandler(orderUse)

	orderR.Post("/create-order",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: "สร้างรายการสั่งซื้อ",
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.CreateOrder)
	orderR.Get("/get-store-order-by-id/:store_id/:order_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: fmt.Sprintf("ดูข้อมูล Order ของหมายเลข %s ใน Store เลขหมายเลข  %s", c.Params("order_id"), c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.GetOrderById)
	orderR.Get("/get-orders-by-store-id/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: fmt.Sprintf("ดูข้อมูล Order ของ Store เลขหมายเลข  %s", c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.GetOrdersByStoreId)
	orderR.Get("/get-orders-by-customer-id/:customer_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: fmt.Sprintf("ดูข้อมูล Order ของ Customer เลขหมายเลข  %s", c.Params("customer_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.GetOrdersByCustomerId)
	orderR.Patch("/update-order-state/:order_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: fmt.Sprintf("อัปเดตสถานะ Order หมายเลข  %s", c.Params("order_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.UpdateOrderTransportDetailAndState)
}
