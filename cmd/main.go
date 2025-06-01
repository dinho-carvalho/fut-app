package main

import (
	"fmt"
	"log"
	"net/http"

	"fut-app/internal/database"
	"fut-app/internal/database/models"
	"fut-app/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	config := database.NewConfig()
	db, err := database.NewDatabase(config)
	if err != nil {
		log.Fatal("❌ Failed to connect to the database")
	}

	fmt.Println("✅ Successfully connected to the database!")

	err = db.AutoMigrate(&models.Player{}, &models.Match{}, &models.Rating{})
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	routes.CreateRoutes(r, db)

	fmt.Println("🚀 Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	fmt.Println("It's time ⚽ ⚽ ⚽ ⚽ ⚽ ⚽ ")
}
