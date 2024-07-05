package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	resUtils "github.com/tboehle/go-response-utils"

	"github.com/Velocyes/mini-go-project/internal/consts"
	"github.com/Velocyes/mini-go-project/internal/model"
)

// Products
func (a *APIServer) getAllProducts(w http.ResponseWriter, r *http.Request) {
	var (
		limit, offset int

		err error

		productStructs   []*model.ProductStruct
		productsResponse model.ProductsResponse
	)

	params := r.URL.Query()
	if limitStr := params.Get("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("[HTTP][Delivery] getAllProducts - Failed to parsing url, return err : %v", err)
			resUtils.WithError(w, r, http.StatusBadRequest, err)
			return
		}
	} else {
		limit = -1
	}

	if offsetStr := params.Get("offset"); offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("[HTTP][Delivery] getAllProducts - Failed to parsing url, return err : %v", err)
			resUtils.WithError(w, r, http.StatusBadRequest, err)
			return
		}
	} else {
		offset = -1
	}

	productStructs, err = a.StoreUsecase.SelectAllProducts(limit, offset)
	if err != nil {
		log.Printf("[HTTP][Delivery] getAllProducts - SelectAllProducts return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	for _, productStruct := range productStructs {
		productsResponse.Products = append(productsResponse.Products, model.ProductResponse{
			ID:          productStruct.ID,
			ProductName: productStruct.ProductName,
			Price:       productStruct.Price,
			Quantity:    productStruct.Quantity,
		})
	}

	resUtils.With(w, r, http.StatusCreated, model.Response{Datas: productsResponse, Message: "Product is succesfully fetched"})
	return
}

