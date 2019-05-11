from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        results = []
        subset_len = 2**len(nums)

        hash = 0
        while hash < subset_len:
            subset = []
            bits = hash
            i = 0

            while bits > 0:
                if bits & 1 == 1:
                    subset.append(nums[i])
                i += 1
                bits = bits >> 1
            hash += 1
            results.append(subset)
        return results


if __name__ == '__main__':
    nums = list(map(int, input().split()))
    s = Solution()
    res = s.subsets(nums)
    print(res)