package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	file, err := os.Open("input-day2.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	tst, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("File reading took %s\n", elapsedTime)

	// Optionally, use the data variable
	fmt.Println(string(t))
}
