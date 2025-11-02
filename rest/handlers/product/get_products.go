package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a product by its ID
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Request body")
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product Not Found")
		return
	}
	util.SendData(w, http.StatusOK, product)

}
