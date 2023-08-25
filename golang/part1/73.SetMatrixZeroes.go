package part1

func setZeroes(matrix [][]int) {
	//find zeros
	zeros := make([][]int, 0)
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}
	for _, zero := range zeros {
		for i := 0; i < n; i++ {
			matrix[zero[0]][i] = 0
		}
		for i := 0; i < m; i++ {
			matrix[i][zero[1]] = 0
		}
	}
}
