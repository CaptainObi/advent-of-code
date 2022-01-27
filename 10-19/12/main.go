package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var paths [][]string

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

	var nodes []string

	for _, i := range output {
		split := strings.Split(i, "-")

		contains := false

		for _, j := range nodes {
			if j == split[0] {
				contains = true
			}
		}

		if !contains {
			nodes = append(nodes, split[0])
		}

		contains = false

		for _, j := range nodes {
			if j == split[1] {
				contains = true
			}
		}

		if !contains {
			nodes = append(nodes, split[1])
		}

	}

	var g Graph
	g.graph = make(map[string][]string, 10)
	g.vertices = make([]string, 10)

	for _, i := range nodes {
		g.addVert(i)
	}

	for _, i := range output {
		split := strings.Split(i, "-")
		g.addEdge(split[0], split[1])
	}

	ch := make(chan []string)

	go g.findALlPaths("start", "end", ch)

	var paths [][]string

	for v := range ch {
		fmt.Printf("v: %v\n", v)
		paths = append(paths, v)
	}

	fmt.Printf("len(paths): %v\n", len(paths))

}

type Graph struct {
	vertices []string
	graph    map[string][]string
}

func (g *Graph) addVert(name string) {
	g.vertices = append(g.vertices, name)
}

func (g *Graph) addEdge(u, v string) {
	g.graph[u] = append(g.graph[u], v)
	g.graph[v] = append(g.graph[v], u)
}

func (g *Graph) findAllPathsUtil(u, d string, visited map[string]int, path []string, ch chan []string) []string {
	visited[u]++
	path = append(path, u)

	if u == d {
		var nArr []string
		for _, i := range path {
			nArr = append(nArr, i)
		}

		ch <- nArr

	} else {
		for _, i := range g.graph[u] {

			// you can visit a single small once

			already_visited_small := false

			for l, p := range visited {
				if strings.ToLower(l) == l {
					if p >= 2 {

						already_visited_small = true
					}
				}
			}
			if i == "aa" {
				fmt.Println(visited[i])
				fmt.Println(visited[i] == 0, !already_visited_small, strings.ToLower(i) != i)
			}
			if (visited[i] == 0 || !already_visited_small || strings.ToLower(i) != i) && i != "start" {
				path = g.findAllPathsUtil(i, d, visited, path, ch)
			}
		}
	}

	path = (path)[:len(path)-1]
	visited[u] = visited[u] - 1
	return path
}

func (g *Graph) findALlPaths(s, d string, ch chan []string) {
	defer close(ch)
	visited := make(map[string]int, 10)
	path := []string{}

	g.findAllPathsUtil(s, d, visited, path, ch)
}
