package config

import (
	"backend-test/src/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(conn DBConfig) (*gorm.DB, error) {

	username := conn.Username
	password := conn.Password
	host := conn.Host
	port := conn.Port
	database := conn.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.Product{})
	DBSeed(db)

	return db, err
}
