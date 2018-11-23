package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"))

	a.Run(os.Getenv("SERVER_PORT"))
}
