package main

import (
	"fmt"
	"os"
	"popcount/algo"
	"strconv"
)

const (
	colorReset = "\033[0m"
	colorGreen = "\033[32m"
)

func main() {
	args := os.Args[1:]
	choice := args[0]
	fmt.Printf("You choose the algorithm %v:\n\n", choice)
	for _, arg := range args[1:] {
		num, err := strconv.Atoi(arg)
		arg = fmt.Sprint(string(colorGreen), arg, string(colorReset))
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
			os.Exit(1)
		}
		switch choice {
		case "1":
			res := fmt.Sprint(string(colorGreen), algo.PopCount1(uint64(num)), string(colorReset))
			fmt.Printf("%v \t set bits in %v (%b)\n", res, arg, num)
		case "2":
			res := fmt.Sprint(string(colorGreen), algo.PopCount2(uint64(num)), string(colorReset))
			fmt.Printf("%v \t set bits in %v (%b)\n", res, arg, num)
		case "3":
			res := fmt.Sprint(string(colorGreen), algo.PopCount3(uint64(num)), string(colorReset))
			fmt.Printf("%v \t set bits in %v (%b)\n", res, arg, num)
		default:
			stop()
		}
	}
}

func stop() {
	fmt.Println("Please chose a numbr between 1 and 3")
	os.Exit(0)
}
