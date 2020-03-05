package main

import "os"

func main() {
	router := SetUpRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	router.Run(":" + port)
}
