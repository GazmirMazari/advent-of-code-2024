package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func findTotalScore(name string) int {
	// Open the file
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed to open the file: %v", err)
	}
	defer file.Close()

	// Preallocate slices based on estimated number of lines
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("failed to get file info: %v", err)
	}
	estimatedLines := int(fileInfo.Size() / 10) // Estimate ~10 bytes per line
	leftList := make([]int, 0, estimatedLines)
	rightList := make([]int, 0, estimatedLines)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) < 2 {
			log.Fatalf("invalid line: %s", line)
		}

		// Convert and append integers directly
		leftNum, _ := strconv.Atoi(words[0]) // Assuming input is valid
		rightNum, _ := strconv.Atoi(words[1])
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// Parallel sorting using goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		sort.Ints(leftList)
		wg.Done()
	}()
	go func() {
		sort.Ints(rightList)
		wg.Done()
	}()

	// Wait for both sorts to complete
	wg.Wait()

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		diff := leftList[i] - rightList[i]
		if diff < 0 {
			diff = -diff
		}
		totalDistance += diff
	}

	return totalDistance
}
