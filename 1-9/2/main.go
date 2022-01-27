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

	depth := 0
	distance := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		split := strings.Split(scanner.Text(), " ")
		command := split[0]
		dist, _ := strconv.Atoi(split[1])

		if command == "down" {
			fmt.Println("down")

			aim += dist
		} else if command == "up" {
			fmt.Println("up")
			aim -= dist
		} else {
			fmt.Println("forward")
			distance += dist
			depth += aim * dist
		}
	}

	fmt.Println(distance, depth, distance*depth)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
