package main

import (
	"log"

	"github.com/trivektor/gorm-resumes-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
dsn := "host=localhost user=root password=root dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(dsn))

	if (err != nil) {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.Resume{}, 
		&models.Intro{},
		&models.Education{}, 
		&models.Position{},
		&models.Reference{},
		&models.Skill{},
	)
}