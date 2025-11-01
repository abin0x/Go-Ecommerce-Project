package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	headerArr := strings.Split(header, " ")
	if len(headerArr) != 2 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	accessToken := headerArr[1]
	tokenParts := strings.Split(accessToken, ".")
	if len(tokenParts) != 3 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	signature := tokenParts[2]

	message := jwtHeader + "." + jwtPayload
	cnf := config.GetConfig()

	byteArrawSecret := []byte(cnf.JwtSecret)
	byteArrayMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrawSecret)
	h.Write(byteArrayMessage)

	hash := h.Sum(nil)
	newSignature := base64UrlEncode(hash)

	if newSignature != signature {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	createdProduct := database.Store(newProduct)
	util.SendData(w, createdProduct, http.StatusCreated)

}
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
