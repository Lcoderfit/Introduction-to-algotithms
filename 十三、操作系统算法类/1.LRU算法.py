"""
一.LRU算法(最近最少使用算法)
1、算法应用场景：适合过滤掉最近最少使用的数据，筛选热点数据
2、计算机理论：最近使用的页面数据会在未来一段时期内仍然被使用,已经很久没有使用的页面很有可能在未来较长的一段时间内仍然不会被使用。
时间复杂度：O()
空间复杂度：O()
"""
# 基于普通dict和list实现
class LRUCache:
    def __init__(self, size=5):
        self.size = size
        # 存储已经缓存的数据
        self.cache = dict()
        # 存储LRU列表
        self.key = list()

    def get(self, key):
        # Python3可以用self.cache.__contains__(key)或者key in self.cache 或 key in self.cache.keys()
        # if key in self.cache:
        if self.cache.__contains__(key):
            self.key.remove(key)
            self.key.insert(0, key)
            return self.cache[key]
        else:
            return None

    def set(self, key, value):
        if self.cache.__contains__(key):
            self.cache[key] = value
            self.key.remove(key)
            self.key.insert(0, key)
        elif len(self.cache) == self.size:
            old_key = self.key.pop()
            self.key.insert(0, key)
            self.cache.pop(old_key)
            self.cache[key] = value
        else:
            self.cache[key] = value
            self.key.insert(0, key)


# Python3的版本
from collections import OrderedDict


class LRUCache1:
    """左边是栈底，右边是栈顶，最近最少使用的在左边"""
    def __init__(self, size=5):
        self.size = size
        self.cache = OrderedDict()
        self.key = list()

    def get(self, key):
        if key in self.cache:
            value = self.cache.pop(key)
            self.cache[key] = value
            return value
        else:
            return None

    def set(self, key, value):
        if key in self.cache:
            self.cache.pop(key)
            self.cache[key] = value
        elif len(self.cache) == self.size:
            self.cache.popitem(last=False)
            self.cache[key] = value
        else:
            self.cache[key] = value


if __name__ == "__main__":
    lru_list = [4, 7, 0, 7, 1, 0, 1, 2, 1, 2, 6]
    # lru_cache = LRUCache()
    lru_cache = LRUCache1()
    for i in lru_list:
        lru_cache.set(i, 0)
        print(lru_cache.key, lru_cache.cache)


