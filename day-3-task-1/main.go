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

	defer func() {
		if err := fileInput.Close(); err != nil {
			log.Fatalf("could not close input: %v", err)
		}
	}()

	scanner := bufio.NewScanner(fileInput)
	numberOfTrees := 0
	currentPositionIndex := 0
	patternWidth := -1
	for scanner.Scan() {

		line := scanner.Text()
		lineArr := strings.Split(line, "")
		if patternWidth < 0 {
			patternWidth = len(lineArr)
		}

		if lineArr[currentPositionIndex] == "#" {
			numberOfTrees += 1
		}

		nextPosition := currentPositionIndex + 3
		if nextPosition >= patternWidth {
			nextPosition = nextPosition - patternWidth
		}
		currentPositionIndex = nextPosition
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not scan input: %v", err)
	}

	fmt.Printf("Trees found: %d", numberOfTrees)
}
