package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		"containers-us-west-23.railway.app",
		"postgres",
		"Woo2oJDYXR4CM7luKpmW",
		"railway",
		"5508")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // nao pode usar em producao
	}

	return db, err
}
