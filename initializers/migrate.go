package initializers

import (
	"github.com/Kachyr/crud/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Post{})
}
