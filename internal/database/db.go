package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		"containers-us-west-116.railway.app",
		"postgres",
		"QSrvkiuqLxo7eyD5UNTi",
		"railway",
		"5703")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // nao pode usar em producao
	}

	return db, err
}
