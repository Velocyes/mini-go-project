package store

import (
	"github.com/Velocyes/mini-go-project/internal/repository"
)

type StoreUsecase struct {
	DatabaseRepository repository.Database
	MidtransRepository repository.Midtrans
}

func InitStoreUsecase(dbRepo repository.Database, midtransRepo repository.Midtrans) *StoreUsecase {
	return &StoreUsecase{
		DatabaseRepository: dbRepo,
		MidtransRepository: midtransRepo,
	}
}
