package main

import (

	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <command> <filename>")
		fmt.Println("Commands: upload, download")
		return
	}

	command := os.Args[1]
	filename := os.Args[2]

	switch command {
	case "upload":
		err := UploadFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
		}
	case "download":
		err := DownloadFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}
