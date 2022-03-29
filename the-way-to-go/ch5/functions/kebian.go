package functions

func Sum(vals ...int) int {
	var total int
	for _, val := range vals {
		total += val
	}
	return total
}
