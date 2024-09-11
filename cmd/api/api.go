package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/codepnw/shopping-cart/service/cart"
	"github.com/codepnw/shopping-cart/service/order"
	"github.com/codepnw/shopping-cart/service/product"
	"github.com/codepnw/shopping-cart/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// User Routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	// Product Routes
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)
	// Cart Routes
	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("server listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
