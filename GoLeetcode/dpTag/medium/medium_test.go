/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:

*/
package medium

import (
	"fmt"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	var m, n int
	for {
		fmt.Scanln(&m, &n)
		matrix := make([][]byte, 0)
		for i := 0; i < m; i++ {
			row := make([]byte, n)
			for j := 0; j < n; j++ {
				_, err := fmt.Scanf("%d", &row[j])
				if err != nil {
					t.Error(err)
					continue
				}
			}
			matrix[i] = row
		}
		ans := MaximalSquare(matrix)
		fmt.Println(ans)
	}
}
