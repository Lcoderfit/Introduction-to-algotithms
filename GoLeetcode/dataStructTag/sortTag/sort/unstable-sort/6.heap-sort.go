/*
https://blog.csdn.net/chandelierds/article/details/91357784
堆排序：不稳定排序，假设arr为5 5; 第一个5会先与最后一个元素替换，即确定了最大值，然后此时两个5的顺序就变了

算法时间复杂度
每一次堆化为O(logn) (从第一层一直堆化到第k层，总共有n个元素，即2^k<n<2^(k+1); k=logn ), 堆化n-1次，则为O(nlogn)
1、最优：O(nlogn)
2、平均：O(nlogn)
3、最差：O(nlogn)
空间复杂度：O()

*/

package unstable_sort

// 堆排序
func HeapSort(arr []int) []int {
	// 从带有孩子节点且索引最大的元素开始
	// 从索引为n/2的节点作为根节点开始堆化，然后一直堆化到以0为根节点的索引；
	// 即构成了一颗 每颗子树 的根节点都比该子树中其他节点的值都大的 二叉树(暂且叫大根树)
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i)
	}
	// 每次构建大根树，可以确保根节点(arr[0])是所有元素中最大的，将最大元素与最后面的一个元素交换，即确定了最大元素
	// 然后每次将除最后一个元素外的其他元素再次进行堆化（由于根节点被替换了，而根几点的两个孩子都是大根树，所以在堆化时根节点与哪颗子树进行了替换，就对哪颗子树再递归的进行堆化操作）
	end := len(arr) - 1
	for end > 0 {
		arr[0], arr[end] = arr[end], arr[0]
		heapify(arr[:end], 0)
		end--
	}
	return arr
}

// 堆化
// 把整个数组看成一颗二叉树，然后将pos作为根节点，首先令temp指向pos节点两个孩子中的最大值，然后将temp与pos节点进行比较
// 如果temp > pos, 则pos与temp交换值，并且递归对temp节点进行堆化
func heapify(arr []int, pos int) {
	end := len(arr) - 1
	// 节点从0开始，数组从0~n-1可以看成是一颗二叉树的层序遍历
	left := 2*pos + 1
	if left > end {
		return
	}

	// temp取两个孩子中的较大的一个
	right := left + 1
	temp := left
	if right <= end && arr[temp] < arr[right] {
		temp = right
	}

	// pos根节点取最大值
	if arr[temp] < arr[pos] {
		return
	}
	arr[temp], arr[pos] = arr[pos], arr[temp]

	// 对temp进行递归
	heapify(arr, temp)
}
