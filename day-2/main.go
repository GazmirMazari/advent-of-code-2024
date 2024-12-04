package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

	lines := strings.Split(content, "\n")
	unsafe := 0
	totalReports := 0

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		totalReports++

		fields := strings.Fields(line)

		numbers := make([]int, 0, len(fields))
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				log.Printf("Error converting field to integer: %v", err)
				numbers = nil
				break
			}
			numbers = append(numbers, num)
		}

		if numbers == nil || len(numbers) < 2 {
			unsafe++
			continue
		}
		//fixed size window here, if the inputs would be differnt for individual cases this turns into for loop
		firstTwoPositionWindow := numbers[1] - numbers[0]

		if firstTwoPositionWindow == 0 || abs(firstTwoPositionWindow) < 1 || abs(firstTwoPositionWindow) > 3 {
			unsafe++
			continue
		}

		// Loop through the rest of the numbers
		for i := 2; i < len(numbers); i++ {
			windowDiff := numbers[i] - numbers[i-1]

			if windowDiff == 0 || !sameSign(firstTwoPositionWindow, windowDiff) || abs(windowDiff) <= 1 || abs(windowDiff) > 3 {
				unsafe++
				break
			}

		}

	}

	safe := totalReports - unsafe
	fmt.Printf("Number of safe reports: %d\n", safe)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Processing took %s\n", elapsedTime)
}

func sameSign(a, b int) bool {
	return (a >= 0 && b >= 0) || (a < 0 && b < 0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
