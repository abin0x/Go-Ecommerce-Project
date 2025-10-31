package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
	"fmt"
)

func main() {
	cnf := config.GetConfig()
	fmt.Println(cnf.Version)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.HttpPort)
	cmd.Serve()
}
