package four

import (
	"regexp"
	"strconv"
	"strings"
)

func valueBetween(value int, min int, max int) bool {
	return value >= min && value <= max
}

func strToInt(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}

var validators = map[string]func(input string) bool{
	"byr": func(input string) bool {
		return valueBetween(strToInt(input), 1920, 2002)
	},
	"iyr": func(input string) bool {
		return valueBetween(strToInt(input), 2010, 2020)
	},
	"eyr": func(input string) bool {
		return valueBetween(strToInt(input), 2020, 2030)
	},
	"hgt": func(input string) bool {
		value := strToInt(regexp.MustCompile(`(in|cm)$`).ReplaceAllString(input, ""))

		if strings.HasSuffix(input, "cm") {
			return valueBetween(value, 150, 193)
		} else if strings.HasSuffix(input, "in") {
			return valueBetween(value, 59, 76)
		}

		return false
	},
	"hcl": func(input string) bool {
		return regexp.MustCompile(`^#[a-f0-9]{6}$`).MatchString(input)
	},
	"ecl": func(input string) bool {
		return regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`).MatchString(input)
	},
	"pid": func(input string) bool {
		return regexp.MustCompile(`^[0-9]{9}$`).MatchString(input)
	},
	"cid": func(input string) bool {
		return false
	},
}

func countValidPassports(input string, validateValues bool) int {
	var passportDelimiter = regexp.MustCompile(`\n{2,}`)
	var fieldExtractor = regexp.MustCompile(`(\w+):([#\w]+)(?:[\n\s]+)?`)
	validPassports := 0

	for _, passport := range passportDelimiter.Split(input, -1) {
		fieldMatches := fieldExtractor.FindAllStringSubmatch(passport, -1)
		foundFields := 0

		for _, match := range fieldMatches {
			if validateValues {
				if validator, ok := validators[match[1]]; ok && validator(match[2]) {
					foundFields++
				}
			} else {
				if match[1] != "cid" {
					foundFields++
				}
			}
		}

		if foundFields == 7 {
			validPassports++
		}
	}

	return validPassports
}
