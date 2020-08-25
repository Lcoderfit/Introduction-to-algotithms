package dp


func GetKthMagicNumber(k int) int {
	p3, p5, p7 := 0, 0, 0
	vec := make([]int, k)
	vec[0] = 1
	for i := 1; i < k; i++ {
		cur := min(vec[p3]*3, min(vec[p5]*5, vec[p7]*7))
		vec[i] = cur
		if cur == vec[p3]*3{
			p3++
		}
		if cur == vec[p5]*5 {
			p5++
		}
		if cur == vec[p7]*7 {
			p7++
		}
	}
	return vec[k - 1]
}