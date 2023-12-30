package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id" form:"id" gorm:"primary_key"`
	Username  string    `json:"username" form:"username"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Role      string    `json:"role" form:"role"`
	Name      string    `json:"name" form:"name"`
	Phone     string    `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
	Image string `json:"image" form:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



