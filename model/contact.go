package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string  `json:"name"`
	Number  string  `json:"number"`
}