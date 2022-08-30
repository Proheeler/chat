package types

import (
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	ID   string
	Size float64
	Name string
}
