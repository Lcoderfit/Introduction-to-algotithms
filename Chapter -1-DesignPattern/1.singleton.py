'''一、单例模式
实现方法：4种
1、用装饰器实现
'''

def singleton(cls):
    instances = {}
    def wrapper(*args, **kwargs):
        if cls not in instances:
            instances[cls] = cls(*args, **kwargs)
        return instances[cls]
    return wrapper

@singleton
class Foo(object):
    pass


if __name__ == '__main__':
    foo1 = Foo()
    foo2 = Foo()
    print(id(foo1), id(foo2))