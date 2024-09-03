package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lahnasti/clientServerApp/client"
)

func main() {
	// Определяем флаги
	uploadFlag := flag.String("upload", "", "Path to the file to upload")
	downloadFlag := flag.String("download", "", "Name of the file to download from the server")

	// Разбираем флаги
	flag.Parse()

	// Определяем, какой флаг был использован
	if *uploadFlag != "" {
		err := client.UploadFile(*uploadFlag)
		if err != nil {
			log.Fatalf("Error uploading file: %v", err)
		}
		fmt.Println("File uploaded successfully.")
	} else if *downloadFlag != "" {
		err := client.DownloadFile(*downloadFlag)
		if err != nil {
			log.Fatalf("Error downloading file: %v", err)
		}
		fmt.Println("File downloaded successfully.")
	} else {
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}
}
