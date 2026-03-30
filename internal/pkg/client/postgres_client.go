package client

import (
	"fmt"
	"log"
	"task_planner_application/internal/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL")
	return db, nil
}
