"""

方法1： 贪心
时间复杂度：O(n)
空间复杂度：O(1)

case1:
-2 1 -3 4 -1 2 1 -5 4
r:
6
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def max_sub_array(nums: List[int]) -> int:
        ans = nums[0]
        count = 0
        for i in range(0, len(nums)):
            count += nums[i]
            ans = max(ans, count)
            if count < 0:
                count = 0
        return ans

    @staticmethod
    def max_sub_array1(nums: List[int]) -> int:
        ans = nums[0]
        for i in range(1, len(nums)):
            # 假设f(i)为以nums[i]结尾的连续子数组的最大和，即f(i) = max(f(i) + nums[i], nums[i])
            # 再通过ans获取每一次迭代的最大和的最大值即可
            if nums[i] + nums[i - 1] > nums[i]:
                nums[i] += nums[i - 1]
            if nums[i] > ans:
                ans = nums[i]
        return ans


if __name__ == '__main__':
    # s = Solution()
    # for line in sys.stdin:
    #     # res = s.max_sub_array([int(i) for i in line.strip().split(" ")])
    #     res = s.max_sub_array1([int(i) for i in line.strip().split(" ")])
    #     print(res)
    print("|".join([u"股份有限公司", u"外商投资投资性公司", u"有限责任公司分公司", u"有限责任公司", u"一人有限公司", u"其他有限责任公司", u"外商投资性公司", u"其他有限公司", u"四达服务公司"]))
