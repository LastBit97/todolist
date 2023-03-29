package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LastBit97/todolist/config"
	"github.com/getsentry/sentry-go"

	"github.com/LastBit97/todolist/middleware"
	"github.com/LastBit97/todolist/router"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "http://8246520e051040cf804cbad4416c5955@sentry.infotecs.int/17",
		Debug:            true,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//initiate Ent Client
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	//set the client to the variable defined in package config
	//this will enable the client intance to be accessed anywhere through the accessor which is a function
	//named GetClient
	config.SetClient(client)

	//initiate router and register all the route
	r := mux.NewRouter()
	r.Use(middleware.Header)
	router.RegisterRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}
