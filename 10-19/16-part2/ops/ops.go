package ops

import "fmt"

func Sum(arr []int64) int64 {
	var res int64

	for _, i := range arr {
		res += i
	}

	return res
}

func Product(arr []int64) int64 {
	fmt.Printf("arr: %v\n", arr)
	var res int64 = 1

	for _, i := range arr {
		res = i * res
	}

	return res
}

func Minimum(arr []int64) int64 {
	lowest := arr[0]

	for _, i := range arr {
		if lowest > i {
			lowest = i
		}
	}

	return lowest
}

func Maximum(arr []int64) int64 {
	highest := arr[0]

	for _, i := range arr {
		if highest < i {
			highest = i
		}
	}

	return highest
}

func GreaterThan(arr [2]int64) int64 {
	if arr[0] > arr[1] {
		return 1
	} else {
		return 0
	}
}

func LessThan(arr [2]int64) int64 {
	if arr[0] < arr[1] {
		return 1
	} else {
		return 0
	}
}

func Equal(arr [2]int64) int64 {
	if arr[0] == arr[1] {
		return 1
	} else {
		return 0
	}
}
