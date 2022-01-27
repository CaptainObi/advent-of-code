package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var twoD [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanned := scanner.Text()

		twoD = append(twoD, []rune(scanned))
	}

	g_array := twoD
	e_array := twoD

	for i := range g_array[0] {
		if len(g_array) == 1 {
			break
		}

		ones := 0
		zeros := 0

		for j := range g_array {

			if g_array[j][i] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		var g rune

		if ones >= zeros {
			g = '1'
		} else {
			g = '0'
		}

		var (
			new_g [][]rune
		)

		for _, j := range g_array {
			if j[i] == g {
				new_g = append(new_g, j)
			}
		}

		g_array = new_g
		fmt.Println(g_array)
	}

	fmt.Println(g_array, len(e_array), "e")

	for i := range e_array[0] {
		if len(e_array) == 1 {
			break
		}

		ones := 0
		zeros := 0

		for j := range e_array {
			if e_array[j][i] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		var e rune

		if ones >= zeros {
			e = '0'
		} else {
			e = '1'
		}

		var (
			new_e [][]rune
		)

		fmt.Println(e)

		for _, j := range e_array {
			if j[i] == e {
				new_e = append(new_e, j)
			}
		}

		e_array = new_e
		fmt.Println(e_array, "O")
	}

	finalA, _ := strconv.ParseInt(string(g_array[0]), 2, 64)
	finalB, _ := strconv.ParseInt(string(e_array[0]), 2, 64)

	fmt.Println(finalA * finalB)

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()

// 	var twoD [][]rune

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		scanned := scanner.Text()

// 		twoD = append(twoD, []rune(scanned))
// 	}

// 	g := []rune{}
// 	e := []rune{}

// 	for i := range twoD[0] {
// 		ones := 0
// 		zeros := 0

// 		for j := range twoD {
// 			if twoD[j][i] == '1' {
// 				ones++
// 			} else {
// 				zeros++
// 			}
// 		}

// 		if ones > zeros {
// 			g = append(g, '1')
// 			e = append(e, '0')
// 		} else {
// 			g = append(g, '0')
// 			e = append(e, '1')
// 		}
// 	}

// 	fmt.Println(string(g), string(e))

// 	g_array := twoD

// 	for len(g_array) != 1 {
// 		for i := 0; i < len(g_array[0]); i++ {
// 			for j := 0; j < len(g_array); j++ {
// 				if len(g_array) != 1 {
// 					if g_array[j][i] != g[i] {
// 						g_array = RemoveIndex(g_array, j)
// 					}
// 				}
// 			}
// 			fmt.Println("Next col", g_array)
// 		}
// 	}

// 	g_str =

// 	e_array := twoD

// 	for len(e_array) != 1 {
// 		for i := 0; i < len(e_array[0]); i++ {
// 			for j := 0; j < len(e_array); j++ {
// 				if len(e_array) != 1 {
// 					if e_array[j][i] != e[i] {
// 						e_array = RemoveIndex(e_array, j)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println(string(g_array[0]), string(e_array[0]))

// 	// fmt.Println(string(g), string(e))

// 	// newG, _ := strconv.ParseInt(string(g), 2, 64)
// 	// newE, _ := strconv.ParseInt(string(e), 2, 64)

// 	// fmt.Println(newG * newE)

// 	// if err := scanner.Err(); err != nil {
// 	// 	fmt.Println(err)
// 	// }
// }

// func RemoveIndex(s [][]rune, index int) [][]rune {
// 	return append(s[:index], s[index+1:]...)
// }
