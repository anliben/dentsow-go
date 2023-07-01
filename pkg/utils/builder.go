package utils

import (
	"gorm.io/gorm"
)

func Builder(query string, value interface{}) interface{} {
	if value == "" {
		return gorm.Expr("1 = 1")
	}
	return gorm.Expr(query, value)
}
