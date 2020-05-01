package main

import (
	"log"
	"net/http"
	"os"
	"shal/router"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("conf.env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	appPort := os.Getenv("PORT")
	log.Println("App running on port: " + appPort)

	r := router.SetupRouter()

	s := &http.Server{
		Addr:           appPort,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
