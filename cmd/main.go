package main

import (
	"fmt"
	"food-api/internal/application/services"
	"food-api/internal/infrastructure/config"
	"food-api/internal/infrastructure/database/cassandra"
	"food-api/internal/interfaces/http/controllers"
	"food-api/internal/interfaces/http/middleware"
	"food-api/internal/interfaces/http/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func main() {
	fmt.Println("Starting server...")

	config.LoadConfig()

	cluster := CreateCluster(config.AppConfig.Cassandra)
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("unable to create session: %v", err)
	}
	defer session.Close()

	cassandra.SeedData(session)

	repo := cassandra.NewCategoryRepository(session)
	service := services.NewCategoryService(repo)
	controller := controllers.NewCategoryController(service)

	router := gin.Default()
	router.Use(middleware.RateLimitMiddleware())

	routes.RegisterCategoryRoutes(router, controller)

	router.Run(":" + config.AppConfig.Server.Port)
}

func CreateCluster(config config.CassandraConfig) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.Quorum
	return cluster
}
