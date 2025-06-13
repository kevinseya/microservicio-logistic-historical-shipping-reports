package main

import (
	"fmt"
	"historical-shipping-reports/config"
	"historical-shipping-reports/database"
	"historical-shipping-reports/graph"
	"historical-shipping-reports/graph/generated"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()
	database.ConnectDB()

	router := gin.Default()

	// Redirect to GraphiQL
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/graphiql")
	})
	router.GET("/graphiql", func(c *gin.Context) {
		handler := playground.Handler("GraphQL playground", "/graphql")
		handler.ServeHTTP(c.Writer, c.Request)
	})

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})

	// Endpoint to GRAPHQL
	router.POST("/graphql", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		handler := handler.NewDefaultServer(schema)
		handler.ServeHTTP(c.Writer, c.Request)
	})

	port := config.AppConfig.Server.Port
	log.Printf("Microservice running on port: %s", port)

	// Run server
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error to initialize server: %v", err)
	}
}
