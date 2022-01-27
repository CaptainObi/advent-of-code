package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	octopi := [][]int{}

	for _, i := range output {
		split := strings.Split(i, "")
		arr := []int{}
		for _, j := range split {
			conv, _ := strconv.Atoi(j)
			arr = append(arr, conv)
		}

		octopi = append(octopi, arr)
	}

	flashes := 0

	bloop := true
	iterations := 0

	for bloop {

		// checks if they are all flashed and if so breaks the loop
		allZero := true

		for _, i := range octopi {
			for _, j := range i {
				if j != 0 {
					allZero = false
				}
			}
		}

		if allZero {
			bloop = false
			break
		}

		octopi = increaseOcto(octopi)

		for {
			i, j := findFirstFlash(octopi)
			if i == -100 {
				break
			}

			fmt.Printf("j: %v\n", j)
			if i != 0 {
				octopi[i-1][j]++
			}
			if j != 0 {
				octopi[i][j-1]++
			}

			if i != 0 && j != 0 {
				octopi[i-1][j-1]++
			}

			if i != 0 && j+1 != len(octopi[i]) {
				octopi[i-1][j+1]++
			}

			if i+1 != len(octopi) {
				octopi[i+1][j]++
			}

			if j+1 != len(octopi[i]) {
				octopi[i][j+1]++
			}

			if i+1 != len(octopi) && j != 0 {
				octopi[i+1][j-1]++
			}

			if i+1 != len(octopi) && j+1 != len(octopi[i]) {
				octopi[i+1][j+1]++
			}

			octopi[i][j] = 200000
			flashes++
		}

		for i := range octopi {
			for j := range octopi[i] {
				if octopi[i][j] >= 200000 {
					octopi[i][j] = 0
				}
			}
		}
		iterations++
	}

	for _, i := range octopi {
		fmt.Println(i)
	}

	fmt.Printf("flashes: %v\n", flashes)
	fmt.Printf("iterations: %v\n", iterations)
}

func increaseOcto(grid [][]int) [][]int {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j]++
		}
	}
	return grid
}

func findFirstFlash(grid [][]int) (int, int) {
	for u, i := range grid {
		for l, j := range i {
			if 9 < j && j < 200000 {
				return u, l
			}
		}
	}

	return -100, -100
}
