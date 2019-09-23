"""
1.100亿个URL中找出重复的URL

1.一行一行读取URL，对每一个URL进行hash然后取模，x = hash(u) % 400
将文件以x进行命名，经过hash取模之后的URL存入对应的文件中
2.读取每个文件，然后用dict进行映射处理，统计重复的URL及其数量
3.归并每个文件的统计结果
"""
# 解法一：Python
import os


class Solution:
    def find_multi_url(self, path):
        with open(path, "rb") as f:
            line = f.readline()
            x = hash(line) % 4000
            wf = open(path + x, "wb")
            wf.write(line)
            wf.close()

        ret = dict()
        for i in range(4000):
            dic = dict()
            if os.path.isfile(path + i):
                f = open(path + i, "rb")
                for line in f:
                    if not dic.get(line):
                        dic[line] = 1
                    else:
                        dic[line] += 1
            ret.update(dic)

        return ret