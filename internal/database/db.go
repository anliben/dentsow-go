package database

import (
	"fiber/pkg/common/models"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		"viaduct.proxy.rlwy.net",
		"postgres",
		"bFFaA1bDGCaEE5Ga2dG413b64Dg3g6f6",
		"railway",
		"44375")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return nil, err
	}

	// AutoMigrate(db)

	return db, err
}

func AutoMigrate(ctx *gorm.DB) {
	ctx.Debug().AutoMigrate(
		&models.User{},
		&models.Budget{},
		&models.Customer{},
		&models.Data{},
		&models.Files{},
		&models.Groups{},
		&models.Procedure{},
		&models.ProposedValue{},
		&models.Tooth{},
	)
}
