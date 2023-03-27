package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/carlosarraes/qcmeback/controllers"
	"github.com/carlosarraes/qcmeback/models"
	"github.com/joho/godotenv"
)

func main() {
	port := 8080
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	app := controllers.App{}
	flag.StringVar(&app.DSN, "dsn", os.Getenv("DB_URI"), "Postgres DSN")
	flag.Parse()

	conn, err := app.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer conn.Close()

	app.DB = &models.Postgres{DB: conn}

	server := app.Routes()

	log.Printf("Server listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
