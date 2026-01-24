package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EliasLd/webgohook/internal/handler"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("WEBHOOK_PORT")
	if port == "" {
		port = "8081"
	}

	secret := os.Getenv("WEBHOOK_SECRET")
	if secret == "" {
		log.Fatalf("WEBHOOK_SECRET must be set")
	}

	http.Handle("/webhook", handler.NewWebhookHandler(secret))

	fmt.Println("Webhook server listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
