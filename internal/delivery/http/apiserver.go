package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"

	"github.com/Velocyes/mini-go-project/internal/consts"
	"github.com/Velocyes/mini-go-project/internal/model"
	"github.com/Velocyes/mini-go-project/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/tboehle/go-response-utils"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

type APIServer struct {
	Port string

	StoreUsecase usecase.StoreUsecase
}

func NewAPIServer(cfg *model.Config, storeUsecase usecase.StoreUsecase) (apiServer *APIServer, err error) {
	if cfg == nil {
		return nil, consts.ErrNilConfig
	}

	return &APIServer{
		Port:         cfg.Server.Port,
		StoreUsecase: storeUsecase,
	}, nil
}

func panicRecovery(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				log.Printf("[HTTP][Panic] Recovering from err %v\n %s", err, buf)
				response.WithError(w, r, http.StatusInternalServerError, errors.New("server got panic"))
			}
		}()

		f(w, r)
	}
}

func (a *APIServer) initNgrokForwarder(ctx context.Context) (ngrok.Forwarder, error) {
	backendUrl, err := url.Parse(fmt.Sprintf("http://localhost:%s", a.Port))
	if err != nil {
		return nil, err
	}

	ngrokSession, err := ngrok.Connect(ctx,
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return nil, err
	}

	ngrokForwarder, err := ngrokSession.ListenAndForward(ctx,
		backendUrl,
		config.HTTPEndpoint(),
	)
	if err != nil {
		return nil, err
	}

	log.Printf("Ngrok URL : %s", ngrokForwarder.URL())
	return ngrokForwarder, nil
}

func (a *APIServer) Run() error {
	ctx := context.Background()

	router := mux.NewRouter()

	// Products
	router.HandleFunc("/products", panicRecovery(a.getAllProducts)).Methods("GET")
	router.HandleFunc("/products/{ids}", panicRecovery(a.getProductsByIDs)).Methods("GET")
	router.HandleFunc("/products", panicRecovery(a.insertNewProducts)).Methods("POST")
	router.HandleFunc("/products", panicRecovery(a.alterProducts)).Methods("PUT")
	router.HandleFunc("/products/{ids}", panicRecovery(a.deleteProductsByIDs)).Methods("DELETE")

	// Orders
	router.HandleFunc("/orders", panicRecovery(a.getAllOrders)).Methods("GET")
	router.HandleFunc("/orders/{ids}", panicRecovery(a.getOrdersByIDs)).Methods("GET")
	router.HandleFunc("/orders", panicRecovery(a.insertNewOrders)).Methods("POST")
	router.HandleFunc("/orders/{ids}", panicRecovery(a.deleteOrdersByIDs)).Methods("DELETE")

	ngrokForwarder, err := a.initNgrokForwarder(ctx)
	if err != nil {
		return err
	}

	go func() error {
		for {
			err = ngrokForwarder.Wait()
			if err == nil {
				return nil
			}
		}
	}()

	log.Printf("[HTTP] Server starting on port %s", a.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", a.Port), router)
}
