package eight

import (
	"fmt"
	"strings"
)

const (
	acc = "acc"
	jmp = "jmp"
	nop = "nop"
)

type argument struct {
	sign  string
	value int
}

type instruction struct {
	operation string
	argument  argument
}

func parseDay2InputLine(s string) (lower int, upper int, char rune, password string, err error) {
	_, err = fmt.Sscanf(s, "%d-%d %c: %s ", &lower, &upper, &char, &password)
	return
}

func parseBootCode(input string) []instruction {
	var operation string
	var sign string
	var value int
	var instructions []instruction

	for _, row := range strings.Split(input, "\n") {
		if _, err := fmt.Sscanf(row, "%3s %1s%d", &operation, &sign, &value); err == nil {
			instructions = append(instructions, instruction{
				operation: operation,
				argument: argument{
					sign:  sign,
					value: value,
				},
			})
		}
	}

	return instructions
}

func GetAnswerCountPartOne(input string) int {
	accumulator := 0
	bootCode := parseBootCode(input)
	executedLines := map[int]bool{}
	fmt.Println(bootCode)

	for line := 0; line < len(bootCode); {
		if _, ok := executedLines[line]; ok {
			break
		}
		code := bootCode[line]
		executedLines[line] = true
		fmt.Println(code)

		if code.operation == acc {
			if code.argument.sign == "+" {
				accumulator += code.argument.value
			} else {
				accumulator -= code.argument.value
			}
			line++
		} else if code.operation == jmp {
			if code.argument.sign == "+" {
				line += code.argument.value
			} else {
				line -= code.argument.value
				if line < 0 {
					line = 0
				}
			}
		} else {
			line++
		}
	}

	return accumulator
}

func getAnswerCountPartTwo(input string) int {
	bootCode := parseBootCode(input)
	fmt.Println(bootCode)
	return 0
}
