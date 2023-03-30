package main

import (
	"log"
	"os"
	"time"

	"github.com/LastBit97/todolist/config"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"

	"github.com/LastBit97/todolist/router"
)

const defaultPort = "8000"

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://9ace5290285b413bbbd534c8ca39e1fc@o4504917501804544.ingest.sentry.io/4504917504819200",
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

	config.SetClient(client)

	log.Println("Server started on port " + port)

	server := gin.Default()
	rg := server.Group("/api")
	router.RegisterRouter(rg)
	log.Fatal(server.Run(":" + port))
}
