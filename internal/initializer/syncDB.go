package initializer

import "ecom-backend/internal/models"

func SyncDB() {
	DB.AutoMigrate(
		&models.Book{},
		&models.User{},
	)
}
