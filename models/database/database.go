package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
	var err error

	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=pinter sslmode=disable")
	if err != nil {
		log.Print(err)
	}
}
