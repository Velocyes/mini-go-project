package repository

import "github.com/Velocyes/mini-go-project/internal/model"

//go:generate mockgen -source=database.go -destination=./../../mocks/repository/mock_database.go
type Database interface {
	SelectAllProducts(limit, offset int) (products []*model.Product, err error)
	SelectProductsByIDs(productIDs []int) (products []*model.Product, err error)
	CreateProducts(products []*model.Product) error
	UpdateProducts(products []*model.Product) error
	DeleteProductsByIDs(productIDs []int) error

	SelectAllOrders(limit, offset int) (orders []*model.Order, err error)
	SelectOrdersByIDs(orderIDs []int) (orders []*model.Order, err error)
	CreateOrders(orders []*model.Order) ([]*model.Order, error)
	DeleteOrdersByIDs(orderIDs []int) error
}
