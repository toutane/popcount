package algo

func PopCount3(x uint64) int {
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
