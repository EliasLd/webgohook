package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/EliasLd/webgohook/internal/deploy"
	"github.com/EliasLd/webgohook/internal/security"
)

type WebhookHandler struct {
	secret string
}

func NewWebhookHandler(secret string) *WebhookHandler {
	return &WebhookHandler{secret: secret}
}

func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannort read request body", http.StatusBadRequest)
		return
	}

	sig := r.Header.Get("X-Hub-Signature-256")
	if !security.VerifyHMAC(body, sig, h.secret) {
		http.Error(w, "invalid signature", http.StatusUnauthorized)
		return
	}

	// Parse github payload
	var payload struct {
		Repository struct {
			Name string `json:"name"`
		} `json:"repository"`
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "invalid json payload", http.StatusBadRequest)
		return
	}

	repo := payload.Repository.Name
	fmt.Printf("Webhook received for repo: %s\n", repo)

	// Trigger corresponding service
	deploy.CallService(repo)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deployment triggered"))
}
