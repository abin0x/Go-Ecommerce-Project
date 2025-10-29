package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	mux := http.NewServeMux()
	mux.Handle("GET /rahim", manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Logger,
		middleware.Hudai,
	))
	mux.Handle("GET /abin", manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Hudai,
		middleware.Logger,
	))
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProductsHandler),
		middleware.Logger,
		middleware.Hudai,
	))
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProductHandler),
		middleware.Logger,
		middleware.Hudai,
	))
	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handlers.GetProductById),
		middleware.Logger,
		middleware.Hudai,
	))

	fmt.Println("Starting server on :8080")
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
