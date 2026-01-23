package deploy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var servicePorts map[string]int

func init() {
	file, err := os.ReadFile("config/services.json")
	if err != nil {
		fmt.Printf("Error reading services.json: %v", err)
		servicePorts = make(map[string]int)
		return
	}

	err = json.Unmarshal(file, &servicePorts)
	if err != nil {
		fmt.Printf("Error parsing services.json: %v", err)
		servicePorts = make(map[string]int)
		return
	}
}

func CallService(service string) {
	port, ok := servicePorts[service]
	if !ok {
		fmt.Printf("Service %s not found\n", service)
		return
	}

	url := fmt.Sprintf("http://127.0.0.1:%d/deploy", port)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		fmt.Printf("Error calling %s: %v\n", service, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Triggered deploy for %s, status %s\n", service, resp.Status)
}
