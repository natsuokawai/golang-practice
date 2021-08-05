package main

import (
	"fmt"
	"strconv"
)

func main() {
	var in string
	var a []int

	fmt.Scan(&in)
	n, _ := strconv.Atoi(in)

	var w string
	for i := 0; i < n; i++ {
		fmt.Scan(&w)
		nw, _ := strconv.Atoi(w)
		a = append(a, nw)
	}
}
