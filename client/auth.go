package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginUser(serverURL, login, password string) (string, error) {
	loginURL := fmt.Sprintf("%s/login", serverURL)
	loginData := map[string]string{
		"login":    login,
		"password": password,
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login data: %v", err)
	}

	resp, err := http.Post(loginURL, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to send login request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to login: %s", resp.Status)
	}

	var result map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	token, ok := result["token"].(string)
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}
