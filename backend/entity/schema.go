package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Firstname string `valid:"required~Firstname cannot be blank"`
	Lastname  string `valid:"required~Lastname cannot be blank"`
	Tel       string `valid:"required~Tel cannot be blank"`
	Room      string `valid:"required~Room cannot be blank"`
	Email     string `valid:"required~Email cannot be blank"`
	Password  string `valid:"required~Password cannot be blank"`
	Role      string `valid:"required~Role cannot be blank"`
	Path      string `valid:"required~Path cannot be blank"`
}
