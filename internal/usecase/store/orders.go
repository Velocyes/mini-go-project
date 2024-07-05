package store

import (
	"log"
	"strconv"

	"github.com/Velocyes/mini-go-project/internal/consts"
	"github.com/Velocyes/mini-go-project/internal/model"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (storeUC *StoreUsecase) SelectAllOrders(limit, offset int) (orderStructs []*model.OrderStruct, err error) {
	var (
		orders []*model.Order
	)

	orders, err = storeUC.DatabaseRepository.SelectAllOrders(limit, offset)
	if err != nil {
		log.Printf("[HTTP][Usecase][Store] SelectAllOrders return err : %v", err)
		return nil, err
	}

	if len(orders) > 0 {
		for _, order := range orders {
			orderStruct := &model.OrderStruct{
				ID:         int(order.ID),
				Status:     int(order.Status),
				TotalPrice: order.TotalPrice,
			}
			for _, orderDetail := range order.OrderDetails {
				orderStruct.Orders = append(orderStruct.Orders, model.OrderDetailStruct{
					ID:        int(orderDetail.ID),
					OrderID:   int(orderDetail.OrderID),
					ProductID: int(orderDetail.ProductID),
					Quantity:  orderDetail.Quantity,
				})
			}
			orderStructs = append(orderStructs, orderStruct)
		}
	}

	return orderStructs, nil
}

func (storeUC *StoreUsecase) SelectOrdersByIDs(orderIDs []int) (orderStructs []*model.OrderStruct, err error) {
	var (
		orders []*model.Order
	)

	orders, err = storeUC.DatabaseRepository.SelectOrdersByIDs(orderIDs)
	if err != nil {
		log.Printf("[HTTP][Usecase][Store] SelectOrdersByIDs return err : %v", err)
		return nil, err
	}

	if len(orders) > 0 {
		for _, order := range orders {
			orderStruct := &model.OrderStruct{
				ID:         int(order.ID),
				Status:     int(order.Status),
				TotalPrice: order.TotalPrice,
			}
			for _, orderDetail := range order.OrderDetails {
				orderStruct.Orders = append(orderStruct.Orders, model.OrderDetailStruct{
					ID:        int(orderDetail.ID),
					OrderID:   int(orderDetail.OrderID),
					ProductID: int(orderDetail.ProductID),
					Quantity:  orderDetail.Quantity,
				})
			}
			orderStructs = append(orderStructs, orderStruct)
		}
	}

	return orderStructs, nil
}

func (storeUC *StoreUsecase) CreateOrders(orderStructs []*model.OrderStruct) error {
	var (
		orders []*model.Order

		// Midtrans
		midtransTD       midtrans.TransactionDetails
		midtransItems    *[]midtrans.ItemDetails
		midtransItemsMap map[int]midtrans.ItemDetails
	)

	if len(orderStructs) > 0 {
		order := &model.Order{}
		midtransItemsMap = map[int]midtrans.ItemDetails{}

		for _, orderStruct := range orderStructs {
			order.Status = uint(orderStruct.Status)
			for _, orderDetailStruct := range orderStruct.Orders {
				order.OrderDetails = append(order.OrderDetails, model.OrderDetail{
					ProductID: uint(orderDetailStruct.ProductID),
					Quantity:  orderDetailStruct.Quantity,
				})

				products, err := storeUC.DatabaseRepository.SelectProductsByIDs([]int{orderDetailStruct.ProductID})
				if err != nil {
					log.Printf("[HTTP][Usecase][Store] CreateOrders - SelectProductsByIDs return err : %v", err)
					return err
				}
				if len(products) > 0 {
					totalPriceProduct := products[0].Price * float64(orderDetailStruct.Quantity)
					order.TotalPrice += totalPriceProduct

					midtransItemsMap[int(products[0].ID)] = midtrans.ItemDetails{
						ID:    strconv.Itoa(int(products[0].ID)),
						Name:  products[0].ProductName,
						Price: int64(products[0].Price),
						Qty:   int32(orderDetailStruct.Quantity),
					}
				} else {
					err = consts.ErrInvalidProductID

					log.Printf("[HTTP][Usecase][Store] CreateOrders - SelectProductsByIDs return err : %v", err)
					return err
				}
			}
			orders = append(orders, order)
		}

		orders, err := storeUC.DatabaseRepository.CreateOrders(orders)
		if err != nil {
			log.Printf("[HTTP][Usecase][Store] CreateOrders return err : %v", err)
			return err
		}

		for _, order := range orders {
			midtransItems = &[]midtrans.ItemDetails{}
			midtransTD = midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(int(order.ID)),
				GrossAmt: int64(order.TotalPrice),
			}
			for _, orderDetail := range order.OrderDetails {
				*midtransItems = append(*midtransItems, midtransItemsMap[int(orderDetail.ProductID)])
			}
		}

		storeUC.MidtransRepository.ChargeTransaction(&coreapi.ChargeReq{
			PaymentType:        coreapi.PaymentTypeGopay,
			TransactionDetails: midtransTD,
			Items:              midtransItems,
		})
		// Ignore error from midtrans, because will return err somehow although the transaction is success
		// if err != nil {
		// 	log.Printf("[HTTP][Usecase][Store] CreateOrders - ChargeTransaction return err : %v", err)
		// 	return err
		// }
		// fmt.Println(coreApiRes)
	}

	return nil
}

func (storeUC *StoreUsecase) DeleteOrdersByIDs(orderIDs []int) error {
	err := storeUC.DatabaseRepository.DeleteOrdersByIDs(orderIDs)
	if err != nil {
		log.Printf("[HTTP][Usecase][Store] CreateOrders - DeleteOrdersByIDs return err : %v", err)
		return err
	}

	return nil
}
