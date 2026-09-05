package models

import "time"

type User struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement"`
	Username     string    `gorm:"column:username;size:255;not null;unique"`
	PasswordHash string    `gorm:"column:password_hash;size:255;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`

	Profile UserProfile `gorm:"foreignKey:UserID"`
	Roles   []Role      `gorm:"many2many:user_roles;joinForeignKey:UserID;joinReferences:RoleID"`
}

func (User) TableName() string {
	return "users"
}

type UserProfile struct {
	UserID    int       `gorm:"column:user_id;primaryKey"`
	FirstName string    `gorm:"column:first_name;size:255;not null"`
	LastName  string    `gorm:"column:last_name;size:255;not null"`
	Address   string    `gorm:"column:address;size:255"`
	Phone     string    `gorm:"column:phone;size:20;not null"`
	Email     string    `gorm:"column:email;size:320;not null;unique"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}

type UserRole struct {
	UserID    int       `gorm:"column:user_id;primaryKey"`
	RoleID    int       `gorm:"column:role_id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
