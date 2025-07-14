package model

import "time"

type UserRoles struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	RoleID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (UserRoles) TableName() string {
	return "user_roles"
}
