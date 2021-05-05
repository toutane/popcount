// This algorithm uses an init function to precompute a table of result, pc,
// for each possible 8-bit value so that the Sum function needn't take 64 steps
// but can return the sum of eight table lookups but with a loop instead of a
// sum .
package algo

var pc2 [256]byte

func init() {
	for i := range pc2 {
		pc2[i] = pc2[i/2] + byte(i&1)
	}
}

func Loop(x uint64) int {
	var res byte
	for i := 0; i < 8; i++ {
		res += pc2[byte(x>>(i*8))]
	}
	return int(res)
}
