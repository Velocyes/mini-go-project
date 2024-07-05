package model

// Request Model
type ProductsRequest struct {
	Products []ProductRequest `json:"products"`
}
type ProductRequest struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type OrdersRequest struct {
	Orders []OrderRequest `json:"orders"`
}
type OrderRequest struct {
	ID           int                  `json:"id"`
	Status       int                  `json:"status"`
	OrderDetails []OrderDetailRequest `json:"order_details"`
}
type OrderDetailRequest struct {
	ID        int `json:"id"`
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// Usecase Model
type ProductStruct struct {
	ID          int
	ProductName string
	Price       float64
	Quantity    int
}

type OrderStruct struct {
	ID         int
	Status     int
	TotalPrice float64
	Orders     []OrderDetailStruct
}
type OrderDetailStruct struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
}

// Response Model
type ProductsResponse struct {
	Products []ProductResponse `json:"products"`
}
type ProductResponse struct {
	ID          int     `json:"id,omitempty"`
	ProductName string  `json:"product_name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
}

type OrdersResponse struct {
	Orders []OrderResponse `json:"orders`
}
type OrderResponse struct {
	ID           int                   `json:"id,omitempty"`
	Status       int                   `json:"status,omitempty"`
	TotalPrice   float64               `json:"total_price,omitempty"`
	OrderDetails []OrderDetailResponse `json:"order_details,omitempty"`
}
type OrderDetailResponse struct {
	ID        int `json:"id,omitempty"`
	OrderID   int `json:"order_id,omitempty"`
	ProductID int `json:"product_id,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}
