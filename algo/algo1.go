package algo

var pc1 [256]byte

func init() {
	for i := range pc1 {
		pc1[i] = pc1[i/2] + byte(i&1)
	}
}

func PopCount1(x uint64) int {
	return int(pc1[byte(x>>(0))] + pc1[byte(x>>(1*8))] + pc1[byte(x>>(2*8))] + pc1[byte(x>>(3*8))] + pc1[byte(x>>(4*8))] + pc1[byte(x>>(5*8))] + pc1[byte(x>>(6*8))] + pc1[byte(x>>(7*8))])
}
