package rejectionSimpling

import "math/rand"

func Rand10() int {
	num := rand7() + (rand7()-1)*7
	for num > 40 {
		num = rand7() + (rand7()-1)*7
	}
	return 1 + (num-1)%10
}

func rand7() int {
	return rand.Intn(7) + 1
}
