package initializer

import (
	"os"

	"github.com/charmbracelet/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	db_name := os.Getenv("DB_NAME") + ".db"
	DB, err = gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database")
		panic(err)
	}
	log.Info("Status :: Database is up and running")
}
