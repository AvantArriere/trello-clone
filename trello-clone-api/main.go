package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var mariaDB *MariaDB

func main() {
	router := SetUpRouter()

	mariaDB = new(MariaDB)

	// mariadb connect
	if err := mariaDB.connect(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mariaDB.close(); err != nil {
			log.Fatal(err)
		}
	}()
	// mariadb migrate
	if err := mariaDB.migrate(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	router.Run(":" + port)
}
