package deploy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Service struct {
	URL      string `json:"url"`
	TokenEnv string `json:"token_env"`
}

var services map[string]Service

func init() {
	file, err := os.ReadFile("config/services.json")
	if err != nil {
		fmt.Printf("Error reading services.json: %v", err)
		services = make(map[string]Service)
		return
	}

	err = json.Unmarshal(file, &services)
	if err != nil {
		fmt.Printf("Error parsing services.json: %v", err)
		services = make(map[string]Service)
		return
	}
}

func CallService(serviceName string) {
	service, ok := services[serviceName]
	if !ok {
		fmt.Printf("Service %s not found\n", service)
		return
	}

	token := os.Getenv(service.TokenEnv)
	if token == "" {
		fmt.Printf("Token for service %s not found in env (%s)\n", serviceName, service.TokenEnv)
	}

	req, err := http.NewRequest("POST", service.URL, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		fmt.Printf("Error creating request for %s: %v\n", serviceName, err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Deploy-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error calling %s: %v\n", serviceName, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Triggered deploy for %s, status %s\n", serviceName, resp.Status)
}
