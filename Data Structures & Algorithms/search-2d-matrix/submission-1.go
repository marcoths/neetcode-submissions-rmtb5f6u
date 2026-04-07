func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lo, hi := 0, m*n-1

	for lo <= hi {
		mid := lo + (hi - lo ) / 2
		val := matrix[mid/n][mid%n]

		if val == target {
			return true
		} else if val < target {
			lo = mid + 1
		} else {
			hi = mid -1
		}
	}
	return false

}
