package main

import (
	"database/sql"
	"fmt"
	"os"

	"practice-go-books/configs"
	"practice-go-books/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/lib/pq" // To register the driver.go mod t
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		os.Exit(1)
	}
	envFile := fmt.Sprintf("%s/.env", wd)
	configs.LoadEnv(envFile)

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", viper.GetString("DATABASE_HOST"), viper.GetString("DATABASE_PORT"), viper.GetString("DATABASE_USER"), viper.GetString("DATABASE_PASSWORD"), viper.GetString("DATABASE_NAME"), viper.GetString("DATABASE_SSLMODE")))
	if err != nil {
		fmt.Printf("Error Connecting to database: %v\n", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
		os.Exit(1)
	}

	router := gin.Default()
	routes.Setup(router, db)

	router.Run(fmt.Sprintf(":%s", viper.Get("PORT")))
}
