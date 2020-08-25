package str


func minimumSwap(s1 string, s2 string) int {
	//相同位置i处，s1[i]和s2[i]为一组
	//如果有两组xy或者两组yx, 则只需交换一次
	//如果有一组xy和一组yx, 则需要交换两次
	//所以优先交换两组xy或者两种yx
	xy, yx := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		} else if s1[i] == 'x' {
			xy++
		} else {
			yx++
		}
	}
	//要使s1可以通过交换得到s2, 则s1和s2中x的总数及y的总数必须为偶数, 否则无法交换，返回-1
	//设有m组xy，有n组yx；则x的总数 == m+n为偶数
	//两种情况：1.m和n均为偶数  2.m和n均为奇数
	if (xy + yx) % 2 == 1 {
		return -1
	}
	//要使得操作数最小，则应先交换两组xy或两组yx
	//如果m和n均为偶数，则 ret = xy/2 + yx/2
	//如果均为奇数，挑出一组xy和一组yx交换(2次)，使剩下的xy和yx为偶数
	//则ret = (xy-1)/2 + (yx-1)/2 + 2 == (xy+1)/2 + (yx+1)/2
	//由于除法向下取整，所以两种情况都可以写成(xy+1)/2 + (yx+1)/2
	return (xy+1)/2 + (yx+1)/2
}