package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(){
    err := godotenv.Load()
    if err != nil{
        panic("Internal Server Error :: Failed to load environment variables.")
    }

    log.Println("Status :: ENV is loaded.")
}
