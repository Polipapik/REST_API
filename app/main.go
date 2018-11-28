package main

import (
	"os"

	"github.com/Polipapik/REST_API/app/routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	//time.Sleep(5 * time.Second)
	a := routes.App{}
	a.Initialize(
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"))

	a.Run(os.Getenv("SERVER_PORT"))
}
