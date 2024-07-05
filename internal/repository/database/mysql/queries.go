package database

import (
	"log"

	"github.com/Velocyes/mini-go-project/internal/model"
)

// Products
func (mySql *mySql) SelectAllProducts(limit, offset int) (products []*model.Product, err error) {
	result := mySql.dbInstance.Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		return []*model.Product{}, result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][SelectAllProducts] Return Rows : %d", result.RowsAffected)
	return products, nil
}

func (mySql *mySql) SelectProductsByIDs(productIDs []int) (products []*model.Product, err error) {
	result := mySql.dbInstance.Find(&products, productIDs)
	if result.Error != nil {
		return []*model.Product{}, result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][SelectProductsByIDs] Return Rows : %d", result.RowsAffected)
	return products, nil
}

func (mySql *mySql) CreateProducts(products []*model.Product) error {
	result := mySql.dbInstance.Create(products)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][CreateProducts] Rows Affected : %d", result.RowsAffected)
	return nil
}

func (mySql *mySql) UpdateProducts(products []*model.Product) error {
	result := mySql.dbInstance.Save(products)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][UpdateProducts] Rows Affected : %d", result.RowsAffected)
	return nil
}

func (mySql *mySql) DeleteProductsByIDs(productIDs []int) error {
	result := mySql.dbInstance.Delete(&model.Product{}, productIDs)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][DeleteProducts] Rows Affected : %d", result.RowsAffected)
	return nil
}

// Orders
func (mySql *mySql) SelectAllOrders(limit, offset int) (orders []*model.Order, err error) {
	result := mySql.dbInstance.Limit(limit).Offset(offset).Preload("OrderDetails").Find(&orders)
	if result.Error != nil {
		return []*model.Order{}, result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][SelectAllOrders] Return Rows : %d", result.RowsAffected)
	return orders, nil
}

func (mySql *mySql) SelectOrdersByIDs(orderIDs []int) (orders []*model.Order, err error) {
	result := mySql.dbInstance.Preload("OrderDetails").Find(&orders, orderIDs)
	if result.Error != nil {
		return []*model.Order{}, result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][SelectOrdersByIDs] Return Rows : %d", result.RowsAffected)
	return orders, nil
}

func (mySql *mySql) CreateOrders(orders []*model.Order) ([]*model.Order, error) {
	result := mySql.dbInstance.Create(orders)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][CreateOrders] Rows Affected : %d", result.RowsAffected)
	return orders, nil
}

func (mySql *mySql) DeleteOrdersByIDs(orderIDs []int) error {
	result := mySql.dbInstance.Delete(&model.Order{}, orderIDs)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("[HTTP][Repository][Database][MySQL][DeleteOrdersByIDs] Rows Affected : %d", result.RowsAffected)
	return nil
}
