package initializer

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectToDB(){
    var err error

    db_name := os.Getenv("DB_NAME") + ".db"
    DB, err = gorm.Open(sqlite.Open(db_name), &gorm.Config{})
    if err != nil{
        panic(err)
    }
    log.Println("Status :: Database is up and running")
}
