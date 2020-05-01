package controller

import (
	"log"
	"net/http"
	models "shal/model"
)

//SetupController sample set up
func SetupController(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome")
}

func TestDatabase(w http.ResponseWriter, r *http.Request) {
	dbString := models.GetDB()
	log.Println(dbString)
}
