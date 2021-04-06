package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	finalResult := 1
	slopes := []struct {
		Right  int
		Bottom int
	}{
		{
			Right:  1,
			Bottom: 1,
		},
		{
			Right:  3,
			Bottom: 1,
		},
		{
			Right:  5,
			Bottom: 1,
		},
		{
			Right:  7,
			Bottom: 1,
		},
		{
			Right:  1,
			Bottom: 2,
		},
	}

	for _, s := range slopes {
		fileInput, err := os.Open("input.txt")
		if err != nil {
			log.Fatalf("could not open input: %v", err)
		}

		scanner := bufio.NewScanner(fileInput)
		numberOfTrees := 0
		currentPositionIndex := 0
		patternWidth := -1

		rightIncr := s.Right
		bottomIncr := s.Bottom
		rowsSeen := 0

		for scanner.Scan() {
			previousRowCount := rowsSeen
			rowsSeen += 1
			if previousRowCount%bottomIncr != 0 {
				continue
			}

			line := scanner.Text()
			lineArr := strings.Split(line, "")
			if patternWidth < 0 {
				patternWidth = len(lineArr)
			}

			if lineArr[currentPositionIndex] == "#" {
				numberOfTrees += 1
			}

			nextPosition := currentPositionIndex + rightIncr
			if nextPosition >= patternWidth {
				nextPosition = nextPosition - patternWidth
			}
			currentPositionIndex = nextPosition
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("could not scan input: %v", err)
		}

		fmt.Printf("Trees found: %d\n", numberOfTrees)
		finalResult *= numberOfTrees
		if err := fileInput.Close(); err != nil {
			log.Fatalf("could not close input: %v", err)
		}
	}

	fmt.Printf("Multiply result: %d\n", finalResult)
}
