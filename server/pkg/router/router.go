package router

import (
	"github.com/gorilla/mux"
	"github.com/mjmichael73/go-simple-http-server-client/server/pkg/handler"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", handler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handler.GetProduct).Methods("GET")
	router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handler.DeleteProduct).Methods("DELETE")
	return router
}
