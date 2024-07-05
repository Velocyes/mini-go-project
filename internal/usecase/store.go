package usecase

import "github.com/Velocyes/mini-go-project/internal/model"

//go:generate mockgen -source=store.go -destination=./../../mocks/usecase/mock_store.go
type StoreUsecase interface {
	SelectAllProducts(limit, offset int) (products []*model.ProductStruct, err error)
	SelectProductsByIDs(productIDs []int) (products []*model.ProductStruct, err error)
	CreateProducts(products []*model.ProductStruct) error
	UpdateProducts(products []*model.ProductStruct) error
	DeleteProductsByIDs(productIDs []int) error

	SelectAllOrders(limit, offset int) (orders []*model.OrderStruct, err error)
	SelectOrdersByIDs(orderIDs []int) (orders []*model.OrderStruct, err error)
	CreateOrders(orders []*model.OrderStruct) error
	DeleteOrdersByIDs(orderIDs []int) error
}
