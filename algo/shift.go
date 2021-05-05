// This algorithm counts bits by shifting its arguments through 64 bit
// positions, testing the rightmost bit each time.
package algo

func Shift(x uint64) int {
	var res int
	input := x
	for i := 0; i < 64; i++ {
		if byte(input&1) == 1 {
			res++
		}
		input = uint64(input >> 1)
	}
	return res
}
