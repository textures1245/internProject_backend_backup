package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accRepo "github.com/nutikuli/internProject_backend/internal/models/account/repository"
	_accUse "github.com/nutikuli/internProject_backend/internal/models/account/usecase"
	_adminRepo "github.com/nutikuli/internProject_backend/internal/models/admin/repository"
	_AdminUse "github.com/nutikuli/internProject_backend/internal/models/admin/usecase"
	_adperRepo "github.com/nutikuli/internProject_backend/internal/models/adminpermission/repository"
	_cutomerHand "github.com/nutikuli/internProject_backend/internal/models/customer/controller/http/v1"
	"github.com/nutikuli/internProject_backend/internal/models/customer/repository"
	_cutomerUse "github.com/nutikuli/internProject_backend/internal/models/customer/usecase"
	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	_storeRepo "github.com/nutikuli/internProject_backend/internal/models/store/repository"
	_storeUse "github.com/nutikuli/internProject_backend/internal/models/store/usecase"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
	_fileUse "github.com/nutikuli/internProject_backend/internal/services/file/usecase"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"
)

func UseAccountRoute(db *sqlx.DB, app fiber.Router) {
	authR := app.Group("/account", func(c *fiber.Ctx) error {
		log.Infof("Account : %v", c.Request().URI().String())
		return c.Next()
	})
	adperRepo := _adperRepo.NewAdminPermissionRepository(db)
	//register

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	fileRepo := _fileRepo.NewFileRepository(db)
	fileUse := _fileUse.NewFileUsecase(fileRepo)
	adminRepo := _adminRepo.NewFileRepository(db)
	storeRep := _storeRepo.NewStoreRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, nil, adminRepo, customerRepo, storeRep)
	customerUse := _cutomerUse.NewCustomerUsecase(customerRepo, accUse)
	customerConn := _cutomerHand.NewCustomerHandler(customerUse)

	AdminUseCase := _AdminUse.NewAdminUsecase(adminRepo, fileRepo, accUse, adperRepo, fileUse)
	storeUse := _storeUse.NewStoreUsecase(storeRep, fileRepo, accUse)

	authR.Post("/register",

		customerConn.CreateCustomerAccount)
	//login
	accConn := NewAccountHandler(accUse, storeUse, customerUse, AdminUseCase)
	authR.Post("/login",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "Account",
				Action: "เข้าสู่ระบบ",
			}

			return logger.LogRequest(c, logAction)
		},
		accConn.Login)

	//OTP
	authR.Post("/otp",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "Account",
				Action: "ส่ง OTP",
			}

			return logger.LogRequest(c, logAction)
		},
		accConn.OTP)
	//resetPassword
	authR.Post("/resetpass",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "Account",
				Action: "รีเซ็ตรหัสผ่าน",
			}

			return logger.LogRequest(c, logAction)
		},
		accConn.UpdatePass)

	// get
	authR.Get("/get-customer", accConn.GetAllDataCustomer)
	authR.Get("/get-store", accConn.GetAllDataStore)
	authR.Get("/get-admin", accConn.GetAllDataAdmin)
}
