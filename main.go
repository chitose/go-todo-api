package main

import (
	"net/http"
	"time"

	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("runtime.env")

	if err != nil {
		log.Fatal("Could not load runtime.env")
	}

	setup()

	r := mux.NewRouter()

	r.HandleFunc("/auth/{provider}", authHandler)
	r.HandleFunc("/auth/{provider}/callback", authCallbackHandler)
	r.HandleFunc("/", homeHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("listening on localhost:3000")
	log.Fatal(srv.ListenAndServe())

}
