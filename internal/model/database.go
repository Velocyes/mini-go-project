package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string
	Price float64
	Quantity int
	OrderDetails []OrderDetail
}

type Order struct {
	gorm.Model
	Status uint
	TotalPrice float64
	OrderDetails []OrderDetail
}

type OrderDetail struct {
	gorm.Model
	OrderID uint
	ProductID uint
	Quantity int
}