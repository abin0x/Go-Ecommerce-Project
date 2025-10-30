package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai, middleware.CorsWithPreflight)

	mux := http.NewServeMux()
	InitRoutes(mux, manager)

	// globalRouter := middleware.CorsWithPreflight(mux)
	wrappedMux := manager.With(mux)
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
