package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /rahim",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)
	mux.Handle("GET /abin",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)
	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProductsHandler),
		),
	)
	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProductHandler),
		),
	)
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		),
	)
}
