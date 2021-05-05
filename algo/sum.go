// This algorithm uses an init function to precompute a table of result, pc,
// for each possible 8-bit value so that the Sum function needn't take 64 steps
// but can return the sum of eight table lookups.
package algo

var pc1 [256]byte

func init() {
	for i := range pc1 {
		pc1[i] = pc1[i/2] + byte(i&1)
	}
}

func Sum(x uint64) int {
	return int(pc1[byte(x>>(0))] + pc1[byte(x>>(1*8))] + pc1[byte(x>>(2*8))] + pc1[byte(x>>(3*8))] + pc1[byte(x>>(4*8))] + pc1[byte(x>>(5*8))] + pc1[byte(x>>(6*8))] + pc1[byte(x>>(7*8))])
}
