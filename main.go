package main

import (
	_ "fmt"
	_ "log"

	"example2.com/m/Products/infraestructure/routes"
	userRoutes"example2.com/m/Users/infraestructure/routes"
	_ "example2.com/m/core"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.ProductRoutes(r)
	userRoutes.UserRoutes(r)
	r.Run()
	/*
	db, err := core.MySQLConection()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Conectaco con la base de datos");

	r := gin.Default()
	routes.ProductRoutes(r)
	r.Run()

	defer db.Close()
*/
	
}