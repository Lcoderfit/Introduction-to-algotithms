"""
实现单例模式的五种方式
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


# 方法二，使用类
class Singleton2:
    def __init__(self, *args, **kwargs):
        pass

    @classmethod
    def instance(cls, *args, **kwargs):
        if not hasattr(Singleton2, "_instance"):
            Singleton2._instance = Singleton2(*args, **kwargs)
        return Singleton2._instance


# 方法三 使用__new__
import threading


class Singleton3:
    _instance_lock = threading.Lock()
    k = 1

    def __init__(self):
        pass

    def __new__(cls, *args, **kwargs):
        if not hasattr(Singleton3, "_instance"):
            with Singleton3._instance_lock:
                Singleton3._instance = super().__new__(cls)
        return Singleton3._instance


print("******方法三******")
obj = Singleton3()
print(obj.k)
obj2 = Singleton3()
obj2.k = 4
print(id(obj), id(obj2))


# 方法四 使用metaclass
import threading


class SingletonType(type):
    _instance_lock = threading.Lock()
    def __call__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            with SingletonType._instance_lock:
                cls._instance = super(SingletonType, cls).__call__(*args, **kwargs)
        return cls._instance


class Foo(metaclass=SingletonType):
    def __init__(self, name):
        self.name = name


obj1 = Foo("name1")
obj2 = Foo("name2")
print("******方法四******")
print(id(obj1), id(obj2))
