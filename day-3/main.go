package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {

	startTime := time.Now()

	file, err := os.Open("input-day-3.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	pattern := `mul\((\d+),(\d+)\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(string(data), -1)

	total := 0

	for _, match := range matches {
		if len(match) == 3 {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			total += num1 * num2
		}
	}

	fmt.Printf("Total sum: %d\n", total)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Processing took %s\n", elapsedTime)
}
