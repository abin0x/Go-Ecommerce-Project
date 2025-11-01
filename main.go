package main

import (
	"ecommerce/cmd"
)

// "ecommerce/cmd"

func main() {
	cmd.Serve()
	// jwt, err := util.CreateJwt("my-secret", util.Payload{
	// 	Sub:         45,
	// 	FirstName:   "Mahmudul",
	// 	LastName:    "hasan",
	// 	Email:       "abin@gmail.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(jwt)
}
