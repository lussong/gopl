package popcount

var pc [256]byte


// pc[i] is the population count of i.
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCount return the population count of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//LoopPopCount return the population count of x
//it is the loop version of popcount
func LoopPopCount(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(7*8))])
	}
	return count
}