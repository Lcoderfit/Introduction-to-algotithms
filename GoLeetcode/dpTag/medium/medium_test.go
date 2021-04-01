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
		ans := MaximalSquare1(matrix)
		fmt.Println(ans)
	}
}

func TestNthUglyNumber2(t *testing.T) {
	n := 10
	ans := NthUglyNumber2(n)
	fmt.Println(ans)
}

func TestMinimumTotal(t *testing.T) {
	triangle1 := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	ans1 := 11
	res1 := MinimumTotal(triangle1)
	if res1 != ans1 {
		t.Errorf("res=%d, ans=%d", res1, ans1)
		return
	} else {
		t.Log("case1正确")
	}

	triangle2 := [][]int{
		{-10},
	}
	ans2 := -10
	res2 := MinimumTotal(triangle2)
	if res2 != ans2 {
		t.Errorf("res=%d, ans=%d", res2, ans2)
		return
	} else {
		t.Log("case2正确")
	}
}

func TestMinimumTotal2(t *testing.T) {
	triangle1 := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	ans1 := 11
	res1 := MinimumTotal2(triangle1)
	if res1 != ans1 {
		t.Errorf("res=%d, ans=%d", res1, ans1)
		return
	} else {
		t.Log("case1正确")
	}

	triangle2 := [][]int{
		{-10},
	}
	ans2 := -10
	res2 := MinimumTotal2(triangle2)
	if res2 != ans2 {
		t.Errorf("res=%d, ans=%d", res2, ans2)
		return
	} else {
		t.Log("case2正确")
	}
}
