package datasource

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	accountRoute "github.com/nutikuli/internProject_backend/internal/models/account/controller/http/v1"
	bankRoute "github.com/nutikuli/internProject_backend/internal/models/bank/controller/http/v1"
	customerRoute "github.com/nutikuli/internProject_backend/internal/models/customer/controller/http/v1"
	orderRoute "github.com/nutikuli/internProject_backend/internal/models/order/controller/http/v1"
	productCateRoute "github.com/nutikuli/internProject_backend/internal/models/product-category/controller/http/v1"
	productRoute "github.com/nutikuli/internProject_backend/internal/models/product/controller/http/v1"
	storeRoute "github.com/nutikuli/internProject_backend/internal/models/store/controller/http/v1"
	adminRoute "github.com/nutikuli/internProject_backend/internal/models/admin/controller/http/v1"
	adminPermissionRoute "github.com/nutikuli/internProject_backend/internal/models/adminpermission/controller/http/v1"
)

type RouteRepository interface {
	// TODO: Implemented model routers
}

func InitRoute(db *sqlx.DB, app *fiber.App) {

	apiEntry := app.Group("/api/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})
	adminRoute.UseAdminRoute(db,apiEntry)
	adminPermissionRoute.UseAdminPermissionRoute(db , apiEntry)
	storeRoute.UseStoreRoute(db, apiEntry)
	orderRoute.UseOrderRoute(db, apiEntry)
	productCateRoute.UseProductCategoryRoute(db, apiEntry)
	productRoute.UseProductRoute(db, apiEntry)
	bankRoute.UseBankRoute(db, apiEntry)
	accountRoute.UseAccountRoute(db, apiEntry)
	customerRoute.UseCustomerRoute(db, apiEntry)

}
