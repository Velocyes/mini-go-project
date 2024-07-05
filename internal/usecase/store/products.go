package store

import (
	"log"

	"github.com/Velocyes/mini-go-project/internal/model"
	"gorm.io/gorm"
)

func (storeUC *StoreUsecase) SelectAllProducts(limit, offset int) (productStructs []*model.ProductStruct, err error) {
	var (
		products []*model.Product
	)

	products, err = storeUC.DatabaseRepository.SelectAllProducts(limit, offset)
	if err != nil {
		log.Printf("[HTTP][Usecase][Store] SelectAllProducts return err : %v", err)
		return nil, err
	}

	if len(products) > 0 {
		for _, product := range products {
			productStructs = append(productStructs, &model.ProductStruct{
				ID:          int(product.ID),
				ProductName: product.ProductName,
				Price:       product.Price,
				Quantity:    product.Quantity,
			})
		}
	}

	return productStructs, nil
}

func (storeUC *StoreUsecase) SelectProductsByIDs(productIDs []int) (productStructs []*model.ProductStruct, err error) {
	var (
		products []*model.Product
	)

	products, err = storeUC.DatabaseRepository.SelectProductsByIDs(productIDs)
	if err != nil {
		log.Printf("[HTTP][Usecase][Store] SelectProductsByIDs return err : %v", err)
		return nil, err
	}

	if len(products) > 0 {
		for _, product := range products {
			productStructs = append(productStructs, &model.ProductStruct{
				ID:          int(product.ID),
				ProductName: product.ProductName,
				Price:       product.Price,
				Quantity:    product.Quantity,
			})
		}
	}

	return productStructs, nil
}

func (storeUC *StoreUsecase) CreateProducts(productStructs []*model.ProductStruct) error {
	var (
		products []*model.Product
	)

	if len(productStructs) > 0 {
		for _, productStruct := range productStructs {
			products = append(products, &model.Product{
				ProductName: productStruct.ProductName,
				Price:       productStruct.Price,
				Quantity:    productStruct.Quantity,
			})
		}

		err := storeUC.DatabaseRepository.CreateProducts(products)
		if err != nil {
			log.Printf("[HTTP][Usecase][Store] CreateProducts return err : %v", err)
			return err
		}
	}

	return nil
}

func (storeUC *StoreUsecase) UpdateProducts(productStructs []*model.ProductStruct) error {
	var (
		products []*model.Product
	)

	if len(productStructs) > 0 {
		for _, productStruct := range productStructs {
			products = append(products, &model.Product{
				Model:       gorm.Model{ID: uint(productStruct.ID)},
				ProductName: productStruct.ProductName,
				Price:       productStruct.Price,
				Quantity:    productStruct.Quantity,
			})
		}

		err := storeUC.DatabaseRepository.UpdateProducts(products)
		if err != nil {
			log.Printf("[HTTP][Usecase][Store] UpdateProducts return err : %v", err)
			return err
		}
	}

	return nil
}

func (storeUC *StoreUsecase) DeleteProductsByIDs(productIDs []int) error {
	err := storeUC.DatabaseRepository.DeleteProductsByIDs(productIDs)
	if err != nil {
		log.Printf("[HTTP][Usecase][Store] DeleteProductsByIDs return err : %v", err)
		return err
	}

	return nil
}
