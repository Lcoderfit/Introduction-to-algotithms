class Solution:
    def LogestArray(self, array):
        if not array:
            return list()
        if len(array) == 1:
            return array
        ret = list()
        length = len(array) - 1
        max = 1
        i = 0
        while i < length:
            tmp = list()
            while i < length and array[i] < array[i + 1]:
                tmp.append(array[i])
                i += 1
            tmp.append(array[i])
            if max == len(tmp):
                ret.append(tmp)
            elif max < len(tmp):
                max = len(tmp)
                ret.clear()
                ret.append(tmp)
            i += 1

        return ret


if __name__ == "__main__":
    array = [int(i) for i in input().split(" ")]
    s = Solution()
    ret = s.LogestArray(array)
    for i in range(len(ret)):
        print(" ".join([str(j) for j in ret[i]]))

