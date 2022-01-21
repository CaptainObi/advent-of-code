package main

import (
	"bufio"
	"bytes"
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

	// final_map := make(map[string]int, 10)

	// final_map["fe"] = 1
	// final_map["afgcd"] = 2
	// final_map["afged"] = 3
	// final_map["bgfe"] = 4
	// final_map["afgcd"] = 5
	// final_map["abcdeg"] = 6
	// final_map["aef"] = 7
	// final_map["abcdefg"] = 8
	// final_map["abgfed"] = 9
	// final_map["abcdef"] = 0

	var sum int

	for _, i := range output {
		e := strings.Split(i, " | ")
		wires, display := strings.Split(e[0], " "), strings.Split(e[1], " ")

		//seg := make(map[rune]rune, 7)
		num := make(map[int]string, 10)
		sixers := []string{}
		fives := []string{}

		for _, j := range wires {
			if len(j) == 2 {
				// 1
				num[1] = j
			} else if len(j) == 3 {
				// 7
				num[7] = j
			} else if len(j) == 7 {
				// 8
				num[8] = j
			} else if len(j) == 4 {
				num[4] = j
			} else if len(j) == 5 {
				fives = append(fives, j)
			} else if len(j) == 6 {
				sixers = append(sixers, j)
			}
		}

		fives_c := make(map[rune]int, 10)

		for _, j := range fives {
			for _, l := range j {
				fives_c[l]++
			}
		}

		top_lines := []rune{}
		for l, p := range fives_c {
			if p >= 3 {
				top_lines = append(top_lines, l)
			}
		}

		for _, j := range sixers {
			for _, l := range top_lines {
				if !strings.ContainsRune(j, l) {
					num[0] = j
					break
				}
			}
		}

		new_sixers := []string{}

		for _, j := range sixers {
			if j != num[0] {
				new_sixers = append(new_sixers, j)
			}
		}

		for u, j := range new_sixers {
			for _, l := range num[1] {
				if !strings.ContainsRune(j, l) {
					num[6] = j
					if u == 0 {
						num[9] = new_sixers[1]
					} else {
						num[9] = new_sixers[0]
					}
					break
				}
			}
		}

		for _, j := range fives {

			if strings.ContainsRune(j, []rune(num[1])[0]) && strings.ContainsRune(j, []rune(num[1])[1]) {

				num[3] = j
				break

			}
		}

		new_fives := []string{}

		for _, j := range fives {
			if j != num[3] {
				new_fives = append(new_fives, j)
			}
		}

		// remove the | from the 4 and turn it into a |_ then match it agains the 5 to get those values :)

		searcher := ""

		for _, j := range num[4] {
			if !strings.ContainsRune(num[1], j) {
				searcher = strings.Join([]string{searcher, string(j)}, "")
			}
		}

		for u, j := range new_fives {
			for _, l := range searcher {
				if !strings.ContainsRune(j, l) {
					num[2] = j
					if u == 0 {
						num[5] = new_fives[1]
					} else {
						num[5] = new_fives[0]
					}
					break
				}
			}
		}

		// fmt.Println(wires, num, display)

		d_arr := []int{}

		for _, j := range display {
			for u, l := range num {
				if unorderedEqual([]rune(l), []rune(j)) {
					d_arr = append(d_arr, u)
				}
			}
		}
		// fmt.Printf("d_arr: %v\n", d_arr)

		var buf bytes.Buffer

		for i := range d_arr {
			buf.WriteString(fmt.Sprintf("%d", d_arr[i]))
		}

		final, _ := strconv.Atoi(buf.String())
		sum += final

		fmt.Println(final)
	}

	/*
		Okay so this is how we are going to label this shit
					aa
				bb	ff
				bb 	ff
					gg
				cc	ee
				cc 	ee
					dd
	*/
	fmt.Println(sum)
}

func unorderedEqual(first, second []rune) bool {
	if len(first) != len(second) {
		return false
	}
	exists := make(map[rune]bool)
	for _, value := range first {
		exists[value] = true
	}
	for _, value := range second {
		if !exists[value] {
			return false
		}
	}
	return true
}
