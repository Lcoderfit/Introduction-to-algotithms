package unstable_sort

import (
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
	var arr []int
	var res, ans []int

	// case1
	arr = []int{10, 4, 2, 7, 9}
	ans = []int{2, 4, 7, 9, 10}
	res = SelectSort(arr)
	if !reflect.DeepEqual(res, ans) {
		t.Errorf("res=%v, ans=%v", res, ans)
	} else {
		t.Log("case正确")
	}

	// case2
	arr = []int{6, 2, 9, 7, 4}
	ans = []int{2, 4, 6, 7, 9}
	res = SelectSort(arr)
	if !reflect.DeepEqual(res, ans) {
		t.Errorf("res=%v, ans=%v", res, ans)
	} else {
		t.Log("case正确")
	}
}
