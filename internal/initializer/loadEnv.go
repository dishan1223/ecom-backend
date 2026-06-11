package initializer

import (
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func MustLoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Status :: Internal Server Error :: Failed to load environment variables.")
		panic("Internal Server Error :: Failed to load environment variables.")
	}

	log.Info("Status :: Environment variables loaded")
}
