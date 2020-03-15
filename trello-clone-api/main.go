package main

import (
	"log"
	"os"
)

var mariaDB *MariaDB

func main() {
	router := SetUpRouter()

	mariaDB = new(MariaDB)

	if err := mariaDB.connect(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mariaDB.close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := mariaDB.migrate(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	router.Run(":" + port)
}
