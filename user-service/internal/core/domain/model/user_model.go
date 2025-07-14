package model

import "time"

type User struct {
	ID         int64 `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	Phone      string
	Photo      string
	Address    string
	Lat        float64
	Lon        float64
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Roles      []Role `gorm:"many2many:user_roles;"`
}
