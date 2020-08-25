package arrTag


//转换成以中心点为轴的四个数值之间的值交换
func Rotate(matrix [][]int) {
	n := len(matrix)
	//长度为(n+1)/2, 中间位置作标为(n-1)/2
	cols := (n + 1) / 2
	rows := n - cols
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] =
				matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}
