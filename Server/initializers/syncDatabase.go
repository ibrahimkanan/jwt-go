package initializers

// syncDatabase.go
import (
	"fmt"
	"jwt-go/models"
	"log"
)

var err error

func SyncDatabase() {

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	fmt.Println("Database migrated successfully")
}
