"""
https://blog.csdn.net/chandelierds/article/details/91357784
堆排序：不稳定排序，如果两个相同元素，a在左边，b在右边(则a在b前面)，右边的先到堆顶，
然后排序之后被交换到堆底，a之后也交换到堆底，则b在a的前面
算法时间复杂度
每一次堆化为O(logn), 堆化n-1次，则为O(nlogn)
1、最优：O(nlogn)
2、平均：O(nlogn)
3、最差：O(nlogn)
空间复杂度：O(1)
"""
class Solution:
    def heapify(self, arr, start, end):
        """
        :param start: 一开始进行堆化的父亲节点
        :param end: 数组下标最大值
        :return:
        """
        root = start
        child = root * 2 + 1
        while child <= end:
            if child + 1 <= end and arr[child] < arr[child + 1]:
                child += 1
            if arr[root] < arr[child]:
                arr[root], arr[child] = arr[child], arr[root]
                root = child
                child = root * 2 + 1
            else:
                break

    def heap_sort(self, arr):
        first = len(arr) // 2 - 1
        for start in range(first, -1, -1):
            self.heapify(arr, start, len(arr) - 1)
        for end in range(len(arr) - 1, 0, -1):
            arr[0], arr[end] = arr[end], arr[0]
            self.heapify(arr, 0, end - 1)
        return arr

    def random_array(self, n):
        ret = list()
        multi = set()
        for i in range(n):
            k = random.randint(1, n)
            while True:
                if k in multi:
                    k = random.randint(1, n)
                else:
                    multi.add(k)
                    ret.append(k)
                    break
        return ret



if __name__ == "__main__":
    s = Solution()
    arr = s.random_array(10)
    print(arr)
    ret = s.heap_sort(arr)
    print(ret)
