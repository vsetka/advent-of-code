package two

import (
	"regexp"
	"strconv"
	"unicode/utf8"
)

type passwordInfo struct {
	FirstIndex          int
	SecondIndex         int
	ConstraintCharacter rune
	Password            string
}

func parsePasswordInfo(input string) passwordInfo {
	passwordInfoMatcher := regexp.MustCompile(`(?P<FirstIndex>\d+)-(?P<SecondIndex>\d+)\s+(?P<ConstraintCharacter>\w{1}):\s+(?P<Password>.+)`)
	match := passwordInfoMatcher.FindStringSubmatch(input)
	passwordInfo := passwordInfo{}

	for i, name := range passwordInfoMatcher.SubexpNames() {
		if i != 0 && name != "" {
			switch name {
			case "FirstIndex":
				passwordInfo.FirstIndex, _ = strconv.Atoi(match[i])
			case "SecondIndex":
				passwordInfo.SecondIndex, _ = strconv.Atoi(match[i])
			case "ConstraintCharacter":
				passwordInfo.ConstraintCharacter, _ = utf8.DecodeRuneInString(match[i])
			case "Password":
				passwordInfo.Password = match[i]
			}
		}
	}

	return passwordInfo
}

func isValidPartOne(input string) bool {
	passwordInfo := parsePasswordInfo(input)
	foundCount := 0

	for _, char := range passwordInfo.Password {
		if char == passwordInfo.ConstraintCharacter {
			foundCount++
		}
	}

	return foundCount >= passwordInfo.FirstIndex && foundCount <= passwordInfo.SecondIndex
}

func isValidPartTwo(input string) bool {
	passwordInfo := parsePasswordInfo(input)
	foundCount := 0

	for idx, char := range passwordInfo.Password {
		if idx+1 == passwordInfo.FirstIndex && char == passwordInfo.ConstraintCharacter {
			foundCount++
		} else if idx+1 == passwordInfo.SecondIndex && char != passwordInfo.ConstraintCharacter {
			foundCount--
		}
	}

	return foundCount == 0
}
