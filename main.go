package main

import (
	"fmt"
	"os"
	"popcount/algo"
	"strconv"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
			os.Exit(1)
		}
		// fmt.Printf("%v: %v", num, algo.PopCount1(uint64(num)))
		//fmt.Printf("%v: %v", num, algo.PopCount2(uint64(num)))
		fmt.Printf("%v: %v", num, algo.PopCount3(uint64(num)))
	}
}
