package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type passport struct {
	BirthYear  string
	IssYear    string
	ExpYear    string
	Height     string
	HairColor  string
	EyeColor   string
	PassportID string
	CountryID  string
	Processed  bool
}

func (p *passport) IsValid() bool {
	return p.BirthYear != "" &&
		p.IssYear != "" &&
		p.ExpYear != "" &&
		p.PassportID != "" &&
		//p.CountryID > 0 &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != ""
}

func (p *passport) Fill(keyval map[string]string) error {
	ok := true
	for k, v := range keyval {
		switch k {
		case "byr":
			p.BirthYear = strings.TrimSpace(v)
			if !ok {
				return fmt.Errorf("invalid birth year: %v", v)
			}
		case "iyr":
			p.IssYear = strings.TrimSpace(v)
			if !ok {
				return fmt.Errorf("invalid issuer year: %v", v)
			}
		case "eyr":
			p.ExpYear = strings.TrimSpace(v)
			if !ok {
				return fmt.Errorf("invalid exp year: %v", v)
			}
		case "hgt":
			p.Height = strings.TrimSpace(v)
			if !ok {
				return fmt.Errorf("invalid height: %v", v)
			}
		case "hcl":
			p.HairColor = strings.TrimSpace(v)
			if !ok {
				return fmt.Errorf("invalid HairColor: %v", v)
			}
		case "ecl":
			p.EyeColor = strings.TrimSpace(v)
			if !ok {
				return fmt.Errorf("invalid EyeColor: %v", v)
			}
		case "pid":
			p.PassportID = strings.TrimSpace(v)
			if !ok {
				return fmt.Errorf("invalid PassportID: %v", v)
			}
		case "cid":
			p.CountryID = strings.TrimSpace(v)
			if !ok {
				return fmt.Errorf("invalid CountryID: %v", v)
			}
		}
	}

	return nil
}

func main() {
	fileInput, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open input: %v", err)
	}

	validPassports := 0
	passports := 0
	var activePassport *passport
	scanner := bufio.NewScanner(fileInput)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			if activePassport == nil {
				passports += 1
				activePassport = new(passport)
			}
		}

		if strings.TrimSpace(line) == "" {
			if activePassport.IsValid() {
				validPassports += 1
			}
			activePassport.Processed = true
			activePassport = nil
			continue
		}

		entries := parsePassportLine(line)
		if err := activePassport.Fill(entries); err != nil {
			log.Fatalf("could not parse passport records: %v", err)
		}
	}

	if activePassport != nil && !activePassport.Processed {
		if activePassport.IsValid() {
			validPassports += 1
		}
		activePassport = nil
	}

	fmt.Printf("passports: %d, Valid Passports: %d\n", passports, validPassports)
	if err := scanner.Err(); err != nil {
		log.Fatalf("could not scan input: %v", err)
	}
	if err := fileInput.Close(); err != nil {
		log.Fatalf("could not close input: %v", err)
	}
}

func parsePassportLine(line string) map[string]string {
	result := make(map[string]string)
	components := strings.Fields(line)
	for _, field := range components {
		fieldParts := strings.Split(field, ":")
		if len(fieldParts) < 2 {
			continue
		}
		var key, value string = fieldParts[0], fieldParts[1]
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if key != "" && value != "" {
			result[key] = value
		}
	}
	return result
}
