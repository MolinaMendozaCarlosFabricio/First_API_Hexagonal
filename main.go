package main

import (
	_ "fmt"
	_ "log"
	"net/http"

	"example2.com/m/Products/infraestructure/routes"
	userRoutes "example2.com/m/Users/infraestructure/routes"
	_ "example2.com/m/core"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	r := gin.Default()
	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:4200"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
        AllowCredentials: true,
        Debug:            true,
    })
	
	routes.ProductRoutes(r)
	userRoutes.UserRoutes(r)

	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}