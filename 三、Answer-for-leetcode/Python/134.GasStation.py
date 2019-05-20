from typing import List


class Solution1:
    '''该方法超时，不予采用'''
    def can_complete_circuit(self, gas: List[int], cost: List[int]) -> int:
        gas_station_len = len(gas)
        for i in range(gas_station_len):
            gas_surplus = 0
            flag = True
            for j in range(gas_station_len):
                index = (i+j) % gas_station_len
                gas_surplus = gas_surplus + gas[index] - cost[index]
                if gas_surplus < 0:
                    flag = False
                    break
            if flag:
                return i
        return -1


class Solution2:
    '''贪心算法'''
    def can_complete_circuit(self, gas: List[int], cost: List[int]) -> int:
        if not len(gas) or not len(cost) or len(gas) != len(cost):
            return -1

        tank, total, start = 0, 0, 0

        for i in range(len(gas)):
            tank += gas[i] - cost[i]

            '''
            如果从某一个站x开始到y站时候油量不足，那么从x到y站之间的任何一站开始到y站都会油量不足;
            在遍历gas过程中，total存储从x到y的油量不足的量
            '''
            if tank < 0:
                start = i + 1
                total += tank
                tank = 0

        if total + tank < 0:
            return -1
        else:
            return start


if __name__ == '__main__':
    while True:
        gas = list(map(int, input().split()))
        cost = list(map(int, input().split()))

        s = Solution2()
        res = s.can_complete_circuit(gas, cost)
        print(res)
