package main

import (
	"fmt"
	"go-rest/database"
	"go-rest/routes"
)

func main() {
	database.DbConnection()
	fmt.Println("Starting server...")
	routes.HandleRequest()
}
