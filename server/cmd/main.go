package main

import (
	"fmt"

	"github.com/lahnasti/clientServerApp/server/routes"
)

func main() {
	r := routes.UserRoutes()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		panic(err)
	}
}