"""
实现单例模式的五种方式

1、单例模式
单例模式：保证一个类仅有一个实例，并提供一个访问它的全局访问点。

一个类的唯一实例

适用性：
当类只能有一个实例而且客户可以从一个众所周知的访问点访问它时。
当这个唯一实例应该是通过子类化可扩展的，并且客户应该无需更改代码就能使用一个扩展的实例时。



在方法二中，可以使用__new__来写一个单例模式，或者__call__编写一个装饰器，用装饰器写一个单例
"""
# 方法一，装饰器
def Singleton1(cls):
    _instance = {}

    def _singleton(*args, **kwargs):
        if cls not in _instance:
            _instance[cls] = cls(*args, **kwargs)
        return _instance[cls]

    return _singleton

@Singleton1
class A(object):
    a = 1

    def __init__(self, x=0):
        self.x = x


a1 = A(2)
a2 = A(3)   #实例已经存在，不会再实例化一个对象
print(a1.x, a2.x)


# 方法二，使用类（类方法）
import time
import threading


class Singleton:
    _instance_lock = threading.Lock()

    def __init__(self, *args, **kwargs):
        time.sleep(1)

    # 类方法必须以一个cls作为参数
    @classmethod
    def instance(cls, *args, **kwargs):
        # 第一个判断是为了减少后面多次创建单例时重复加锁
        if not hasattr(cls, "_instance"):
            # 在第二个判断前加锁，是为了支持多线程
            with cls._instance_lock:
                if not hasattr(cls, "_instance"):
                    cls._instance = cls(*args, **kwargs)
        return cls._instance


def task(arg):
    obj = Singleton.instance()
    print(obj)

for i in range(10):
    t = threading.Thread(target=task, args=[i,])
    t.start()


# 方法三 使用__new__
import time
import threading


class Singleton(object):
    _instance_lock = threading.Lock()

    def __init__(self, *args, **kwargs):
        time.sleep(1)
        pass

    def __new__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            with cls._instance_lock:
                if not hasattr(cls, "_instance"):
                    cls._instance = object.__new__(cls)
        return cls._instance


def task(arg):
    obj = Singleton()
    print(obj)

for i in range(10):
    t = threading.Thread(target=task, args=[i, ])
    t.start()


# 方法四 使用metaclass(__call__方法, 继承type)
import time
import threading


class SingleType(type):
    _instance_lock = threading.Lock()

    # 实例化一个对象时，先调用类的__call__方法，然后在其中调用__new__方法和__init__来创建和初始化实例
    def __call__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            with cls._instance_lock:
                if not hasattr(cls, "_instance"):
                    cls._instance = super(SingleType, cls).__call__(*args, **kwargs)
        return cls._instance


class Foo(metaclass=SingleType):
    def __init__(self, name):
        self.name = name


obj1 = Foo("name1")
obj2 = Foo("name2")
print(obj1, obj2)