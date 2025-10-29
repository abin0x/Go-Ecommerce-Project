package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a product by its ID
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
