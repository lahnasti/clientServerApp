package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterUser(url, login, password string) error {
	payload := map[string]string{
		"login": login,
		"password": password,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url+"/register", "application/json", bytes.NewReader(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("registration failed: %s", resp.Status)
	}
	return nil
}
