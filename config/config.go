package config

import (
	"fmt"
	"os"
	"strconv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
}

func LoadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	configurations := Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}

}
