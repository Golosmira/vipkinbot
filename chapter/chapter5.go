package main

type Operate func(int, int) int

type Expression struct {
	X, Y      int
	Operation Operate
}

var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}
