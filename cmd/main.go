package main

import (
	"log"

	"github.com/trivektor/gorm-resumes-api/pkg/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=root dbname=postgres port=5432"
	var err error

	database.DB, err = gorm.Open(postgres.Open(dsn))

	if (err != nil) {
		log.Fatalln(err)
	}
}