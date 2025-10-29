package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProductsHandler))
	mux.Handle("POST /createproduct", http.HandlerFunc(handlers.CreateProductHandler))

	fmt.Println("Starting server on :8080")
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
