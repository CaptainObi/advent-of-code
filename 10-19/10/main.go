package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanned := scanner.Text()

		output = append(output, scanned)
	}

	opening_chars := []rune{'<', '{', '[', '('}
	//closing_chars := []rune{'>', '}', ']', ')'}

	var scores ArrInt
	for _, i := range output {

		var current_open []rune
		score := 0

		broken := false
		for _, j := range i {
			if contains(j, opening_chars) {
				current_open = append(current_open, j)
			} else if opposite(current_open[len(current_open)-1]) != j {
				//score += scoreF(j)
				broken = true
				break
			} else {
				current_open = current_open[:len(current_open)-1]
			}
		}
		if !broken {

			for i2, j := 0, len(current_open)-1; i2 < j; i2, j = i2+1, j-1 {
				current_open[i2], current_open[j] = current_open[j], current_open[i2]
			}

			fmt.Printf("current_open: %v\n", len(current_open))

			for _, l := range current_open {
				// fmt.Printf("score1: %v\n", score*5)

				score *= 5
				score += scoreF(opposite(l))
				//fmt.Printf("score2: %v\n", score)
			}

			//fmt.Printf("score: %v\n", score)

			scores = append(scores, score)
		}
		fmt.Println("===")
	}
	fmt.Println((scores))
	sort.Sort(scores)
	fmt.Printf("scores: %v\n", len(scores))
	fmt.Printf("scores[len(scores)/2]: %v\n", scores[len(scores)/2])

}

type ArrInt []int

func (a ArrInt) Len() int           { return len(a) }
func (a ArrInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ArrInt) Less(i, j int) bool { return a[i] < a[j] }

func contains(c rune, arr []rune) bool {
	for _, i := range arr {
		if i == c {
			return true
		}
	}
	return false
}

func opposite(r rune) rune {
	if r == '{' {
		return '}'
	} else if r == '[' {
		return ']'
	} else if r == '<' {
		return '>'
	} else {
		return ')'
	}
}

func scoreF(r rune) int {
	if r == ')' {
		return 1
	} else if r == ']' {
		return 2
	} else if r == '}' {
		return 3
	} else {
		return 4
	}
}
