package main

import (
	"log"
	"net/http"
	"task_planner_application/internal/api/router"
	"task_planner_application/internal/pkg/client"
	"task_planner_application/internal/pkg/config"
	"task_planner_application/internal/task/handlers"
	"task_planner_application/internal/task/infra"
	"task_planner_application/internal/task/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Initialize Database Client (PostgreSQL)
	db, err := client.NewPostgresClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 3. Initialize Redis Client
	redisClient, err := client.NewRedisClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}

	// 4. Initialize Dependency Injection
	taskRepo := infra.NewTaskRepo(db)
	taskService := services.NewTaskService(taskRepo, redisClient)
	taskHandler := handlers.NewTaskHandler(taskService)

	// 5. Initialize API Router
	app := router.NewChiRouter(taskHandler)

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting application on port %s...", port)
	if err := http.ListenAndServe(":"+port, app); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
