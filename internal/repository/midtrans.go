package repository

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

//go:generate mockgen -source=midtrans.go -destination=./../../mocks/repository/mock_midtrans.go

type Midtrans interface {
	GetCoreAPIClient() (*coreapi.Client)
	
	// ChargeTransaction : Do `/charge` API request to Midtrans Core API return `coreapi.Response` with `coreapi.ChargeReq`
	ChargeTransaction(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error)
}
