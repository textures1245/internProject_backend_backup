package entities

type StoreUpdatedReq struct {
	StoreName     string `json:"store_name" form:"store_name" binding:"required"`
	StoreLocation string `json:"store_location" form:"store_location" binding:"required"`
	Name          string `json:"acc_name" form:"acc_name" binding:"required"`
	Phone         string ` json:"acc_phone" form:"acc_phone" binding:"required"`
	Location      string ` json:"acc_location" form:"acc_location" binding:"required"`
	Email         string ` json:"acc_email" form:"acc_email" binding:"required"`
	Status        bool   ` json:"acc_status" form:"acc_status" binding:"required"`
}
