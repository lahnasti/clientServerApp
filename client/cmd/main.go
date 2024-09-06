package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/lahnasti/clientServerApp/client"
)

func main() {
	// Определяем флаги
	uploadFlag := flag.String("upload", "", "Path to the file to upload")
	downloadFlag := flag.String("download", "", "Name of the file to download from the server")
	registerFlag := flag.Bool("register", false, "Register a new user")
	authFlag := flag.Bool("auth", false, "Login to get a token")
	loginFlag := flag.String("login", "", "Username for registration/login")
	passwordFlag := flag.String("password", "", "Password for registration/login")
	tokenFlag := flag.String("token", "", "JWT token for authentication")
	serverURL := flag.String("url", "http://server:8080", "URL of the server")

	flag.Parse()

	if *loginFlag == "" || *passwordFlag == "" {
		log.Printf("Username and password are required for registration/login.")
	}

	var token string
	var err error

	fmt.Println("Registering user...")
	if *registerFlag {
		err = client.RegisterUser(*serverURL, *loginFlag, *passwordFlag)
		if err != nil {
			log.Printf("Error registering user: %v", err)
		}
		fmt.Println("User registered successfully.")
	}

	fmt.Println("Logging in...")
	if *authFlag {
		token, err = client.LoginUser(*serverURL, *loginFlag, *passwordFlag)
		if err != nil {
			log.Fatalf("Error logging in: %v", err)
		}
		fmt.Println("Login successful. Token:", token)
	} else if *tokenFlag != "" {
		token = *tokenFlag
	}

	if token == "" && !*authFlag {
		log.Println("Token is required for file operations. Use -login flag to get a token.")
	}

	for {
		if *uploadFlag != "" {
			if token == "" {
				log.Fatal("Token is required for file operations. Use -login flag to get a token.")
			}
			err = client.UploadFile(*serverURL, *uploadFlag, token)
			if err != nil {
				log.Printf("Error uploading file: %v", err)
			} else {
				fmt.Println("File uploaded successfully.")
			}

		} else if *downloadFlag != "" {
			if token == "" {
				log.Fatal("Token is required for file operations. Use -login flag to get a token.")
			}
			err = client.DownloadFile(*serverURL, *downloadFlag, token)
			if err != nil {
				log.Printf("Error downloading file: %v", err)
			} else {
				fmt.Println("File downloaded successfully.")
			}

		} else {
			fmt.Println("Usage:")
			flag.PrintDefaults()
		}
		time.Sleep(5 * time.Minute)
	}
}
