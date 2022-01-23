package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func getVersion(bb []bool) (int64, error) {
	return getValue(bb[:3])
}

func getType(bb []bool) (int64, error) {
	return getValue(bb[3:6])
}

func stripLiteral(bb []bool) []bool {
	if bb[0] {
		// recursive
		return stripLiteral(
			bb[5:],
		)
	} else {
		return bb[5:]
	}
}

func stripHeader(bb []bool) []bool {
	return bb[6:]
}

// stripLiteralPacket include the header
func stripLiteralPacket(bb []bool) []bool {
	return stripLiteral(stripHeader(bb))
}

// handleOperator include the header
func handleOperator(bb []bool) ([]bool, int64, error) {
	v, err := getVersion(bb)
	// t, err := getType(bb)

	if err != nil {
		return nil, 0, err
	}

	stripped := stripHeader(bb)

	if stripped[0] {

		val, err := getValue(stripped[1:12])

		sum := v

		if err != nil {
			return nil, 0, err
		}

		stripped = stripped[12:]

		for i := 0; i < int(val); i++ {
			nSum := int64(0)
			stripped, nSum, err = handlePacket(stripped)

			sum += nSum
		}

		return stripped, sum, err
	} else {
		val, err := getValue(stripped[1:16])

		sum := v

		if err != nil {
			return nil, 0, err
		}

		fmt.Printf("stripped: %v\n", stripped[16:])

		stripped = stripped[16:]

		orgLen := len(stripped)

		for len(stripped) > orgLen-int(val) {
			nSum := int64(0)
			stripped, nSum, err = handlePacket(stripped)

			sum += nSum
		}

		return stripped, sum, err
	}
}

func handlePacket(bb []bool) ([]bool, int64, error) {
	fmt.Println("====")
	fmt.Println(getVersion(bb))

	if t, _ := getType(bb); t == 4 {
		return handleLiteral(bb)
	} else {
		return handleOperator(bb)
	}
}

func handleLiteral(bb []bool) ([]bool, int64, error) {

	b, e := getVersion(bb)
	r := stripLiteral(stripHeader(bb))

	return r, b, e
}
