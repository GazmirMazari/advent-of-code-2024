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
	//split the content into lines
	lines :=strings.Split(content, "\n")
	unsafe := 0
	for _, line :=range lines {
		fields := strings.Fields(line)

    
        numbers := make([]int, 0, len(fields))
        for _, field := range fields {
            num, err := strconv.Atoi(field)
            if err != nil {
                
                continue
            }
            numbers = append(numbers, num)
        }

		
		//check each line if if is safe or not
		for i:=1; i <len(numbers); i++{
			diff := abs(numbers[i] - numbers[i-1])
			if diff <= 1 || diff > 3 {
				unsafe++
			}
		}
	}

	//substract length of lines here
	safe := len(lines) - unsafe
	fmt.Printf("lines being printed are %d", safe)
	elapsedTime := time.Since(startTime)
	fmt.Printf("File reading took %s\n", elapsedTime)

}

func abs (x int) int {
	if x < 0 {
		return -x
	}
	return x
}
