package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /abin", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))

	mux.Handle("GET /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProductsHandler))))
	mux.Handle("POST /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.CreateProductHandler))))
	mux.Handle("GET /products/{id}", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProductById))))

	fmt.Println("Starting server on :8080")
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
