package handlers

import (
	"net/http"

	"ecommerce/database"
	"ecommerce/util"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200)
}
