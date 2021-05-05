// Popcount is a program that returns the number of set bits, that is, bit whose
// value is 1, in a uint64 value, which is called its population count.
package main

import (
	"fmt"
	"os"
	"popcount/algo"
	"strconv"
	"strings"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

func main() {
	args := os.Args[1:]
	choice, _ := strconv.Atoi(args[0])
	if choice > 3 || choice == 0 {
		stop()
	}
	for i, arg := range args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			str := fmt.Sprint(err)
			strColored := fmt.Sprint(string(colorRed), strings.TrimPrefix(str, "strconv.Atoi: parsing "), string(colorReset)) + "\n"
			if i == 0 && len(args) > 2 { // Print algorithm sentence if there is others command-line arguments.
				fmt.Printf("\nYou choose algorithm %v:\n\n", choice)
			}
			if i == 0 && len(args) == 2 { // Add new line if there is only on command-line argument (wrong).
				strColored = "\n" + strColored
			}
			fmt.Fprintf(os.Stderr, strColored)
			continue // Go to next arguments.
		}
		arg = fmt.Sprint(string(colorGreen), arg, string(colorReset))
		if i == 0 { // Make algorithm sentence  appears just once.
			fmt.Printf("\nYou choose algorithm %v:\n\n", choice)
		}
		switch choice {
		case 1:
			res := fmt.Sprint(string(colorGreen), algo.Sum(uint64(num)), string(colorReset))
			fmt.Printf("%v \t set bits in %v (%b)\n", res, arg, num)
		case 2:
			res := fmt.Sprint(string(colorGreen), algo.Loop(uint64(num)), string(colorReset))
			fmt.Printf("%v \t set bits in %v (%b)\n", res, arg, num)
		case 3:
			res := fmt.Sprint(string(colorGreen), algo.Shift(uint64(num)), string(colorReset))
			fmt.Printf("%v \t set bits in %v (%b)\n", res, arg, num)
		default:
			stop()
		}
	}
}

// Stop function stop the program if the algotithm exeeds 3.
func stop() {
	fmt.Println(string(colorRed), "\nPlease chose a number between 1 and 3 for algorithm", string(colorReset))
	os.Exit(0)
}
