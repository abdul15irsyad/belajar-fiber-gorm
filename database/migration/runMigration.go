package migration

import (
	"abdul15irsyad/belajar-fiber-gorm/database"
	"abdul15irsyad/belajar-fiber-gorm/models/entity"
	"log"
)

func RunMigration() {
	database.DB.AutoMigrate(entity.User{})
	log.Println("Running auto migrate")
}
