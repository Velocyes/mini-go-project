package main

import (
	"log"

	"github.com/Velocyes/mini-go-project/internal/config"
	"github.com/Velocyes/mini-go-project/internal/delivery/http"
	"github.com/Velocyes/mini-go-project/internal/model"
	"github.com/Velocyes/mini-go-project/internal/repository"
	mySQL "github.com/Velocyes/mini-go-project/internal/repository/database/mysql"
	midtransCoreApi "github.com/Velocyes/mini-go-project/internal/repository/midtrans/coreapi"
	"github.com/Velocyes/mini-go-project/internal/usecase"
	storeUC "github.com/Velocyes/mini-go-project/internal/usecase/store"
)

func main() {
	var (
		cfg *model.Config

		dbRepository repository.Database
		storeUsecase usecase.StoreUsecase
	)

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("[HTTP][Config] Initialization Failed, return err : %v", err)
	}

	dbRepository, err = mySQL.InitMySQL(cfg)
	if err != nil {
		log.Fatalf("[HTTP][Repository][Database][MySQL] Initialization Failed, return err : %v", err)
	}

	midtransRepsitory, err := midtransCoreApi.InitCoreAPIMidtrans(cfg)
	if err != nil {
		log.Fatalf("[HTTP][Repository][Midtrans][CoreAPI] Initialization Failed, return err : %v", err)
	}

	storeUsecase = storeUC.InitStoreUsecase(dbRepository, midtransRepsitory)

	apiServer, err := http.NewAPIServer(cfg, storeUsecase)
	if err := apiServer.Run(); err != nil {
		log.Printf("[HTTP] Error starting the server, return err : %v\n", err)
	}
}
