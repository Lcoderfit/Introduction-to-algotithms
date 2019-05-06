'''一、单例模式
实现方法：4种
1、用装饰器实现
2、用__new__实现
'''
def singleton1(cls):
    instances = {}
    def wrapper(*args, **kwargs):
        if cls not in instances:
            instances[cls] = cls(*args, **kwargs)
        return instances[cls]
    return wrapper

@singleton1
class Foo1(object):
    pass


class Singleton2(object):
    '''new方法返回当前类的实例'''
    def __new__(cls, *args, **kwargs):
        if not hasattr(cls, '_instance'):
            cls._instance = super(Singleton2, cls).__new__(cls, *args, **kwargs)
        return cls._instance


class Foo2(Singleton2):
    pass


class Singleton3(type):
    '''用元类创建'''
    def __call__(cls, *args, **kwargs):
        if not hasattr(cls, '_instance'):
            cls._instance = super(Singleton3, cls).__call__(*args, **kwargs)
        return cls._instance


class Foo3(metaclass=Singleton3):
    pass


if __name__ == '__main__':
    #用装饰器
    foo1 = Foo1()
    foo2 = Foo1()
    print(id(foo1), id(foo2))

    #用__new__()
    foo3 = Foo2()
    foo4 = Foo2()
    print(id(foo3), id(foo4))

    #用元类
    foo5 = Foo3()
    foo6 = Foo3()
    print(id(foo5), id(foo6))
