package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	score := findTotalScore("input.txt")
	fmt.Printf("Total Score: %d\n", score)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findTotalScore(name string) int {
	// Open the file
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed to read the file: %v", err)
	}
	defer file.Close()

	// Initialize slices for the two columns
	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) < 2 {
			log.Fatalf("Invalid line: %s", line)
		}

		leftNum, err1 := strconv.Atoi(words[0])
		rightNum, err2 := strconv.Atoi(words[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Invalid number in line: %s", line)
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	return totalDistance

}
