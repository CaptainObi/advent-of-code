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

	fish := []int{}

	for _, i := range strings.Split(output[0], ",") {
		j, _ := strconv.Atoi(i)
		fish = append(fish, j)
	}

	fmt.Println(fish)

	fmap := make(map[int]int, 10)

	for _, i := range fish {
		fmap[i]++
	}

	for days := 0; days < 256; days++ {
		// for i := range fish {
		// 	fish[i]--
		// }

		// for i := range fish {
		// 	if fish[i] == -1 {
		// 		fish[i] = 6
		// 		fish = append(fish, 8)
		// 	}
		// }

		for i := 0; i <= 8; i++ {

			fmap[i-1] = fmap[i]
			fmap[i] = 0
		}

		negative := fmap[-1]
		fmap[8] += negative
		fmap[6] += negative
		fmap[-1] = 0

	}

	var sum int

	for _, i := range fmap {
		sum += i
	}

	fmt.Println(sum)
}
