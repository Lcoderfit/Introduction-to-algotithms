class Solution:
    def FindPos(self, array, number):
        if not array:
            return (-1, -1)
        left, right = 0, len(array) - 1
        while left <= right:
            mid = (left + right) // 2
            if array[mid] == number:
                break
            elif array[mid] > number:
                right = mid - 1
            else:
                left = mid + 1
        ret = (-1, -1)
        i, j = mid - 1, mid + 1
        if array[mid] == number:
            while i >= 0  and array[i] == number:
                i -= 1
            while (j < len(array)) and (array[j] == number):
                j += 1
            ret = (i + 1, j - 1)
        return ret


if __name__ == "__main__":
    n = int(input())
    array = [int(i) for i in input().split(" ")]
    number = int(input())
    s = Solution()
    ret = s.FindPos(array, number)
    print(ret[0], ret[1])
