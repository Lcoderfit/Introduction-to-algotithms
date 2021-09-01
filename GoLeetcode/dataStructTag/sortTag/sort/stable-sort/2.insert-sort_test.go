package stable_sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	var arr []int
	var res, ans []int

	// case1
	arr = []int{5, 7, 2, 0, 3, 6, 9}
	res = InsertSort(arr)
	ans = []int{0, 2, 3, 5, 6, 7, 9}
	if !reflect.DeepEqual(res, ans) {
		t.Errorf("res=%v, ans=%v", res, ans)
	} else {
		t.Log("case1正确")
	}

	// case2
	arr = []int{8, 0, 4, 3, 2, 7}
	res = InsertSort(arr)
	ans = []int{0, 2, 3, 4, 7, 8}
	if !reflect.DeepEqual(res, ans) {
		t.Errorf("res=%v, ans=%v", res, ans)
	} else {
		t.Log("case1正确")
	}
}
