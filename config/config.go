package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
	Db          *DBConfig
}

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables: ", err)
		os.Exit(1)
	}

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

	httpPortInt, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("HTTP_PORT must be a number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB Host is required")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB PORT is required")
		os.Exit(1)
	}
	dbPrt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("DB Port must be number")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB NAME is required")
		os.Exit(1)
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB USER is required")
		os.Exit(1)
	}
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		fmt.Println("DB PASS is required")
		os.Exit(1)
	}
	enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")
	enbleSSLMode, err := strconv.ParseBool(enableSslMode)
	if err != nil {
		fmt.Println("invalid enable ssl mode value", err)
		os.Exit(1)
	}
	// if enableSslMode == "" {
	// 	fmt.Println("DB PASS is required")
	// 	os.Exit(1)
	// }

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          int(dbPrt),
		Name:          dbName,
		User:          dbUser,
		Password:      dbPass,
		EnableSSLMODE: enbleSSLMode,
	}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPortInt,
		JwtSecret:   jwtSecretKey,
		Db:          dbConfig,
	}

}

func GetConfig() *Config {

	if configurations == nil {
		loadConfig()
	}
	return configurations
}
