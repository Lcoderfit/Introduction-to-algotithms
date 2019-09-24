"""
3.数组中的重复的数字
时间复杂度：O(n)
空间复杂度：O(1)
"""
class Solution:
    def duplicate(self, numbers, duplication):
        for i in range(len(numbers)):
            while numbers[i] != i:
                if numbers[i] == numbers[numbers[i]]:
                    duplication[0] = numbers[i]
                    return True
                t = numbers[i]
                numbers[i], numbers[t] = numbers[t], numbers[i]
        return False


if __name__ == "__main__":
    numbers = [2, 3, 1, 0, 2, 5, 3]
    duplication = [0]
    s = Solution()
    ret = s.duplicate(numbers, duplication)
    print(ret, duplication[0])