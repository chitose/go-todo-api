package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
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

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", 15*time.Second,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()

	// jwt middleware
	r.Use(jwtMiddleware)

	r.HandleFunc("/auth/{provider}", authHandler)
	r.HandleFunc("/auth/{provider}/callback", authCallbackHandler)
	r.HandleFunc("/", homeHandler)

	setupUserRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// run the server in a goroutine -> non block
	go func() {
		log.Println("listening on localhost:3000")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// accept gracefull shutdown when quit via SIGINT (CTRL+C)
	// SIGKILL, SIGQUIT or SIGTERM will not be caught
	signal.Notify(c, os.Interrupt)

	// block until we receive our signal
	<-c

	// create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
