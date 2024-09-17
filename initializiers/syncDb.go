package initializiers

import "github.com/f1nn-ach/go-jwt/models"

func SyncDb() {
	DB.AutoMigrate(&models.Book{})
}
