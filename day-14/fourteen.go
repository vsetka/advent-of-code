package fourteen

import (
	"fmt"
	"math"
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

type maskInfo struct {
	masks        []string
	modifierBits map[int]bool
}

func getMaskInfo(mask string) maskInfo {
	mi := maskInfo{
		masks:        []string{},
		modifierBits: map[int]bool{},
	}
	floatingBitCount := 0

	for idx, bit := range strings.Split(mask, "") {
		if bit == "X" {
			floatingBitCount++
		}
		if bit != "0" {
			mi.modifierBits[idx] = true
		}
	}

	for i := int64(0); i < int64(math.Pow(float64(2), float64(floatingBitCount))); i++ {
		variation := mask
		format := "%0" + fmt.Sprint(floatingBitCount) + "s"
		combination := fmt.Sprintf(format, strconv.FormatInt(i, 2))
		for _, bit := range strings.Split(combination, "") {
			variation = strings.Replace(variation, "X", bit, 1)
		}
		mi.masks = append(mi.masks, variation)
	}

	return mi
}

func getAnswerCountPartTwo(input string) int64 {
	maskMatcher := regexp.MustCompile(`mask\s+=\s+(\w+)`)
	opMatcher := regexp.MustCompile(`mem\[(\d+)\]\s+=\s+(\d+)`)
	memory := map[uint64]int64{}
	var mi maskInfo

	for _, row := range strings.Split(input, "\n") {
		if maskMatch := maskMatcher.FindStringSubmatch(row); maskMatch != nil {
			mi = getMaskInfo(maskMatch[1])
			continue
		}

		opMatch := opMatcher.FindStringSubmatch(row)
		address, _ := strconv.ParseUint(opMatch[1], 10, 64)
		paddedAddress := fmt.Sprintf("%036s", strconv.FormatUint(address, 2))
		value, _ := strconv.ParseInt(opMatch[2], 10, 64)

		for _, maskVariation := range mi.masks {
			decodedAddress := paddedAddress
			for modifier := range mi.modifierBits {
				decodedAddress = decodedAddress[:modifier] + string(maskVariation[modifier]) + decodedAddress[modifier+1:]
			}
			addressInt, _ := strconv.ParseUint(decodedAddress, 2, 64)
			memory[addressInt] = value
		}
	}

	sum := int64(0)
	for _, value := range memory {
		sum += value
	}

	return sum
}
