package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var arr []int
	var res, ans []int

	// case1
	arr = []int{5, 4, 1, 3, 2}
	ans = []int{1, 2, 3, 4, 5}
	res = BubbleSort(arr)
	if !reflect.DeepEqual(res, ans) {
		t.Errorf("res=%v, ans=%v", res, ans)
	} else {
		t.Log("case1正确")
	}

	// case2
	arr = []int{10, 6, 2}
	ans = []int{2, 6, 10}
	res = BubbleSort(arr)
	if !reflect.DeepEqual(res, ans) {
		t.Errorf("res=%v, ans=%v", res, ans)
	} else {
		t.Log("case2正确")
	}
}
