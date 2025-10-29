package handlers

import (
	"net/http"

	"ecommerce/product"
	"ecommerce/util"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, product.ProductList, 200)
}
