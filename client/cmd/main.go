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

	// Разбираем флаги
	flag.Parse()

	for {
		// Определяем, какой флаг был использован
		if *uploadFlag != "" {
			err := client.UploadFile(*uploadFlag)
			if err != nil {
				log.Printf("Error uploading file: %v", err)
			} else {
				fmt.Println("File uploaded successfully.")
			}
			// Ожидание перед следующей проверкой
			time.Sleep(10 * time.Second) // Задержка в 10 секунд
		} else if *downloadFlag != "" {
			err := client.DownloadFile(*downloadFlag)
			if err != nil {
				log.Printf("Error downloading file: %v", err)
			} else {
				fmt.Println("File downloaded successfully.")
			}
			// Ожидание перед следующей проверкой
			time.Sleep(10 * time.Second) // Задержка в 10 секунд
		} else {
			// Если нет аргументов, выводим инструкцию и ждём 10 секунд
			fmt.Println("Usage:")
			flag.PrintDefaults()
			time.Sleep(10 * time.Second) // Задержка в 10 секунд
		}
	}
}
