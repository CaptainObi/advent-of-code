package main

import (
	"bufio"
	"fmt"
	"os"
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

	patterns := make(map[string]rune, len(output))
	template := output[0]

	for i, j := range output {
		if i >= 2 {
			split := strings.Split(j, " -> ")

			patterns[split[0]] = rune(split[1][0])
		}
	}

	// fmt.Printf("patterns: %v\n", patterns)

	for i := 0; i < 3; i++ {
		fmt.Println(template)
		//fmt.Printf("i: %v\n", i)
		n_template := ""
		for j := 0; j < len(template)-1; j++ {
			n_template += string(template[j])

			n_template += string(patterns[string(string(template[j])+string(template[j+1]))])
		}
		n_template += string(template[len(template)-1])

		template = n_template
		//fmt.Printf("len(template): %v\n", len(template))

		// freq_map := make(map[rune]int)

		// for _, i := range template {
		// 	freq_map[i]++
		// }

		// min, max := 10000000, 0

		// for _, j := range freq_map {
		// 	if min > j {
		// 		min = j
		// 	} else if max < j {
		// 		max = j
		// 	}
		// }

		// fmt.Println(max, min)
	}

	freq_map := make(map[rune]int)

	for _, i := range template {
		freq_map[i]++
	}

	min, max := 10000000, 0

	for _, j := range freq_map {
		if min > j {
			min = j
		} else if max < j {
			max = j
		}
	}

	fmt.Println(max - min)
}
