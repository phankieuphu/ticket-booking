package entity

type UpdateRole struct {
	ID         int
	Permission []int
	Name       string
}

type CreateRole struct {
	Permission []int
	Name       string
}

type Role struct {
	ID             int
	RoleName       string
	RolePermission []int
}
