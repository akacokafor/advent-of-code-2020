package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	var inputs []int

	for scanner.Scan() {
		inputNumber := 0
		fmt.Sscanf(scanner.Text(), "%d", &inputNumber)
		for index, entry := range inputs {
			for _, secondEntry := range inputs[index:] {
				if entry+secondEntry+inputNumber == 2020 {
					fmt.Printf("Answer is: %d\n", entry*secondEntry*inputNumber)
					break
				}
			}
		}
		inputs = append(inputs, inputNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not scan input: %v", err)
	}
}
