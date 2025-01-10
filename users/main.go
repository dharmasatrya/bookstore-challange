package main

import (
	"GC1P3-order/routes"
	"os"
)

func main() {
	router := routes.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Logger.Fatal(router.Start(":" + port))

}
