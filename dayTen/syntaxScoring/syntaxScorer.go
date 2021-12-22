package syntaxscoring

import (
	"sort"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	return (*s)[len(*s)-1], true
}

var (
	incorrectScoring map[string]int = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	incompleteScoring map[string]int = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	chunkPairs map[string]string = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
)

func ScoreCorruptedLines(data []string) (incorrectScore int, incompleteScore int) {
	lineIncorectClosingCharacters := make([]string, 0)
	lineIncompleteClosingCharacters := make([]string, 0)
	for _, line := range data {
		var chunkStack Stack
		for idx := range line {
			stackPeek, isNotEmpty := chunkStack.Peek()
			if _, ok := chunkPairs[string(line[idx])]; !isNotEmpty || ok {
				chunkStack.Push(string(line[idx]))
			} else {
				if chunkPairs[stackPeek] != string(line[idx]) {
					lineIncorectClosingCharacters = append(lineIncorectClosingCharacters, string(line[idx]))
					break
				} else {
					chunkStack.Pop()
				}
			}

			if idx == len(line)-1 && len(chunkStack) > 0 {
				var sb strings.Builder
				for len(chunkStack) > 0 {
					val, _ := chunkStack.Pop()
					sb.WriteString(chunkPairs[val])
				}

				lineIncompleteClosingCharacters = append(lineIncompleteClosingCharacters, sb.String())
			}
		}
	}

	for _, val := range lineIncorectClosingCharacters {
		incorrectScore += incorrectScoring[val]
	}

	incompleteScores := make([]int, 0)
	for _, incLine := range lineIncompleteClosingCharacters {
		score := 0
		for _, ch := range incLine {
			score *= 5
			score += incompleteScoring[string(ch)]
		}

		incompleteScores = append(incompleteScores, score)
	}

	sort.Ints(incompleteScores)
	incompleteScore = incompleteScores[len(incompleteScores)/2]

	return
}
