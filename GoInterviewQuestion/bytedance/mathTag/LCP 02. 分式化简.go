package mathTag


func Fraction(cont []int) []int {
	//a为分子，b为分母
	//但是最后一次转换不用交换分子分母，所以最后a为分母，b为分子
	a, b := 0, 1
	for i := len(cont) - 1; i >= 0; i-- {
		a += cont[i]*b
		a, b = b, a
	}
	g := gcd(a, b)
	return []int{b/g, a/g}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}