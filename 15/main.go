package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	vis       bool
	dist, val int
	c         Coord
}

type Coord struct {
	i, j int
}

type NodeArr []Node
type NodeGrid [][]Node

func (a NodeArr) Len() int           { return len(a) }
func (a NodeArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NodeArr) Less(i, j int) bool { return a[i].dist < a[j].dist }

func (ng NodeGrid) findAllUnvisited() NodeArr {
	res := NodeArr{}

	for _, i := range ng {
		for _, j := range i {
			if !j.vis && j.dist != math.MaxInt {
				res = append(res, j)
			}
		}
	}
	return res
}

func (ng NodeGrid) notEmpty() bool {
	for _, i := range ng {
		for _, j := range i {
			if !j.vis {
				return true
			}
		}
	}

	return false
}

func (ng NodeGrid) getLowestDistanceNode() Node {
	nodes := ng.findAllUnvisited()

	sort.Sort(nodes)

	return nodes[0]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var output []string
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanned := scanner.Text()

		output = append(output, scanned)
	}

	// split the strings into a 2d array
	var o_grid NodeGrid

	for b, i := range output {
		split := strings.Split(i, "")

		var arr []Node

		for p, j := range split {
			conv, _ := strconv.Atoi(j)

			arr = append(arr, Node{val: conv, dist: math.MaxInt, vis: false, c: Coord{i: b, j: p}})
		}

		var e_arr []Node

		for j := 0; j < 5; j++ {
			for _, e := range arr {
				n_v := computeNewValue(e.val, j)
				e_arr = append(e_arr, Node{val: n_v, dist: math.MaxInt, vis: false, c: Coord{i: b, j: j*len(arr) + e.c.j}})
			}
		}

		o_grid = append(o_grid, e_arr)
	}

	var grid NodeGrid

	for i := 0; i < 5; i++ {
		for _, j := range o_grid {
			row := NodeArr{}
			for _, e := range j {
				n_v := computeNewValue(e.val, i)
				row = append(row, Node{val: n_v, dist: math.MaxInt, vis: false, c: Coord{j: e.c.j, i: len(grid)}})
			}

			grid = append(grid, row)
		}
	}

	for _, i := range grid {
		// fmt.Printf("i: %v\n", i)
		just_ints := []int{}

		for _, j := range i {
			just_ints = append(just_ints, j.val)
		}

		// f.WriteString(just_ints + "\n")
	}

	grid[0][0].val = 0
	grid[0][0].dist = 0

	for grid.notEmpty() {
		n := grid.getLowestDistanceNode()
		i := n.c.i
		j := n.c.j

		// fmt.Printf("\"h1\": %v\n", "h1")
		//	fmt.Println(len(grid) - 1)

		fmt.Printf("n: %v\n", n)
		f.WriteString(fmt.Sprintf("%v", n) + "\n")

		if i == len(grid)-1 && j == len(grid[i])-1 {
			break

		}

		neighbors := []Coord{}
		if i != 0 {
			neighbors = append(neighbors, Coord{i: i - 1, j: j})
		}
		if j != 0 {
			neighbors = append(neighbors, Coord{i: i, j: j - 1})
		}
		if i+1 != len(grid) {
			neighbors = append(neighbors, Coord{i: i + 1, j: j})
		}
		if j+1 != len(grid[i]) {
			neighbors = append(neighbors, Coord{i: i, j: j + 1})
		}

		for _, b := range neighbors {
			c := grid[b.i][b.j]

			if !c.vis {
				if c.dist > n.dist+n.val {
					grid[b.i][b.j].dist = n.dist + n.val
				}
			}
		}

		grid[i][j].vis = true

	}

	fmt.Println(grid[len(grid)-1][len(grid[0])-1].dist + grid[len(grid)-1][len(grid[0])-1].val)
}

func computeNewValue(old, increase int) int {
	res := old + increase

	for res > 9 {
		res -= 9
	}

	return res
}
