package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"otumian-empire/go-ecom/src/config"
	"otumian-empire/go-ecom/src/web"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Go E-com api")

	env, err := config.GetEnvirons()
	if err != nil {
		log.Println("An error occurred reading env")
		log.Fatalln(err)
	}

	connectionString := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v", env.Database, env.DatabaseUsername, env.DatabasePassword, env.DatabaseHost, env.DatabasePort, env.DatabaseName, env.SslMode)

	db, err := sql.Open(env.DatabaseDriverName, connectionString)
	if err != nil {
		log.Println("Database connection error")
		log.Fatalln(err)
	}

	defer db.Close()

	if err := recover(); err != nil {
		log.Println("SERVER_RECOVER_FROM_ERROR")
		log.Println(err)
	}

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"isSuccessful": true,
			"message":      "Hello there... This is a test endpoint",
		})
	})

	// api v1
	web.NewHandler(router.Group("/api/v1"), db)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", env.ServerPort), router))
}
