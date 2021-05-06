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

var max, maxB int

func main() {
	args := os.Args[1:]
	for _, arg := range args[1:] { // Set the number with the most bits (max) and the one with the most digits in base 10. (maxB)
		num, _ := strconv.Atoi(arg)
		str := fmt.Sprintf("(%b)", num)
		if len(str) > max {
			max = len(str)
		}
		if num > maxB {
			maxB = num
		}
	}
	choice, _ := strconv.Atoi(args[0])
	if choice > 3 || choice == 0 {
		stop()
	}
	for i, arg := range args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			str := fmt.Sprint(err)
			strColored := "  " + fmt.Sprint(string(colorRed), strings.TrimPrefix(str, "strconv.Atoi: parsing "), string(colorReset)) + "\n"
			if i == 0 && len(args) > 2 { // Print algorithm sentence if there are other command-line arguments.
				fmt.Printf("\nYou choose algorithm %v:\n\n", choice)
			}
			if i == 0 && len(args) == 2 { // Add new line if there is only one command-line argument (but badly formatted).
				strColored = "\n" + strColored
			}
			fmt.Fprintf(os.Stderr, strColored)
			continue // Go to next arguments.
		}
		if i == 0 { // Make algorithm sentence  appears just once.
			fmt.Printf("\nYou choose algorithm %v:\n\n", choice)
		}
		switch choice {
		case 1:
			res := fmt.Sprint(string(colorGreen), algo.Sum(uint64(num)), string(colorReset))
			fmt.Print(formatOutput(res, num))
		case 2:
			res := fmt.Sprint(string(colorGreen), algo.Loop(uint64(num)), string(colorReset))
			fmt.Print(formatOutput(res, num))
		case 3:
			res := fmt.Sprint(string(colorGreen), algo.Shift(uint64(num)), string(colorReset))
			fmt.Print(formatOutput(res, num))
		default:
			stop()
		}
	}
}

// CountDig returns the number of digits if a integer.
func countDig(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}

// FormatOutput formats the output with color and padding.
func formatOutput(res string, num int) string {
	arg := fmt.Sprint(string(colorGreen), num, string(colorReset))
	byteStr := fmt.Sprintf("(%b)", num)
	pad := max - len(byteStr) + countDig(maxB) - countDig(num) + 3
	padStr := strings.Repeat(" ", pad) + byteStr
	return fmt.Sprintf("   %v \t set bits in %v %s\n", res, arg, padStr)
}

// Stop function stop the program if the algotithm exeeds 3.
func stop() {
	fmt.Println(string(colorRed), "\nPlease chose a number between 1 and 3 for algorithm", string(colorReset))
	os.Exit(0)
}
