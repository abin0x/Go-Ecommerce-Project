package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a product by its ID
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.Get(pId)

	if product == nil {
		util.SendError(w, 404, "product not found")
		return
	}
	util.SendData(w, product, 200)

}
