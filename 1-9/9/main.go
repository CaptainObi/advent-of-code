package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// defines some methods on []int that are needed for sort.sort
type IntArr []int

func (a IntArr) Len() int           { return len(a) }
func (a IntArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntArr) Less(i, j int) bool { return a[i] < a[j] }

// a node of the 2d array

type Node struct {
	val int
	vis bool
}

// a set of coords to make it easy to pass

type Coord struct {
	i, j int
}

func main() {
	// all of this just reads the file
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

	// splits up the numbers in to a 2d array where all the nodes are unvisited

	numbers := [][]Node{}

	for _, i := range output {
		arr := []Node{}

		for _, j := range strings.Split(i, "") {
			conv, _ := strconv.Atoi(j)
			arr = append(arr, Node{val: conv, vis: false})
		}

		numbers = append(numbers, arr)
	}

	// list of basins, im using the type IntArr so that I can sort it, IntArr is just a type I made to replace []int

	basins := IntArr{}

	// loops over all the numbers in the 2D array
	for b := range numbers {
		for l := range numbers[b] {

			// the number has been visited or its a nine it skips
			if numbers[b][l].vis || numbers[b][l].val == 9 {
				continue
			}

			// debuging ignore
			fmt.Println(b, l, "===")
			// size of the basin
			basin_size := 0

			// this is the stack inits it with the current coords
			stack := []Coord{{i: b, j: l}}

			// until the stack isn't empty
			for len(stack) != 0 {
				// pops the current value off of the stack
				n_struct := stack[0]
				i := n_struct.i
				j := n_struct.j

				// pops the stack
				stack = stack[1:]

				// debug
				fmt.Println(i, j)

				// increases the size of the basin
				basin_size++
				numbers[i][j].vis = true

				// checks all the neighbors, if they are not equal to 9 and they haven't been visited yet,
				// it adds it to the stack and marks them at visited
				neighbors := []Coord{}
				if i != 0 {
					neighbors = append(neighbors, Coord{i: i - 1, j: j})
				}
				if j != 0 {
					neighbors = append(neighbors, Coord{i: i, j: j - 1})
				}
				if i+1 != len(numbers) {
					neighbors = append(neighbors, Coord{i: i + 1, j: j})
				}
				if j+1 != len(numbers[i]) {
					neighbors = append(neighbors, Coord{i: i, j: j + 1})
				}

				// goes through all teh neighbors and checks if they are nines or visited
				// if not it appends them to the stack and marks them as visited

				for _, d := range neighbors {
					if numbers[d.i][d.j].val != 9 && !numbers[d.i][d.j].vis {
						numbers[d.i][d.j].vis = true
						stack = append(stack, Coord{i: d.i, j: d.j})
					}
				}
			}

			// appends the current basin to the list of basins once the stack is over
			basins = append(basins, basin_size)

		}
	}

	// sorts all the basins by size
	sort.Sort(basins)
	fmt.Printf("basins: %v\n", basins)

	last := basins[basins.Len()-1]
	sec_last := basins[basins.Len()-2]
	third_last := basins[basins.Len()-3]

	fmt.Println(last * sec_last * third_last)
}
