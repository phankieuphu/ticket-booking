package models

import "time"

type Role struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	RoleName  string    `gorm:"column:role_name;size:50;not null;unique"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`

	Permissions []Permission `gorm:"many2many:role_permissions;joinForeignKey:RoleID;joinReferences:PermissionID"`
}

func (Role) TableName() string {
	return "roles"
}

type RolePermission struct {
	RoleID       int       `gorm:"column:role_id;primaryKey"`
	PermissionID int       `gorm:"column:permission_id;primaryKey"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
