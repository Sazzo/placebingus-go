package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)



func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Hello).Methods("GET")
	router.HandleFunc("/image", Image).Methods("GET")

	fmt.Println("PLACEBINGUS - Golang Edition")
	fmt.Println("Made by FelipeSazz :3")
	fmt.Println("Running at http://localhost:8080/image")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "bingus my beloved"}

	json.NewEncoder(w).Encode(response)
}

func Image(w http.ResponseWriter, r *http.Request) {
	imageBox := packr.NewBox("/assets/images")
	imagesList := imageBox.List()
	
	rand.Seed(time.Now().Unix())
	randomImage := imagesList[rand.Intn(len(imagesList))]

	image, _ := imageBox.Open(randomImage)
	imageBytes, _ := imageBox.Find(randomImage)
	mimeType := mimetype.Detect(imageBytes)

	w.Header().Set("Content-Type", mimeType.String())
	io.Copy(w, image)
}