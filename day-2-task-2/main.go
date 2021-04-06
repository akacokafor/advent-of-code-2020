package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type passwordRule struct {
	MinAppearance int
	MaxAppearance int
	SubjectChar   byte
}

func (p *passwordRule) Validate(password string) bool {

	positionOnePresent := password[(p.MinAppearance-1)] == p.SubjectChar
	positionTwoPresent := password[(p.MaxAppearance-1)] == p.SubjectChar
	if positionOnePresent && positionTwoPresent {
		return false
	}

	return positionOnePresent || positionTwoPresent
}

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
	validPassword := 0

	for scanner.Scan() {
		var input passwordRule
		var testSubject string
		fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &input.MinAppearance, &input.MaxAppearance, &input.SubjectChar, &testSubject)
		if input.Validate(testSubject) {
			validPassword += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not scan input: %v", err)
	}

	fmt.Printf("Valid Password: %d\n", validPassword)

}
