"""
时间限制：C/C++语言 1000MS；其他语言 3000MS
内存限制：C/C++语言 65536KB；其他语言 589824KB
题目描述：
设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。

写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

输入
第一行输入缓存容量，包含一个整数N，1≤N≤10。

接下来N行，每一行是一个put或者get的操作。若输入一个数字代表get指定的key，若输入两个数字，则前者为key，后者为value，进行put操作。

输出
缓存里最终键值对，按访问时间升序排序，逐行输出
"""
import sys

from collections import OrderedDict


class LRUCache:
    def __init__(self, size=5):
        self.cache = OrderedDict()
        self.size = size

    def get(self, key):
        if key in self.cache:
            value = self.cache[key]
            del self.cache[key]
            self.cache[key] = value
            return value
        else:
            return -1

    def put(self, key, value):
        if size == len(self.cache):
            self.cache.popitem(last=False)
            self.cache[key] = value
        elif key in self.cache:
            del self.cache[key]
            self.cache[key] = value
        else:
            self.cache[key] = value

    def lru_print(self):
        for k, v in self.cache.items():
            print(k, v)


if __name__ == "__main__":
    size, times = [int(i) for i in input().split(" ")]
    s = LRUCache(size)
    for i in range(times):
        opr_list = [int(i) for i in input().split(" ")]
        if len(opr_list) == 2:
            s.put(opr_list[0], opr_list[1])
        elif len(opr_list) == 1:
            s.get(opr_list[0])
    s.lru_print()
