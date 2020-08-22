package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var DB *gorm.DB

func InitDatabase() error {
	db, err := gorm.Open("postgres", os.Getenv("GORM_CONN_STR"))
	if err != nil {
		return err
	}
	if os.Getenv("GIN_MODE") == "debug" {
		db.LogMode(true)
	}
	DB = db
	return nil
}
