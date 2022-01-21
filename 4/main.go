package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoCard struct {
	values [5][5]Value
	scored bool
}

type Value struct {
	number  int
	punched bool
}

func (b BingoCard) score(final int) int {
	unchecked := []int{}
	for _, i := range b.values {
		for _, j := range i {
			if j.punched == false {
				unchecked = append(unchecked, j.number)
			}
		}
	}

	sum := 0

	for _, i := range unchecked {
		sum += i
	}

	return sum * final

}

func (b *BingoCard) checkWin() bool {
	for _, i := range b.values {

		allChecked := true

		for _, j := range i {
			if j.punched == false {
				allChecked = false
			}
		}

		if allChecked == true {
			b.scored = true
			return true
		}
	}

	for j := range b.values[0] {
		allChecked := true
		for i := range b.values {
			if b.values[i][j].punched == false {
				allChecked = false
			}
		}
		if allChecked == true {
			b.scored = true
			return true
		}
	}

	return false
}

func (b *BingoCard) punch(i, j int) {
	b.values[i][j].punched = true
}

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

	numbers := []int{}

	for _, m := range strings.Split(output[0], ",") {
		conv, _ := strconv.Atoi(m)
		numbers = append(numbers, conv)
	}

	cards := output[2:]
	bingoCards := []*BingoCard{}

	for i := 0; i < len(cards); i += 6 {
		rows := [5][5]Value{}
		for j := 0; j < 5; j++ {

			if len(cards)-1 < i+j {
				break
			}

			row := [5]Value{}

			stringRow := strings.Split(cards[i+j], " ")

			if len(stringRow) != 0 {
				for n, l := range stringRow {
					p, _ := strconv.Atoi(l)
					row[n] = Value{p, false}
				}
			}

			rows[j] = row

		}

		bingoCards = append(bingoCards, &BingoCard{values: rows, scored: false})
	}

	for _, i := range numbers {
		for _, j := range bingoCards {
			for m, k := range j.values {
				for n, l := range k {
					if i == l.number {
						j.punch(m, n)
					}
				}
			}

			// fmt.Println("OOOO")
			// for _, o := range j.values {
			// 	p := []bool{}

			// 	for _, m := range o {
			// 		p = append(p, m.punched)
			// 	}

			// 	// fmt.Println(p)
			// }

			if j.scored == false {
				if j.checkWin() == true {
					fmt.Println("UWU")
					fmt.Println(j.score(i))

				}
			}
		}
	}

}
