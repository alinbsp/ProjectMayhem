package main

import(
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Using this structure we will store participant information
// Each repository will be mapped to a list of participants
type Repository struct {
	ID           int    `gorm:"primary_key"`
	Identifier   string
	Participants string
}

func main() {
	const connStr = "postgresql://mayhem@localhost:26257/mayhem?sslmode=disable"

	// Connect to the database
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error when connecting to the database:\n", err)
	}
	defer db.Close()

	// Cleanup any old data
	if db.HasTable(&Repository{}) {
		db.DropTable(&Repository{})
	}

	// Ensure the table is created
	db.AutoMigrate(&Repository{})

	// Insert some test data
	db.Create(&Repository{
		ID: 1,
		Identifier: "ProjectMayhem",
		Participants: "127.0.0.1;8.8.8.8",
	})

	// Print out all rows
	var repositories []Repository
	db.Find(&repositories)
	fmt.Println("Repositories found in the database:")
	for _, repo := range repositories {
		fmt.Printf("%d %s %s\n", repo.ID, repo.Identifier, repo.Participants)
	}
}

