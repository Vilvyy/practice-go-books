package main

import (
	"fmt"
	"os"

	"practice-go-books/configs"
	"practice-go-books/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		os.Exit(1)
	}
	envFile := fmt.Sprintf("%s/.env", wd)
	configs.LoadEnv(envFile)

	router := gin.Default()
	routes.Setup(router, nil)
	router.Run(fmt.Sprintf(":%s", viper.Get("PORT")))
}
