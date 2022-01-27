package main

import (
	"bufio"
	"fmt"
	"math"
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

	crabs := []int{}

	for _, i := range strings.Split(output[0], ",") {
		j, _ := strconv.Atoi(i)
		crabs = append(crabs, j)
	}

	fmt.Println(crabs)

	lowest := 100
	highest := 0

	for _, i := range crabs {
		if lowest > i {
			lowest = i
		}
		if highest < i {
			highest = i
		}
	}

	fmt.Println(lowest, highest)

	fuelValues := []int{}

	for i := lowest; i <= highest; i++ {
		fuel := 0

		for _, j := range crabs {
			diff := int(math.Abs(float64(i - j)))

			fuel += generateTriangle(diff)
		}

		fmt.Printf("i: %v\n", i)
		fmt.Printf("fuel: %v\n", fuel)

		fuelValues = append(fuelValues, fuel)
	}

	bestFuel := 1000000000000

	for _, j := range fuelValues {
		if bestFuel > j {
			bestFuel = j
		}
	}

	fmt.Printf("bestFuel: %v\n", bestFuel)

	fmt.Print(generateTriangle(1), generateTriangle(2), generateTriangle(3), generateTriangle(4))
}

func generateTriangle(steps int) int {

	sum := 0

	for i := 0; i <= steps; i++ {
		sum += i
	}

	return sum
}
