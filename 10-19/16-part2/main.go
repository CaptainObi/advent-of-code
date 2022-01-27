package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/CaptainObi/advent-of-code/10-19/16-part2/ops"
)

func convertHex(h string) []bool {

	var res []bool

	for _, i := range h {
		switch i {
		case '0':
			res = append(res, []bool{false, false, false, false}...)
		case '1':
			res = append(res, []bool{false, false, false, true}...)
		case '2':
			res = append(res, []bool{false, false, true, false}...)
		case '3':
			res = append(res, []bool{false, false, true, true}...)
		case '4':
			res = append(res, []bool{false, true, false, false}...)
		case '5':
			res = append(res, []bool{false, true, false, true}...)
		case '6':
			res = append(res, []bool{false, true, true, false}...)
		case '7':
			res = append(res, []bool{false, true, true, true}...)
		case '8':
			res = append(res, []bool{true, false, false, false}...)
		case '9':
			res = append(res, []bool{true, false, false, true}...)
		case 'A':
			res = append(res, []bool{true, false, true, false}...)
		case 'B':
			res = append(res, []bool{true, false, true, true}...)
		case 'C':
			res = append(res, []bool{true, true, false, false}...)
		case 'D':
			res = append(res, []bool{true, true, false, true}...)
		case 'E':
			res = append(res, []bool{true, true, true, false}...)
		case 'F':
			res = append(res, []bool{true, true, true, true}...)
		}
	}

	return res

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanned := scanner.Text()

	bits := convertHex(scanned)

	fmt.Printf("bits: %v\n", bits)

	fmt.Println(handleOperator(bits))
}

func binaryToBools(b string) []bool {
	var res []bool

	for _, i := range b {
		if i == '0' {
			res = append(res, false)
		} else {
			res = append(res, true)
		}
	}

	return res
}

func asBits(val uint64) []bool {
	return binaryToBools(strconv.FormatInt(int64(val), 2))
}

func getValue(bb []bool) (int64, error) {
	var str []rune

	for _, i := range bb {
		if i {
			str = append(str, '1')
		} else {
			str = append(str, '0')
		}
	}

	return strconv.ParseInt(string(str), 2, 64)
}

// func getVersion(bb []bool) (int64, error) {
// 	return getValue(bb[:3])
// }

func getType(bb []bool) (int64, error) {
	return getValue(bb[3:6])
}

// returns rest, res
func stripLiteral(bb []bool) ([]bool, []bool) {
	var res []bool
	res = append(res, bb[1:5]...)
	if bb[0] {
		bb, recur := stripLiteral(bb[5:])

		return bb, append(res, recur...)

	} else {
		return bb[5:], res
	}
}

func stripHeader(bb []bool) []bool {
	return bb[6:]
}

// stripLiteralPacket include the header
// func stripLiteralPacket(bb []bool) []bool {
// 	return stripLiteral(stripHeader(bb))
// }

// handleOperator include the header
func handleOperator(bb []bool) ([]bool, int64, error) {
	// _, err := getVersion(bb)
	t, err := getType(bb)
	var values []int64
	if err != nil {
		return nil, 0, err
	}

	stripped := stripHeader(bb)

	if stripped[0] {

		val, err := getValue(stripped[1:12])

		if err != nil {
			return nil, 0, err
		}

		stripped = stripped[12:]

		for i := 0; i < int(val); i++ {
			nSum := int64(0)
			stripped, nSum, err = handlePacket(stripped)

			values = append(values, nSum)
		}
	} else {
		val, err := getValue(stripped[1:16])

		if err != nil {
			return nil, 0, err
		}

		stripped = stripped[16:]

		orgLen := len(stripped)

		for len(stripped) > orgLen-int(val) {
			nSum := int64(0)
			stripped, nSum, err = handlePacket(stripped)

			values = append(values, nSum)
		}
	}

	switch t {
	case 0:
		return stripped, ops.Sum(values), err
	case 1:
		return stripped, ops.Product(values), err
	case 2:
		return stripped, ops.Minimum(values), err
	case 3:
		return stripped, ops.Maximum(values), err
	case 5:
		return stripped, ops.GreaterThan([2]int64{values[0], values[1]}), err
	case 6:
		return stripped, ops.LessThan([2]int64{values[0], values[1]}), err
	case 7:
		return stripped, ops.Equal([2]int64{values[0], values[1]}), err
	default:
		return nil, 0, err
	}
}

func handlePacket(bb []bool) ([]bool, int64, error) {
	if t, _ := getType(bb); t == 4 {
		return handleLiteral(bb)
	} else {
		return handleOperator(bb)
	}
}

func handleLiteral(bb []bool) ([]bool, int64, error) {

	// _, e := getVersion(bb)
	r, val := stripLiteral(stripHeader(bb))
	res, e := getValue(val)
	return r, res, e
}
