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
	currentGroupUniqueAnswers := make(map[string]bool)
	currentGroupIndex := 0
	scanner := bufio.NewScanner(fileInput)
	for scanner.Scan() {
		newLine := scanner.Text()
		if strings.TrimSpace(newLine) == "" {
			groupCounts = append(groupCounts, 0)
			currentGroupUniqueAnswers = make(map[string]bool)
			currentGroupIndex += 1
			continue
		}

		answers := strings.Split(newLine, "")
		for _, v := range answers {
			currentGroupUniqueAnswers[v] = true
		}
		groupCounts[currentGroupIndex] = len(currentGroupUniqueAnswers)
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
