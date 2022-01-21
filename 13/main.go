package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fold struct {
	is_y  bool
	value int
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

	coords := [][2]int{}
	folds := []Fold{}

	for _, i := range output {
		if len(i) > 0 {
			if i[0] != 'f' {
				split := strings.Split(i, ",")
				conva, _ := strconv.Atoi(split[1])
				convb, _ := strconv.Atoi(split[0])

				coords = append(coords, [2]int{conva, convb})
			} else {
				split := strings.Split(i, " ")

				second_split := strings.Split(split[2], "=")

				var fold Fold

				if second_split[0] == "y" {
					fold.is_y = true
				} else {
					fold.is_y = false
				}

				conv, _ := strconv.Atoi(second_split[1])

				fold.value = conv

				folds = append(folds, fold)
			}
		}
	}

	// fmt.Println(coords)
	// fmt.Println(folds)

	dots := make([][]bool, 1500)

	for i := range dots {
		dots[i] = make([]bool, 1500)
	}

	for _, i := range coords {
		dots[i[0]][i[1]] = true
	}

	for _, i := range folds {
		fold(dots, i)
	}

	// printPaper(dots)

	// sum := 0
	// for _, i := range dots {
	// 	for _, j := range i {
	// 		if j {
	// 			sum++
	// 		}
	// 	}
	// }
	// fmt.Printf("sum: %v\n", sum)

	hx, hy := findLen(dots)

	fmt.Printf("hx: %v\n", hx)
	fmt.Printf("hy: %v\n", hy)

	n_dots := [][]bool{}

	for i := 0; i <= hy; i++ {

		row := []bool{}

		for j := 0; j <= hx; j++ {
			fmt.Println(i, j)
			fmt.Println(dots[i][j])
			row = append(row, dots[i][j])
		}

		n_dots = append(n_dots, row)
	}

	printPaper(n_dots)

}

func fold(paper [][]bool, f Fold) [][]bool {
	if !f.is_y {
		highest_x, _ := findLen(paper)
		for i := 0; i < (highest_x+2)/2; i++ {
			for j := range paper {
				if paper[j][f.value+i] {
					paper[j][f.value-i] = true
					paper[j][f.value+i] = false
				}

			}
		}
	} else {
		_, highest_y := findLen(paper)
		for j := 0; j < (highest_y+2)/2; j++ {
			for i := range paper[0] {
				if paper[f.value+j][i] {
					paper[f.value-j][i] = true
				}

				paper[f.value+j][i] = false
			}
		}

	}

	return paper
}

func findLen(paper [][]bool) (int, int) {
	highest_y := 0
	highest_x := 0
	for i := range paper {
		for j, p := range paper[i] {
			if p && i > highest_y {
				highest_y = i
			}
			if p && j > highest_x {
				highest_x = j
			}
		}
	}

	return highest_x, highest_y

}

func printPaper(dots [][]bool) {
	for _, i := range dots {
		chars := []string{}
		for _, j := range i {
			if j {
				chars = append(chars, "O")
			} else {
				chars = append(chars, ".")
			}
		}

		fmt.Println(chars)
	}
}
