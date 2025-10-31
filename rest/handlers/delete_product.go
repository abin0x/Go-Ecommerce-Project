package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "please give me a valid product id", 400)
		return
	}

	database.Delete(pId)

	util.SendData(w, "successfully deleted product", 201)
}
