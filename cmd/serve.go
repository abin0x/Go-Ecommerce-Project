package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	// globalRouter := middleware.CorsWithPreflight(mux)
	// wrappedMux := manager.WrapMux(
	// 	mux,
	// 	middleware.Logger,
	// 	middleware.Hudai,
	// 	middleware.CorsWithPreflight,
	// )
	wrappedMux := manager.WrapMux(nil, mux)
	InitRoutes(mux, manager)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
