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

func TestWiggleMaxLength(t *testing.T) {
	var ans, res int
	var nums []int
	nums = []int{1, 7, 4, 9, 2, 5}
	ans = 6
	res = WiggleMaxLength(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	nums = []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	ans = 7
	res = WiggleMaxLength(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	nums = []int{0, 0}
	ans = 1
	res = WiggleMaxLength(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}
}

func TestWiggleMaxLength2(t *testing.T) {
	var ans, res int
	var nums []int
	nums = []int{1, 7, 4, 9, 2, 5}
	ans = 6
	res = WiggleMaxLength2(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	nums = []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	ans = 7
	res = WiggleMaxLength2(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	nums = []int{0, 0}
	ans = 1
	res = WiggleMaxLength2(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}
}

func TestNumMatrix_SumRegion(t *testing.T) {
	var matrix [][]int
	matrix = [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	var ans, res int
	nm := ConstructorNumMatrix(matrix)
	//ans = 8
	//res = nm.SumRegion(2, 1, 4, 3)
	//if res != ans {
	//	t.Errorf("res=%d, ans=%d", res, ans)
	//} else {
	//	t.Log("case正确")
	//}

	ans = 11
	res = nm.SumRegion(1, 1, 2, 2)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	//ans = 12
	//res = nm.SumRegion(1, 2, 2, 4)
	//if res != ans {
	//	t.Errorf("res=%d, ans=%d", res, ans)
	//} else {
	//	t.Log("case正确")
	//}
}

func TestCountSubstrings(t *testing.T) {
	var s string
	var ans, res int
	s = "aaa"
	ans = 6
	res = CountSubstrings(s)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}
}

func TestMaxSumDivThree(t *testing.T) {
	var nums []int
	var res, ans int
	nums = []int{3, 6, 5, 1, 8}
	ans = 18
	res = MaxSumDivThree(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	nums = []int{4}
	ans = 0
	res = MaxSumDivThree(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	nums = []int{1, 2, 3, 4, 4}
	ans = 12
	res = MaxSumDivThree(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}
}

func TestMaxSumDivThree2(t *testing.T) {
	var nums []int
	var res, ans int
	nums = []int{3, 6, 5, 1, 8}
	ans = 18
	res = MaxSumDivThree2(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	nums = []int{4}
	ans = 0
	res = MaxSumDivThree2(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}

	nums = []int{1, 2, 3, 4, 4}
	ans = 12
	res = MaxSumDivThree2(nums)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}
}

func TestFindMaxForm(t *testing.T) {
	var strList []string
	var m, n int
	m, n = 5, 3
	var ans, res int
	strList = []string{"10", "0001", "111001", "1", "0"}
	ans = 4
	res = FindMaxForm(strList, m, n)
	if res != ans {
		t.Errorf("res=%d, ans=%d", res, ans)
	} else {
		t.Log("case正确")
	}
}
