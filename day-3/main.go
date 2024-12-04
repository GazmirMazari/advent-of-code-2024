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

	file, err := os.Open("input-day-3.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	tst, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	content := string(tst)

	fmt.Println(content)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Processing took %s\n", elapsedTime)
}
