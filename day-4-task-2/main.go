package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	minBirthYear = 1920
	maxBirthYear = 2002

	minIssuerYear = 2010
	maxIssuerYear = 2020

	minExpYear = 2020
	maxExpYear = 2030

	unitInch = "in"
	unitCm   = "cm"
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

func (p *passport) fillBirthYear(v string) error {
	birthYearStr := strings.TrimSpace(v)
	err := p.validateNumber(birthYearStr, minBirthYear, maxBirthYear)
	if err != nil {
		return err
	}
	p.BirthYear = birthYearStr
	return nil
}

func (p *passport) fillIssuerYear(v string) error {
	issYear := strings.TrimSpace(v)
	err := p.validateNumber(issYear, minIssuerYear, maxIssuerYear)
	if err != nil {
		return err
	}
	p.IssYear = issYear
	return nil
}
func (p *passport) fillExpYear(v string) error {
	expYear := strings.TrimSpace(v)
	err := p.validateNumber(expYear, minExpYear, maxExpYear)
	if err != nil {
		return err
	}
	p.ExpYear = expYear
	return nil
}

func (p *passport) fillHeight(v string) error {
	height := strings.TrimSpace(v)
	unitPosition := len(height) - 2
	unit := height[unitPosition:]
	if unit != unitInch && unit != unitCm {
		return fmt.Errorf("invalid unit: %v", unit)
	}

	value := height[:unitPosition]
	if unit == unitCm {
		minVal, maxVal := 150, 193
		err := p.validateNumber(value, minVal, maxVal)
		if err != nil {
			return err
		}
	}

	if unit == unitInch {
		minVal, maxVal := 59, 76
		err := p.validateNumber(value, minVal, maxVal)
		if err != nil {
			return err
		}
	}

	p.Height = height
	return nil
}

func (p *passport) fillHairColor(v string) error {
	hairColor := strings.TrimSpace(v)
	hairColorRegex, err := regexp.Compile(`^#[0-9a-f]{6}$`)
	if err != nil {
		return err
	}

	if ok := hairColorRegex.MatchString(hairColor); !ok {
		return fmt.Errorf("invalid hair color: %v", hairColor)
	}
	p.HairColor = hairColor
	return nil
}

func (p *passport) fillEyeColor(v string) error {
	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	eyeColor := strings.TrimSpace(v)
	if !inArray(eyeColor, validEyeColors) {
		return fmt.Errorf("invalid eye color: %v", eyeColor)
	}
	p.EyeColor = eyeColor
	return nil
}

func (p *passport) fillPassportID(v string) error {
	passportID := strings.TrimSpace(v)

	passportIDRegex, err := regexp.Compile(`^[0-9]{9}$`)
	if err != nil {
		return err
	}

	if ok := passportIDRegex.MatchString(passportID); !ok {
		return fmt.Errorf("invalid passport id: %v", passportID)
	}

	p.PassportID = passportID
	return nil
}

func (p *passport) fillCountryID(v string) error {
	p.CountryID = strings.TrimSpace(v)
	return nil
}

func (p *passport) Fill(keyval map[string]string) error {
	handlers := map[string]func(string) error{
		"byr": p.fillBirthYear,
		"iyr": p.fillIssuerYear,
		"eyr": p.fillExpYear,
		"hgt": p.fillHeight,
		"hcl": p.fillHairColor,
		"ecl": p.fillEyeColor,
		"pid": p.fillPassportID,
		"cid": p.fillCountryID,
	}

	for k, v := range keyval {
		if handler, ok := handlers[k]; ok {
			if err := handler(v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *passport) validateNumber(num string, min, max int) error {
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return err
	}
	if numInt < min || numInt > max {
		return fmt.Errorf("%s outside expected range: min=%d max=%d", num, min, max)
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
			continue
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

func inArray(item string, elements []string) bool {
	for _, v := range elements {
		if v == item {
			return true
		}
	}

	return false
}
