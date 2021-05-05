package algo

var pc [256]byte // Population count (table lookup).

// Precompute the population count for comparison by algorithm.
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
