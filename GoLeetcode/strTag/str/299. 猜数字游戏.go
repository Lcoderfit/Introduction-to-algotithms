package str

import "strconv"

func getHint(secret string, guess string) string {
	//两个slice分别存储相同位置的不同元素
	sMap, gMap := make([]int, 10), make([]int, 10)
	//a, b分别存储公牛和奶牛的数量
	a, b := 0, 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			//byte可以作为下标，但是不能作为map[int]int的键
			sMap[secret[i]-'0']++
			gMap[guess[i]-'0']++
		}
	}
	//字符串由0-9的数字字符组成
	for i := 0; i < 10; i++ {
		b += min(sMap[i], gMap[i])
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

