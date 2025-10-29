package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	util.SendData(w, newProduct, http.StatusCreated)

}
