package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	endpoints [2]Point
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

	lines := []Line{}

	for _, i := range output {
		split := strings.Split(i, " -> ")

		one := strings.Split(split[0], ",")
		two := strings.Split(split[1], ",")

		onex, _ := strconv.Atoi(one[0])
		oney, _ := strconv.Atoi(one[1])
		twox, _ := strconv.Atoi(two[0])
		twoy, _ := strconv.Atoi(two[1])

		lines = append(lines, Line{[2]Point{{onex, oney}, {twox, twoy}}})

	}

	fmt.Println(lines)

	grid := make([][]int, 1000)

	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	// for _, i := range grid {
	// 	fmt.Println(i)
	// }

	for _, i := range lines {
		one_x := i.endpoints[0].x
		one_y := i.endpoints[0].y
		two_x := i.endpoints[1].x
		two_y := i.endpoints[1].y

		if one_x == two_x {
			var high_y, low_y int

			if one_y > two_y {
				high_y = one_y
				low_y = two_y
			} else {
				high_y = two_y
				low_y = one_y
			}

			for j := low_y; j <= high_y; j++ {
				grid[j][one_x]++
			}

		} else if one_y == two_y {
			var high_x, low_x int

			if one_x > two_x {
				high_x = one_x
				low_x = two_x
			} else {
				high_x = two_x
				low_x = one_x
			}
			fmt.Println(high_x, low_x)

			for j := low_x; j <= high_x; j++ {
				grid[one_y][j]++
			}
		} else {
			// the line is diagonal
			// so it can be diagonal in 4 directions and we need to be able to handle of them

			between_x := generateBetween(one_x, two_x)
			between_y := generateBetween(one_y, two_y)

			for l := range between_x {
				grid[between_y[l]][between_x[l]]++
			}

		}
	}

	for _, i := range grid {
		fmt.Println(i)
	}

	var count int

	for _, i := range grid {
		for _, j := range i {
			if j > 1 {
				count++
			}
		}
	}

	fmt.Println(count, len(lines))
}

func generateBetween(a, b int) []int {

	if a < b {
		rs := []int{}
		for i := a; i <= b; i++ {
			rs = append(rs, i)
		}
		return rs
	} else {
		rs := []int{}
		for i := a; b <= i; i-- {
			rs = append(rs, i)
		}

		return rs
	}
}
