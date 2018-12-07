package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Polipapik/REST_API/app/models"
	"github.com/Polipapik/REST_API/app/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var env *routes.Env

func init() {
	log.Println("Connection...")

	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_ADDRESS"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL"))

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection successful")

	env = &routes.Env{Router: mux.NewRouter(), Country: &models.GormDB{DB: db}}
	env.InitializeRoutes()
}

func main() {
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), env.Router))
}
