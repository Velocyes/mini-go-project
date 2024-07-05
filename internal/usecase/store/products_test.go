package store

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Velocyes/mini-go-project/internal/model"
	"github.com/Velocyes/mini-go-project/internal/repository"
	mockRepo "github.com/Velocyes/mini-go-project/mocks/repository"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

var (
	expectedError = errors.New("expected error")
)

func TestUsecase_SelectAllProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)
	limit, offset := -1, -1
	mockProductIDs := []int{1}
	mockResult := []*model.Product{
		&model.Product{
			Model:       gorm.Model{ID: uint(mockProductIDs[0])},
			ProductName: "product_name",
			Price:       50000.0,
			Quantity:    10,
		},
	}
	result := []*model.ProductStruct{
		&model.ProductStruct{
			ID:          mockProductIDs[0],
			ProductName: "product_name",
			Price:       50000.0,
			Quantity:    10,
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
		result  []*model.ProductStruct
		wantErr bool
		mocks   func()
	}{
		{
			name: "TestCase1-SelectAllProducts return error",
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
				mockDBRepo.EXPECT().SelectAllProducts(limit, offset).Return(nil, expectedError)
			},
		},
		{
			name: "TestCase1-SelectAllProducts success return zero result",
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
				mockDBRepo.EXPECT().SelectAllProducts(limit, offset).Return([]*model.Product{}, nil)
			},
		},
		{
			name: "TestCase1-SelectAllProducts success",
			args: args{
				limit:  limit,
				offset: offset,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  result,
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().SelectAllProducts(limit, offset).Return(mockResult, nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			res, err := storeUC.SelectAllProducts(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectAllProducts() got error, wantErr is %v", tt.wantErr)
			} else if !reflect.DeepEqual(res, tt.result) {
				t.Errorf("SelectAllProducts() got %v, want %v", res, tt.result)
			}
		})

	}
}

func TestUsecase_SelectProductsByIDs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)
	mockProductIDs := []int{1}
	mockResult := []*model.Product{
		&model.Product{
			Model:       gorm.Model{ID: uint(mockProductIDs[0])},
			ProductName: "product_name",
			Price:       50000.0,
			Quantity:    10,
		},
	}
	result := []*model.ProductStruct{
		&model.ProductStruct{
			ID:          mockProductIDs[0],
			ProductName: "product_name",
			Price:       50000.0,
			Quantity:    10,
		},
	}

	type args struct {
		productIDs []int
	}
	type fields struct {
		DatabaseRepository repository.Database
		MidtransRepository repository.Midtrans
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		result  []*model.ProductStruct
		wantErr bool
		mocks   func()
	}{
		{
			name: "TestCase1-SelectProductsByIDs return error",
			args: args{
				productIDs: mockProductIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  nil,
			wantErr: true,
			mocks: func() {
				mockDBRepo.EXPECT().SelectProductsByIDs(mockProductIDs).Return(nil, expectedError)
			},
		},
		{
			name: "TestCase1-SelectProductsByIDs success return zero result",
			args: args{
				productIDs: mockProductIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  nil,
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().SelectProductsByIDs(mockProductIDs).Return([]*model.Product{}, nil)
			},
		},
		{
			name: "TestCase1-SelectProductsByIDs success",
			args: args{
				productIDs: mockProductIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			result:  result,
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().SelectProductsByIDs(mockProductIDs).Return(mockResult, nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			res, err := storeUC.SelectProductsByIDs(tt.args.productIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectProductsByIDs() got error, wantErr is %v", tt.wantErr)
			} else if !reflect.DeepEqual(res, tt.result) {
				t.Errorf("SelectProductsByIDs() got %v, want %v", res, tt.result)
			}
		})

	}
}

func TestUsecase_UpdateProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)
	mockProductIDs := []int{1}
	mockRequest := []*model.ProductStruct{
		&model.ProductStruct{
			ID:          mockProductIDs[0],
			ProductName: "product_name",
			Price:       50000.0,
			Quantity:    10,
		},
	}
	mockDBRequest := []*model.Product{
		&model.Product{
			Model:       gorm.Model{ID: uint(mockProductIDs[0])},
			ProductName: "product_name",
			Price:       50000.0,
			Quantity:    10,
		},
	}

	type args struct {
		productStructs []*model.ProductStruct
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
			name: "TestCase1-UpdateProducts request body is empty",
			args: args{
				productStructs: []*model.ProductStruct{},
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: false,
			mocks:   func() {},
		},
		{
			name: "TestCase1-UpdateProducts return error",
			args: args{
				productStructs: mockRequest,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: true,
			mocks: func() {
				mockDBRepo.EXPECT().UpdateProducts(mockDBRequest).Return(expectedError)
			},
		},
		{
			name: "TestCase1-UpdateProducts success",
			args: args{
				productStructs: mockRequest,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().UpdateProducts(mockDBRequest).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			err := storeUC.UpdateProducts(tt.args.productStructs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProducts() got error, wantErr is %v", tt.wantErr)
			}
		})

	}
}

func TestUsecase_CreateProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)
	mockRequest := []*model.ProductStruct{
		&model.ProductStruct{
			ProductName: "product_name",
			Price:       50000.0,
			Quantity:    10,
		},
	}
	mockDBRequest := []*model.Product{
		&model.Product{
			ProductName: "product_name",
			Price:       50000.0,
			Quantity:    10,
		},
	}

	type args struct {
		productStructs []*model.ProductStruct
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
			name: "TestCase1-CreateProducts request body is empty",
			args: args{
				productStructs: []*model.ProductStruct{},
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: false,
			mocks:   func() {},
		},
		{
			name: "TestCase1-CreateProducts return error",
			args: args{
				productStructs: mockRequest,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: true,
			mocks: func() {
				mockDBRepo.EXPECT().CreateProducts(mockDBRequest).Return(expectedError)
			},
		},
		{
			name: "TestCase1-CreateProducts success",
			args: args{
				productStructs: mockRequest,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().CreateProducts(mockDBRequest).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			err := storeUC.CreateProducts(tt.args.productStructs)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProducts() got error, wantErr is %v", tt.wantErr)
			}
		})

	}
}

func TestUsecase_DeleteProductsByIDs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBRepo := mockRepo.NewMockDatabase(mockCtrl)
	mockMidtransRepo := mockRepo.NewMockMidtrans(mockCtrl)
	mockProductIDs := []int{1}

	type args struct {
		productIDs []int
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
			name: "TestCase1-DeleteProductsByIDs return error",
			args: args{
				productIDs: mockProductIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: true,
			mocks: func() {
				mockDBRepo.EXPECT().DeleteProductsByIDs(mockProductIDs).Return(expectedError)
			},
		},
		{
			name: "TestCase1-DeleteProductsByIDs success",
			args: args{
				productIDs: mockProductIDs,
			},
			fields: fields{
				DatabaseRepository: mockDBRepo,
				MidtransRepository: mockMidtransRepo,
			},
			wantErr: false,
			mocks: func() {
				mockDBRepo.EXPECT().DeleteProductsByIDs(mockProductIDs).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mocks()

		t.Run(tt.name, func(t *testing.T) {
			storeUC := InitStoreUsecase(tt.fields.DatabaseRepository, tt.fields.MidtransRepository)
			err := storeUC.DeleteProductsByIDs(tt.args.productIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteProductsByIDs() got error, wantErr is %v", tt.wantErr)
			}
		})

	}
}
