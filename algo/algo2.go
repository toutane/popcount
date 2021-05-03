package algo

var pc2 [256]byte

func init() {
	for i := range pc2 {
		pc2[i] = pc2[i/2] + byte(i&1)
	}
}

func PopCount2(x uint64) int {
	var res byte
	for i := 0; i < 8; i++ {
		res += pc2[byte(x>>(i*8))]
	}
	return int(res)
}
