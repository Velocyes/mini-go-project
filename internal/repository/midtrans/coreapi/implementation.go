package coreapi

import (
	"github.com/Velocyes/mini-go-project/internal/consts"
	"github.com/Velocyes/mini-go-project/internal/model"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type CoreAPI struct {
	client coreapi.Client
}

func InitCoreAPIMidtrans(cfg *model.Config) (*CoreAPI, error) {
	if cfg == nil {
		return nil, consts.ErrNilConfig
	}

	coreApiClient := coreapi.Client{}
	coreApiClient.New(cfg.Midtrans.ServerKey, midtrans.Sandbox)

	return &CoreAPI{
		client: coreApiClient,
	}, nil
}

func (c *CoreAPI) GetCoreAPIClient() (*coreapi.Client) {
	return &c.client
}
