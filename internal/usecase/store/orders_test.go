package store

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/Velocyes/mini-go-project/internal/consts"
	"github.com/Velocyes/mini-go-project/internal/model"
	"github.com/Velocyes/mini-go-project/internal/repository"
	mockRepo "github.com/Velocyes/mini-go-project/mocks/repository"
	"github.com/golang/mock/gomock"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

func TestUsecase_SelectAllOrders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)
	limit, offset := -1, -1

	mockOrderIDs := []int{1}
	mockOrderDetailIDs := []int{1}

	mockProductIDs := []int{1}
	mockProductQuantity := map[int]int{
		mockProductIDs[0]: 5,
	}
	mockProductPrices := map[int]float64{
		mockProductIDs[0]: 50000,
	}
	mockOrders := map[int]model.Order{
		mockProductIDs[0]: model.Order{
			Model:      gorm.Model{ID: uint(mockProductIDs[0])},
			Status:     consts.OnProgress,
			TotalPrice: mockProductPrices[mockProductIDs[0]] * float64(mockProductQuantity[mockProductIDs[0]]),
			OrderDetails: []model.OrderDetail{
				model.OrderDetail{
					Model:     gorm.Model{ID: uint(mockOrderDetailIDs[0])},
					OrderID:   uint(mockOrderIDs[0]),
					ProductID: uint(mockProductIDs[0]),
					Quantity:  mockProductQuantity[0],
				},
			},
		},
	}
	mockResult := []*model.OrderStruct{
		&model.OrderStruct{
			ID:         mockOrderIDs[0],
			Status:     consts.OnProgress,
			TotalPrice: mockProductPrices[mockProductIDs[0]] * float64(mockProductQuantity[mockProductIDs[0]]),
			Orders: []model.OrderDetailStruct{
				model.OrderDetailStruct{
					ID:        mockOrderDetailIDs[0],
					OrderID:   mockOrderIDs[0],
					ProductID: mockProductIDs[0],
					Quantity:  mockProductQuantity[0],
				},
			},
		},
	}

	type args struct {
		limit  int
		offset int
	}
	type fields struct {
		DatabaseRepository repository.Database
		MidtransRepository repository.Midtrans
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		result  []*model.OrderStruct
		wantErr bool
		mocks   func()
	}{
		{
			name: "TestCase1-SelectAllOrders return error",
			args: args{
				limit:  limit,
				offset: offset,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  nil,
			wantErr: true,
			mocks: func() {
				mockDBRepo.EXPECT().SelectAllOrders(limit, offset).Return(nil, expectedError)
			},
		},
		{
			name: "TestCase1-SelectAllOrders success return zero result",
			args: args{
				limit:  limit,
				offset: offset,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  nil,
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().SelectAllOrders(limit, offset).Return([]*model.Order{}, nil)
			},
		},
		{
			name: "TestCase1-SelectAllOrders success",
			args: args{
				limit:  limit,
				offset: offset,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  mockResult,
			wantErr: false,
			mocks: func() {
				mockOrder := mockOrders[mockProductIDs[0]]
				mockDBRepo.EXPECT().SelectAllOrders(limit, offset).Return([]*model.Order{
					&mockOrder,
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			res, err := storeUC.SelectAllOrders(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectAllOrders() got error, wantErr is %v", tt.wantErr)
			} else if !reflect.DeepEqual(res, tt.result) {
				t.Errorf("SelectAllOrders() got %v, want %v", res, tt.result)
			}
		})
	}
}

func TestUsecase_SelectOrdersByIDs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)
	mockOrderIDs := []int{1}
	mockOrderDetailIDs := []int{1}

	mockProductIDs := []int{1}
	mockProductQuantity := map[int]int{
		mockProductIDs[0]: 5,
	}
	mockProductPrices := map[int]float64{
		mockProductIDs[0]: 50000,
	}
	mockOrders := map[int]model.Order{
		mockProductIDs[0]: model.Order{
			Model:      gorm.Model{ID: uint(mockProductIDs[0])},
			Status:     consts.OnProgress,
			TotalPrice: mockProductPrices[mockProductIDs[0]] * float64(mockProductQuantity[mockProductIDs[0]]),
			OrderDetails: []model.OrderDetail{
				model.OrderDetail{
					Model:     gorm.Model{ID: uint(mockOrderDetailIDs[0])},
					OrderID:   uint(mockOrderIDs[0]),
					ProductID: uint(mockProductIDs[0]),
					Quantity:  mockProductQuantity[0],
				},
			},
		},
	}
	mockResult := []*model.OrderStruct{
		&model.OrderStruct{
			ID:         mockOrderIDs[0],
			Status:     consts.OnProgress,
			TotalPrice: mockProductPrices[mockProductIDs[0]] * float64(mockProductQuantity[mockProductIDs[0]]),
			Orders: []model.OrderDetailStruct{
				model.OrderDetailStruct{
					ID:        mockOrderDetailIDs[0],
					OrderID:   mockOrderIDs[0],
					ProductID: mockProductIDs[0],
					Quantity:  mockProductQuantity[0],
				},
			},
		},
	}

	type args struct {
		orderIDs []int
	}
	type fields struct {
		DatabaseRepository repository.Database
		MidtransRepository repository.Midtrans
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		result  []*model.OrderStruct
		wantErr bool
		mocks   func()
	}{
		{
			name: "TestCase1-SelectOrdersByIDs return error",
			args: args{
				orderIDs: mockOrderIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  nil,
			wantErr: true,
			mocks: func() {
				mockDBRepo.EXPECT().SelectOrdersByIDs(mockProductIDs).Return(nil, expectedError)
			},
		},
		{
			name: "TestCase1-SelectOrdersByIDs success return zero result",
			args: args{
				orderIDs: mockOrderIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  nil,
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().SelectOrdersByIDs(mockOrderIDs).Return([]*model.Order{}, nil)
			},
		},
		{
			name: "TestCase1-SelectOrdersByIDs success",
			args: args{
				orderIDs: mockOrderIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  mockResult,
			wantErr: false,
			mocks: func() {
				mockOrder := mockOrders[mockProductIDs[0]]
				mockDBRepo.EXPECT().SelectOrdersByIDs(mockOrderIDs).Return([]*model.Order{
					&mockOrder,
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			res, err := storeUC.SelectOrdersByIDs(tt.args.orderIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectOrdersByIDs() got error, wantErr is %v", tt.wantErr)
			} else if !reflect.DeepEqual(res, tt.result) {
				t.Errorf("SelectOrdersByIDs() got %v, want %v", res, tt.result)
			}
		})
	}
}

func TestUsecase_CreateOrders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)

	mockProductIDs := []int{1}
	mockProductQuantity := map[int]int{
		mockProductIDs[0]: 5,
	}
	mockProductPrices := map[int]float64{
		mockProductIDs[0]: 50000,
	}
	mockRequest := []*model.OrderStruct{
		&model.OrderStruct{
			Status:     consts.OnProgress,
			TotalPrice: mockProductPrices[mockProductIDs[0]] * float64(mockProductQuantity[mockProductIDs[0]]),
			Orders: []model.OrderDetailStruct{
				model.OrderDetailStruct{
					ProductID: mockProductIDs[0],
					Quantity:  mockProductQuantity[mockProductIDs[0]],
				},
			},
		},
	}
	mockDBOrder := map[int]model.Order{
		mockProductIDs[0]: model.Order{
			Status:     consts.OnProgress,
			TotalPrice: mockProductPrices[mockProductIDs[0]] * float64(mockProductQuantity[mockProductIDs[0]]),
			OrderDetails: []model.OrderDetail{
				model.OrderDetail{
					ProductID: uint(mockProductIDs[0]),
					Quantity:  mockProductQuantity[mockProductIDs[0]],
				},
			},
		},
	}
	mockDBProduct := []*model.Product{
		&model.Product{
			Model:       gorm.Model{ID: uint(mockProductIDs[0])},
			ProductName: "product_test",
			Price:       mockProductPrices[mockProductIDs[0]],
			Quantity:    mockProductQuantity[mockProductIDs[0]],
		},
	}

	mockDBProductMap := map[int]*model.Product{
		mockProductIDs[0]: mockDBProduct[0],
	}
	mockMidtransTD := midtrans.TransactionDetails{
		OrderID:  strconv.Itoa(int(mockDBOrder[mockProductIDs[0]].ID)),
		GrossAmt: int64(mockDBOrder[int(mockProductIDs[0])].TotalPrice),
	}
	mockMidtransItems := &[]midtrans.ItemDetails{}
	for _, orderDetail := range mockDBOrder[mockProductIDs[0]].OrderDetails {
		tempMockDBProduct := mockDBProductMap[int(orderDetail.ProductID)]

		*mockMidtransItems = append(*mockMidtransItems, midtrans.ItemDetails{
			ID:    strconv.Itoa(int(tempMockDBProduct.ID)),
			Name:  tempMockDBProduct.ProductName,
			Price: int64(tempMockDBProduct.Price),
			Qty:   int32(orderDetail.Quantity),
		})
	}

	type args struct {
		orderStructs []*model.OrderStruct
	}
	type fields struct {
		DatabaseRepository repository.Database
		MidtransRepository repository.Midtrans
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		wantErr bool
		mocks   func()
	}{
		{
			name: "TestCase1-CreateOrders request body is empty",
			args: args{
				orderStructs: []*model.OrderStruct{},
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: false,
			mocks:   func() {},
		},
		{
			name: "TestCase1-CreateOrders SelectProductsByIDs return error",
			args: args{
				orderStructs: mockRequest,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: true,
			mocks: func() {
				mockDBRepo.EXPECT().SelectProductsByIDs([]int{mockProductIDs[0]}).Return(nil, expectedError)
			},
		},
		{
			name: "TestCase1-CreateOrders return error",
			args: args{
				orderStructs: mockRequest,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: true,
			mocks: func() {
				mockOrder := mockDBOrder[mockProductIDs[0]]

				mockDBRepo.EXPECT().SelectProductsByIDs([]int{mockProductIDs[0]}).Return(mockDBProduct, nil)
				mockDBRepo.EXPECT().CreateOrders([]*model.Order{&mockOrder}).Return(nil, expectedError)
			},
		},
		{
			name: "TestCase1-CreateOrders success",
			args: args{
				orderStructs: mockRequest,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: false,
			mocks: func() {
				mockOrder := mockDBOrder[mockProductIDs[0]]

				mockDBRepo.EXPECT().SelectProductsByIDs([]int{mockProductIDs[0]}).Return(mockDBProduct, nil)
				mockDBRepo.EXPECT().CreateOrders([]*model.Order{&mockOrder}).Return([]*model.Order{&mockOrder}, nil)

				mockCoreApiReq := &coreapi.ChargeReq{
					PaymentType:        coreapi.PaymentTypeGopay,
					TransactionDetails: mockMidtransTD,
					Items:              mockMidtransItems,
				}
				mockMidtransRepo.EXPECT().ChargeTransaction(mockCoreApiReq)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			err := storeUC.CreateOrders(tt.args.orderStructs)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrders() got error, wantErr is %v", tt.wantErr)
			}
		})
	}
}

func TestUsecase_DeleteOrdersByIDs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)
	mockOrderIDs := []int{1}

	type args struct {
		orderIDs []int
	}
	type fields struct {
		DatabaseRepository repository.Database
		MidtransRepository repository.Midtrans
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		wantErr bool
		mocks   func()
	}{
		{
			name: "TestCase1-DeleteOrdersByIDs return error",
			args: args{
				orderIDs: mockOrderIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: true,
			mocks: func() {
				mockDBRepo.EXPECT().DeleteOrdersByIDs(mockOrderIDs).Return(expectedError)
			},
		},
		{
			name: "TestCase1-DeleteOrdersByIDs success",
			args: args{
				orderIDs: mockOrderIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().DeleteOrdersByIDs(mockOrderIDs).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			err := storeUC.DeleteOrdersByIDs(tt.args.orderIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteOrdersByIDs() got error, wantErr is %v", tt.wantErr)
			}
		})
	}
}