func (a *APIServer) getProductsByIDs(w http.ResponseWriter, r *http.Request) {
	var (
		muxVars map[string]string

		productIDs       []int
		productStructs   []*model.ProductStruct
		productsResponse model.ProductsResponse
	)

	muxVars = mux.Vars(r)
	strIds, ok := muxVars["ids"]
	if !ok {
		err := consts.ErrInvalidUrlVariable

		log.Printf("[HTTP][Delivery] getProductsByIDs - Failed to parsing url, return err : %v", err)
		resUtils.WithError(w, r, http.StatusBadRequest, err)
		return
	}

	splittedStrIDs := strings.Split(strIds, ",")
	for _, strID := range splittedStrIDs {
		intID, err := strconv.Atoi(strID)
		if err != nil {
			log.Printf("[HTTP][Delivery] getProductsByIDs - Failed to parsing url, return err : %v", err)
			resUtils.WithError(w, r, http.StatusInternalServerError, err)
			return
		}
		productIDs = append(productIDs, intID)
	}

	productStructs, err := a.StoreUsecase.SelectProductsByIDs(productIDs)
	if err != nil {
		log.Printf("[HTTP][Delivery] getProductsByIDs - SelectProductsByIDs return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	for _, productStruct := range productStructs {
		productsResponse.Products = append(productsResponse.Products, model.ProductResponse{
			ID:          productStruct.ID,
			ProductName: productStruct.ProductName,
			Price:       productStruct.Price,
			Quantity:    productStruct.Quantity,
		})
	}

	resUtils.With(w, r, http.StatusCreated, model.Response{Datas: productsResponse, Message: "Product is succesfully fetched"})
	return
}

func (a *APIServer) insertNewProducts(w http.ResponseWriter, r *http.Request) {
	var (
		productsRequest model.ProductsRequest

		productStructs []*model.ProductStruct
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[HTTP][Delivery] insertNewProducts - Failed read request body, return err : %v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := json.Unmarshal(body, &productsRequest); err != nil {
		log.Printf("[HTTP][Delivery] insertNewProducts - Failed unmarshal request body, return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	for _, productRequest := range productsRequest.Products {
		productStructs = append(productStructs, &model.ProductStruct{
			ProductName: productRequest.ProductName,
			Price:       productRequest.Price,
			Quantity:    productRequest.Quantity,
		})
	}

	err = a.StoreUsecase.CreateProducts(productStructs)
	if err != nil {
		log.Printf("[HTTP][Delivery] insertNewProducts - CreateProducts return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	resUtils.With(w, r, http.StatusCreated, model.Response{Message: "Product is succesfully created"})
	return
}

func (a *APIServer) alterProducts(w http.ResponseWriter, r *http.Request) {
	var (
		productsRequest model.ProductsRequest

		productStructs []*model.ProductStruct
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[HTTP][Delivery] alterProducts - Failed read request body, return err : %v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := json.Unmarshal(body, &productsRequest); err != nil {
		log.Printf("[HTTP][Delivery] alterProducts - Failed unmarshal request body, return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	for _, productRequest := range productsRequest.Products {
		productStructs = append(productStructs, &model.ProductStruct{
			ID:          productRequest.ID,
			ProductName: productRequest.ProductName,
			Price:       productRequest.Price,
			Quantity:    productRequest.Quantity,
		})
	}

	err = a.StoreUsecase.UpdateProducts(productStructs)
	if err != nil {
		log.Printf("[HTTP][Delivery] alterProducts - UpdateProducts return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	resUtils.With(w, r, http.StatusOK, model.Response{Message: "Product is succesfully updated"})
	return
}

func (a *APIServer) deleteProductsByIDs(w http.ResponseWriter, r *http.Request) {
	var (
		muxVars map[string]string

		productIDs       []int
		productsResponse model.ProductsResponse
	)

	muxVars = mux.Vars(r)
	strIds, ok := muxVars["ids"]
	if !ok {
		err := consts.ErrInvalidUrlVariable

		log.Printf("[HTTP][Delivery] deleteProductsByIDs - Failed to parsing url, return err : %v", err)
		resUtils.WithError(w, r, http.StatusBadRequest, err)
		return
	}

	splittedStrIDs := strings.Split(strIds, ",")
	for _, strID := range splittedStrIDs {
		intID, err := strconv.Atoi(strID)
		if err != nil {
			log.Printf("[HTTP][Delivery] deleteProductsByIDs - Failed to parsing url, return err : %v", err)
			resUtils.WithError(w, r, http.StatusInternalServerError, err)
			return
		}
		productIDs = append(productIDs, intID)
	}

	err := a.StoreUsecase.DeleteProductsByIDs(productIDs)
	if err != nil {
		log.Printf("[HTTP][Delivery] deleteProductsByIDs - DeleteProductsByIDs return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	resUtils.With(w, r, http.StatusOK, model.Response{Datas: productsResponse, Message: "Product is succesfully deleted"})
	return
}

// Orders
func (a *APIServer) getAllOrders(w http.ResponseWriter, r *http.Request) {
	var (
		limit, offset int

		err error

		orderStructs   []*model.OrderStruct
		ordersResponse model.OrdersResponse
	)

	params := r.URL.Query()
	if limitStr := params.Get("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("[HTTP][Delivery] getAllOrders - Failed to parsing url, return err : %v", err)
			resUtils.WithError(w, r, http.StatusBadRequest, err)
			return
		}
	} else {
		limit = -1
	}

	if offsetStr := params.Get("offset"); offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("[HTTP][Delivery] getAllOrders - Failed to parsing url, return err : %v", err)
			resUtils.WithError(w, r, http.StatusBadRequest, err)
			return
		}
	} else {
		offset = -1
	}

	orderStructs, err = a.StoreUsecase.SelectAllOrders(limit, offset)
	if err != nil {
		log.Printf("[HTTP][Delivery] getAllOrders - SelectAllOrders return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	for _, orderStruct := range orderStructs {
		orderResponse := &model.OrderResponse{
			ID:         int(orderStruct.ID),
			Status:     int(orderStruct.Status),
			TotalPrice: orderStruct.TotalPrice,
		}
		for _, orderDetailStruct := range orderStruct.Orders {
			orderResponse.OrderDetails = append(orderResponse.OrderDetails, model.OrderDetailResponse{
				ID:        int(orderDetailStruct.ID),
				OrderID:   int(orderDetailStruct.OrderID),
				ProductID: int(orderDetailStruct.ProductID),
				Quantity:  orderDetailStruct.Quantity,
			})
		}
		ordersResponse.Orders = append(ordersResponse.Orders, *orderResponse)
	}

	resUtils.With(w, r, http.StatusCreated, model.Response{Datas: ordersResponse, Message: "Order is succesfully fetched"})
	return
}

func (a *APIServer) getOrdersByIDs(w http.ResponseWriter, r *http.Request) {
	var (
		muxVars map[string]string

		orderIDs       []int
		orderStructs   []*model.OrderStruct
		ordersResponse model.OrdersResponse
	)

	muxVars = mux.Vars(r)
	strIds, ok := muxVars["ids"]
	if !ok {
		err := consts.ErrInvalidUrlVariable

		log.Printf("[HTTP][Delivery] getOrdersByIDs - Failed to parsing url, return err : %v", err)
		resUtils.WithError(w, r, http.StatusBadRequest, err)
		return
	}

	splittedStrIDs := strings.Split(strIds, ",")
	for _, strID := range splittedStrIDs {
		intID, err := strconv.Atoi(strID)
		if err != nil {
			log.Printf("[HTTP][Delivery] getOrdersByIDs - Failed to parsing url, return err : %v", err)
			resUtils.WithError(w, r, http.StatusInternalServerError, err)
			return
		}
		orderIDs = append(orderIDs, intID)
	}

	orderStructs, err := a.StoreUsecase.SelectOrdersByIDs(orderIDs)
	if err != nil {
		log.Printf("[HTTP][Delivery] getOrdersByIDs - SelectOrdersByIDs return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	for _, orderStruct := range orderStructs {
		orderResponse := &model.OrderResponse{
			ID:         int(orderStruct.ID),
			Status:     int(orderStruct.Status),
			TotalPrice: orderStruct.TotalPrice,
		}
		for _, orderDetailStruct := range orderStruct.Orders {
			orderResponse.OrderDetails = append(orderResponse.OrderDetails, model.OrderDetailResponse{
				ID:        int(orderDetailStruct.ID),
				OrderID:   int(orderDetailStruct.OrderID),
				ProductID: int(orderDetailStruct.ProductID),
				Quantity:  orderDetailStruct.Quantity,
			})
		}
		ordersResponse.Orders = append(ordersResponse.Orders, *orderResponse)
	}

	resUtils.With(w, r, http.StatusCreated, model.Response{Datas: ordersResponse, Message: "Order is succesfully fetched"})
	return
}

func (a *APIServer) insertNewOrders(w http.ResponseWriter, r *http.Request) {
	var (
		ordersRequest model.OrdersRequest

		orderStructs []*model.OrderStruct
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[HTTP][Delivery] insertNewOrders - Failed read request body, return err : %v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := json.Unmarshal(body, &ordersRequest); err != nil {
		log.Printf("[HTTP][Delivery] insertNewOrders - Failed unmarshal request body, return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	for _, orderRequest := range ordersRequest.Orders {
		orderStruct := &model.OrderStruct{
			Status: orderRequest.Status,
		}
		for _, orderDetailRequest := range orderRequest.OrderDetails {
			orderStruct.Orders = append(orderStruct.Orders, model.OrderDetailStruct{
				ProductID: orderDetailRequest.ProductID,
				Quantity:  orderDetailRequest.Quantity,
			})
		}
		orderStructs = append(orderStructs, orderStruct)
	}

	err = a.StoreUsecase.CreateOrders(orderStructs)
	if err != nil {
		log.Printf("[HTTP][Delivery] insertNewOrders - CreateOrders return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	resUtils.With(w, r, http.StatusCreated, model.Response{Message: "Order is succesfully created"})
	return
}

func (a *APIServer) deleteOrdersByIDs(w http.ResponseWriter, r *http.Request) {
	var (
		muxVars map[string]string

		orderIDs       []int
		OrdersResponse model.OrdersResponse
	)

	muxVars = mux.Vars(r)
	strIds, ok := muxVars["ids"]
	if !ok {
		err := consts.ErrInvalidUrlVariable

		log.Printf("[HTTP][Delivery] deleteOrdersByIDs - Failed to parsing url, return err : %v", err)
		resUtils.WithError(w, r, http.StatusBadRequest, err)
		return
	}

	splittedStrIDs := strings.Split(strIds, ",")
	for _, strID := range splittedStrIDs {
		intID, err := strconv.Atoi(strID)
		if err != nil {
			log.Printf("[HTTP][Delivery] deleteOrdersByIDs - Failed to parsing url, return err : %v", err)
			resUtils.WithError(w, r, http.StatusInternalServerError, err)
			return
		}
		orderIDs = append(orderIDs, intID)
	}

	err := a.StoreUsecase.DeleteOrdersByIDs(orderIDs)
	if err != nil {
		log.Printf("[HTTP][Delivery] deleteOrdersByIDs - DeleteOrdersByIDs return err :%v", err)
		resUtils.WithError(w, r, http.StatusInternalServerError, err)
		return
	}

	resUtils.With(w, r, http.StatusOK, model.Response{Datas: OrdersResponse, Message: "Order is succesfully deleted"})
	return
}
