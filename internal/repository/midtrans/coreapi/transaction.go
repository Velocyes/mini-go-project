package coreapi

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (c *CoreAPI) ChargeTransaction(req *coreapi.ChargeReq) (coreApiRes *coreapi.ChargeResponse, coreApiErr *midtrans.Error) {
	return c.client.ChargeTransaction(req)
}