package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()
	InitRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
