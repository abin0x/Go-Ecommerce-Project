package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first name"`
	LastName    string `json:"last name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner string `json:"is_shop_owner"`
}
