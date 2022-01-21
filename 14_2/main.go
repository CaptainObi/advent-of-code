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

	pairs := make(map[string]int, 10)

	for j := 0; j < len(template)-1; j++ {
		pairs[string(template[j])+string(template[j+1])]++

	}

	for j := 0; j < 40; j++ {
		// fmt.Printf("pairs: %v\n", pairs)
		n_pairs := make(map[string]int)
		for i, v := range pairs {
			n_char := patterns[i]

			// pairs[i] = 0

			// fmt.Println(i, string(patterns[i]))
			// fmt.Println(string(i[0])+string(n_char), string(n_char)+string(i[1]))

			n_pairs[string(i[0])+string(n_char)] += v
			n_pairs[string(n_char)+string(i[1])] += v
		}

		pairs = n_pairs
	}

	fmt.Printf("pairs: %v\n", pairs)

	freq_map := make(map[rune]int, 20)

	for i, j := range pairs {
		freq_map[rune(i[0])] += j
	}

	freq_map[rune(template[len(template)-1])]++

	fmt.Printf("freq_map: %v\n", freq_map)

	min, max := 10000000000000000, 0

	for _, j := range freq_map {
		if min > j {
			min = j
		} else if max < j {
			max = j
		}
	}

	fmt.Println(max - min)
	fmt.Println(max, min)

}
