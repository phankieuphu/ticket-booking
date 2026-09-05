package models

import "time"

type Permission struct {
	ID             int       `gorm:"column:id;primaryKey;autoIncrement"`
	PermissionName string    `gorm:"column:permission_name;size:50;not null;unique"`
	Resource       string    `gorm:"column:resource;size:255;not null"`
	Action         string    `gorm:"column:action;size:50;not null"`
	Description    string    `gorm:"column:description;size:255"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Permission) TableName() string {
	return "permissions"
}
