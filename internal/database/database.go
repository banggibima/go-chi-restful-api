package database

import (
	"fmt"
	"log"

	"github.com/banggibima/go-chi-restful-api/internal/config"
	"github.com/banggibima/go-chi-restful-api/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() (*gorm.DB, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
		return nil, err
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
		return nil, err
	}

	err = AutoMigrate(db)
	if err != nil {
		log.Fatalf("error performing auto migration: %v", err)
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatalf("error auto-migrating User table: %v", err)
		return err
	}

	return nil
}
