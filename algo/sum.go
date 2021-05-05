// This algorithm uses an init function to precompute a table of result, pc,
// for each possible 8-bit value so that the Sum function needn't take 64 steps
// but can return the sum of eight table lookups.
package algo

func Sum(x uint64) int {
	return int(pc[byte(x>>(0))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] + pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] + pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
}
