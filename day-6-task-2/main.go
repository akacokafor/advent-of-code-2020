package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileInput, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open input: %v", err)
	}

	groupCounts := []int{
		0,
	}
	currentGroupUniqueAnswers := make(map[string]int)
	currentGroupIndex := 0
	currentGroupSize := 0

	scanner := bufio.NewScanner(fileInput)
	for scanner.Scan() {
		newLine := scanner.Text()
		if strings.TrimSpace(newLine) == "" {
			for _, v := range currentGroupUniqueAnswers {
				if v == currentGroupSize {
					groupCounts[currentGroupIndex] += 1
				}
			}
			currentGroupSize = 0
			groupCounts = append(groupCounts, 0)
			currentGroupUniqueAnswers = make(map[string]int)
			currentGroupIndex += 1
			continue
		}

		currentGroupSize += 1
		answers := strings.Split(newLine, "")
		for _, v := range answers {
			currentValue := 0
			if _, ok := currentGroupUniqueAnswers[v]; ok {
				currentValue = currentGroupUniqueAnswers[v]
			}
			currentGroupUniqueAnswers[v] = currentValue + 1
		}
	}

	for _, v := range currentGroupUniqueAnswers {
		if v == currentGroupSize {
			groupCounts[currentGroupIndex] += 1
		}
	}

	totalSum := 0
	for _, v := range groupCounts {
		totalSum += v
	}

	fmt.Printf("Sum: %d\n", totalSum)
	if err := scanner.Err(); err != nil {
		log.Fatalf("could not scan input: %v", err)
	}
	if err := fileInput.Close(); err != nil {
		log.Fatalf("could not close input: %v", err)
	}
}
