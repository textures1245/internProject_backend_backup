package entities

type AdminRegisterReq struct {
	Name         string `json:"name" from:"name" binding:"required"`
	Password     string `json:"password" from:"password" binding:"required"`
	Phone        string `json:"phone" from:"phone" binding:"required"`
	Location     string `json:"location" from:"location" binding:"required"`
	Email        string `json:"email" from:"email" binding:"required"`
	Role         string `json:"role" from:"role" binding:"required"`
	Status       bool   `json:"status" from:"status" binding:"required"`
	PermissionID string `json:"permissionid" from:"permissionid" binding:"required"`
} 



func (a *AdminRegisterReq) GetEmail() *string {
	return &a.Email
}

func (a *AdminRegisterReq) GetPassword() *string {
	return &a.Password
}

func (a *AdminRegisterReq) GetRole() *string {
	return &a.Role
}
