package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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


	content := string(tst)
	//split the content into lines
	lines :=strings.Split(content, "\n")

	for _, line :=range lines {
		fields := strings.Fields(line)

        // Convert fields to integers
        numbers := make([]int, 0, len(fields))
        for _, field := range fields {
            num, err := strconv.Atoi(field)
            if err != nil {
                log.Printf("Error converting field to integer on line %d: %v", lineNumber, err)
                continue
            }
            numbers = append(numbers, num)
        }

	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("File reading took %s\n", elapsedTime)

	// Optionally, use the data variable
	fmt.Println(string(tst))
}
