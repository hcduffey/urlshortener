package main

import (
	"log"
)

func main() {
	initializeRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
