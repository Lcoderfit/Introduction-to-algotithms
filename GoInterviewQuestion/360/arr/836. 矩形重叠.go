package arr


func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec2[0] >= rec1[2] ||
		rec2[1] >= rec1[3] ||
		rec2[2] <= rec1[0] ||
		rec2[3] <= rec1[1])
}

//二：
func IsRectangleOverlap(rec1 []int, rec2 []int) bool {
	return (max(rec1[0], rec2[0]) < min(rec1[2], rec2[2])) &&
		(max(rec1[1], rec2[1]) < min(rec1[3], rec2[3]))
}