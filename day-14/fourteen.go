package fourteen

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getAnswerCountPartOne(input string) int64 {
	// 0000 0000 0000 0000 0000 0000 0000 0000 0000
	// ^ most significant bit	(2^35)							^ least significant bit (2^0)

	maskMatcher := regexp.MustCompile(`mask\s+=\s+(\w+)`)
	opMatcher := regexp.MustCompile(`mem\[(\d+)\]\s+=\s+(\d+)`)
	memory := map[uint16]int64{} // max memory address is 65434, fits in a uint16, value is 36 bits so we must use int64
	var mask string

	for _, row := range strings.Split(input, "\n") {
		if maskMatch := maskMatcher.FindStringSubmatch(row); maskMatch != nil {
			mask = maskMatch[1]
			continue
		}

		opMatch := opMatcher.FindStringSubmatch(row)
		u64, _ := strconv.ParseUint(opMatch[1], 10, 16)
		address := uint16(u64)
		value, _ := strconv.ParseInt(opMatch[2], 10, 64)
		maskedValue := fmt.Sprintf("%036s", strconv.FormatInt(value, 2))

		for idx, bit := range strings.Split(mask, "") {
			if bit != "X" {
				maskedValue = maskedValue[:idx] + bit + maskedValue[idx+1:]
			}
		}

		maskedInt, _ := strconv.ParseInt(maskedValue, 2, 64)
		memory[address] = maskedInt
	}

	sum := int64(0)
	for _, value := range memory {
		sum += value
	}

	return sum
}

func getAnswerCountPartTwo(input string) int {
	return 0
}

func GetAnswerCountPartOne(input string) int64 {
	return getAnswerCountPartOne(input)
}

func GetAnswerCountPartTwo(input string) int {
	return getAnswerCountPartTwo(input)
}
