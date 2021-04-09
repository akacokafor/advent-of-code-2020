package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

func main() {
	fileInput, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open input: %v", err)
	}

	maxSeatId := 0
	var seatIds []int
	scanner := bufio.NewScanner(fileInput)
	for scanner.Scan() {
		_, _, seatId, err := decodeBoardingPass(scanner.Text())
		if err != nil {
			continue
		}

		seatIds = append(seatIds, seatId)
		if maxSeatId < seatId {
			maxSeatId = seatId
		}
	}

	sort.Ints(seatIds)
	missingSeat := -1
	for index, seatId := range seatIds {
		if prevIndex := index - 1; prevIndex > 0 {
			prevItem := seatIds[prevIndex]
			if seatId-prevItem > 1 {
				missingSeat = prevItem + 1
				break
			}
		}
	}

	fmt.Printf("Missing seat id: %d\n", missingSeat)

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not scan input: %v", err)
	}
	if err := fileInput.Close(); err != nil {
		log.Fatalf("could not close input: %v", err)
	}
}

func decodeBoardingPass(input string) (row int, column int, seatId int, err error) {
	if len(input) < 10 {
		return 0, 0, 0, fmt.Errorf("invalid boarding pass")
	}

	boardinPassRegex := regexp.MustCompile(`^[FB]{7}[LR]{3}$`)
	if ok := boardinPassRegex.MatchString(input); !ok {
		return 0, 0, 0, fmt.Errorf("invalid boarding pass input: %v", input)
	}

	row, err = getRowCount(input[:7])
	if err != nil {
		return 0, 0, 0, err
	}

	column, err = getColumnCount(input[7:])
	if err != nil {
		return 0, 0, 0, err
	}

	seatId = (row * 8) + column
	return row, column, seatId, nil
}

func getRowCount(input string) (int, error) {
	if len(input) < 7 {
		return 0, fmt.Errorf("invalid boarding pass")
	}

	rowInput := input[:7]
	rowRegex := regexp.MustCompile(`^[FB]{7}$`)
	if ok := rowRegex.MatchString(rowInput); !ok {
		return 0, fmt.Errorf("invalid row input: %v", rowInput)
	}
	lowerBound := 0
	upperBound := 127
	for _, entry := range rowInput {
		if entry == 'F' {
			upperBound = ((lowerBound + upperBound + 1) / 2) - 1
		}

		if entry == 'B' {
			lowerBound = ((lowerBound + upperBound + 1) / 2)
		}
	}

	if rowInput[6] == 'F' {
		return lowerBound, nil
	}

	return upperBound, nil
}

func getColumnCount(input string) (int, error) {
	if len(input) < 3 {
		return 0, fmt.Errorf("invalid boarding pass")
	}

	columnInput := input[:3]
	columnRegex := regexp.MustCompile(`^[LR]{3}$`)
	if ok := columnRegex.MatchString(columnInput); !ok {
		return 0, fmt.Errorf("invalid column input: %v", columnInput)
	}

	lowerBound := 0
	upperBound := 7
	for _, entry := range columnInput {
		if entry == 'L' {
			upperBound = ((lowerBound + upperBound + 1) / 2) - 1
		}

		if entry == 'R' {
			lowerBound = ((lowerBound + upperBound + 1) / 2)
		}
	}

	if columnInput[2] == 'L' {
		return lowerBound, nil
	}

	return upperBound, nil
}
