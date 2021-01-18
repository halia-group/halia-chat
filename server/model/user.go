package model

import "time"

type User struct {
	ID        int `gorm:"primary_key"`
	Username  string
	Password  string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (User) TableName() string {
	return "user"
}
