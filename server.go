package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LastBit97/todolist/config"
	"github.com/LastBit97/todolist/middleware"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/LastBit97/todolist/router"
)

const defaultPort = "8000"

func main() {

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              "http://8246520e051040cf804cbad4416c5955@sentry.infotecs.int/17",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

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
	server.Use(sentrygin.New(sentrygin.Options{}))
	server.Use(middleware.SentryTraceMiddleware())
	rg := server.Group("/api")
	router.RegisterRouter(rg)
	log.Fatal(server.Run(":" + port))
}
