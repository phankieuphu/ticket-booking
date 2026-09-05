package dto

type UpdateRole struct {
	ID         int    `json:"id" binding:"required"`
	Permission []int  `json:"permission" binding:"required"`
	Name       string `json:"name"`
}

type CreateRole struct {
	Permission []int  `json:"permission" binding:"required"`
	Name       string `json:"name"`
}

type Role struct {
	ID             int    `json:"id"`
	RoleName       string `json:"role_name"`
	RolePermission []int  `json:"role_permission"`
}
