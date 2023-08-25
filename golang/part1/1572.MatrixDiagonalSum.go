package part1

func diagonalSum(mat [][]int) int {
	nl := len(mat)
	total := 0
	for i := 0; i < nl; i++ {
		total += mat[i][i] + mat[i][nl-1-i]
	}
	if nl%2 != 0 {
		total -= mat[(nl-1)/2][(nl-1)/2]
	}
	return total
}
