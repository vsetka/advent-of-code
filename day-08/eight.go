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

type executionNode struct {
	lineNumber  int
	instruction instruction
	parents     []int
	child       *int
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

func runCode(bootCode []instruction) (int, bool) {
	accumulator := 0
	loopFound := false
	executedLines := map[int]bool{}

	for line := 0; line < len(bootCode); {
		if _, ok := executedLines[line]; ok {
			loopFound = true
			break
		}
		code := bootCode[line]
		executedLines[line] = true

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

	return accumulator, loopFound
}

func getAnswerCountPartOne(input string) int {
	bootCode := parseBootCode(input)
	accumulator, _ := runCode(bootCode)

	return accumulator
}

// generateDigraph generates a graphviz digraph for visual debugging
func generateDigraph(executionPlan []*executionNode) string {
	graph := []string{}
	for _, item := range executionPlan {
		var from string
		var to string

		if item.instruction.operation == jmp {
			if item.child != nil {
				from = fmt.Sprintf(
					"jmp_%d_to_%d",
					item.lineNumber,
					executionPlan[*item.child].lineNumber,
				)
			} else {
				from = fmt.Sprintf(
					"jmp_%d_to_end",
					item.lineNumber,
				)
			}
		} else {
			from = fmt.Sprintf(
				"%s_%d",
				item.instruction.operation,
				item.lineNumber,
			)
		}

		if item.lineNumber == 0 {
			graph = append(graph, fmt.Sprintf("start -> %s", from))
		}

		var destinationLine int
		if item.child != nil {
			destinationLine = *item.child
		} else {
			destinationLine = item.lineNumber + 1
		}

		if destinationLine >= len(executionPlan) {
			to = "end"
		} else {
			if executionPlan[destinationLine].instruction.operation == jmp {
				if executionPlan[destinationLine].child != nil {
					to = fmt.Sprintf(
						"jmp_%d_to_%d",
						executionPlan[destinationLine].lineNumber,
						executionPlan[*executionPlan[destinationLine].child].lineNumber,
					)
				} else {
					to = fmt.Sprintf(
						"jmp_%d_to_end",
						executionPlan[destinationLine].lineNumber,
					)
				}
			} else {
				to = fmt.Sprintf(
					"%s_%d",
					executionPlan[destinationLine].instruction.operation,
					executionPlan[destinationLine].lineNumber,
				)
			}
		}

		graph = append(graph, fmt.Sprintf("%s -> %s", from, to))
	}

	return fmt.Sprintf("digraph G {\n%s\n}", strings.Join(graph, "\n\t"))
}

func generateExecutionPlan(bootCode []instruction) []*executionNode {
	execution := make([]*executionNode, len(bootCode))

	for line := 0; line < len(bootCode); line++ {
		if currentNode := execution[line]; currentNode != nil {
			if execution[line-1].instruction.operation != jmp {
				currentNode.parents = append(currentNode.parents, line-1)
			}
		} else {
			execution[line] = &executionNode{
				lineNumber:  line,
				instruction: bootCode[line],
				parents:     []int{},
			}

			if line > 0 {
				if execution[line-1].instruction.operation != jmp {
					execution[line].parents = append(execution[line].parents, line-1)
				}
			}
		}

		if bootCode[line].operation == jmp {
			var child int

			if bootCode[line].argument.sign == "+" {
				child = line + bootCode[line].argument.value
			} else {
				child = line - bootCode[line].argument.value
			}

			if child < len(bootCode) {
				currentNode := execution[line]
				currentNode.child = &child

				if childNode := execution[child]; childNode != nil {
					childNode.parents = append(childNode.parents, currentNode.lineNumber)
				} else {
					execution[child] = &executionNode{
						lineNumber:  child,
						instruction: bootCode[child],
						parents:     []int{line},
					}
				}
			}
		}
	}

	return execution
}

func getParentlessLeafNodes(executionPlan []*executionNode, endNode executionNode) []executionNode {
	leafNodes := []executionNode{}

	// start at the end
	for i := endNode.lineNumber; i >= 0; {
		if len(executionPlan[i].parents) == 0 {
			leafNodes = append(leafNodes, *executionPlan[i])
			break
		} else if len(executionPlan[i].parents) == 1 {
			i = executionPlan[i].parents[0]
		} else {
			for _, parent := range executionPlan[i].parents {
				parentLeafNodes := getParentlessLeafNodes(executionPlan, *executionPlan[parent])
				leafNodes = append(leafNodes, parentLeafNodes...)
			}
			break
		}
	}

	return leafNodes
}

func getJumpCandidates(executionPlan []*executionNode, leafNode executionNode) []executionNode {
	candidates := []executionNode{}
	startIndex := leafNode.lineNumber - 1

	if startIndex < 0 {
		return candidates
	}

	for i := executionPlan[startIndex].lineNumber; i >= 0; {
		if executionPlan[i].instruction.operation == jmp {
			candidates = append(candidates, *executionPlan[i])
			break
		} else if len(executionPlan[i].parents) == 0 {
			break
		} else if len(executionPlan[i].parents) == 1 {
			if executionPlan[i].instruction.operation == jmp {
				candidates = append(candidates, *executionPlan[i])
			}
			i = executionPlan[i].parents[0]
		} else {
			for _, parent := range executionPlan[i].parents {
				if executionPlan[parent].instruction.operation == jmp {
					candidates = append(candidates, *executionPlan[parent])
				} else {
					parentCandidates := getJumpCandidates(executionPlan, *executionPlan[parent])
					candidates = append(candidates, parentCandidates...)
				}
			}
			break
		}
	}

	return candidates
}

func getAnswerCountPartTwo(input string) int {
	// 0. load boot code
	bootCode := parseBootCode(input)

	// 1. generate execution graph
	execution := generateExecutionPlan(bootCode)

	// 1a. optionally, dump to graphviz digraph for debugging
	// graphviz := generateDigraph(execution)
	// fmt.Println(graphviz)

	// 2. identify cut-off (start at the last instruction and go up to all leaf nodes)
	leafNodes := getParentlessLeafNodes(execution, *execution[len(execution)-1])

	// 3. traverse up from each leafNode until we hit a jmp instruction, replace it with nop and run the program
	for _, leaf := range leafNodes {
		for _, candidate := range getJumpCandidates(execution, leaf) {
			// try to flip the jmp candidate to nop and run the code
			bootCode[candidate.lineNumber].operation = nop
			accumulator, loopDetected := runCode(bootCode)

			// if the code ran through to the end, we're done
			if !loopDetected {
				return accumulator
			} else {
				// otherwise, flip the nop back to jmp and continue on to the next candidate
				bootCode[candidate.lineNumber].operation = jmp
			}
		}
	}

	// none of the candidates produced valid code, so we give up
	return -1
}
