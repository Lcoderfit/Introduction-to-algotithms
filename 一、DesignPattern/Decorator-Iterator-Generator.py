# # from typing import List
# # from functools import cmp_to_key
# #
# #
# # class Solution(object):
# #     def min_number(self, nums: List) -> str:
# #         if any(nums):
# #             lmb = lambda x, y: int(x + y) - int(y + x)
# #             nums = sorted([str(x) for x in nums], key=cmp_to_key(lmb))
# #             return ''.join(nums)
# #         return '0'
# #
# #
# # if __name__ == "__main__":
# #     numbers = input().split(",")
# #     s = Solution()
# #     ret = s.min_number(numbers)
# #     print(ret)
#
# # from functools import wraps
# #
# #
# # def validate_serializer(serializer):
# #     def validate(*args, **kwargs):
# #         @wraps(serializer)
# #         def handler(*args, **kwargs)
#
#
# #
# # import functools
#
#
# import functools
#
#
# def outter(param):
#     def wrapper(f):
#         @functools.wraps(f)
#         def handler(*args, **kwargs):
#             param(args[0])
#             print("Hello")
#             return f(*args, **kwargs)
#         return handler
#     return wrapper
#
#
# def p(name):
#     print(name)
#
#
# @outter(p)
# def k(name):
#     print("world")
#
#
# if __name__ == "__main__":
#     k("luhu")
#
#
#
# # def user_login_data(params):
# #     def wrapper(real_func):
# #         @functools.wraps(real_func)
# #         def handler(*args, **kwargs):
# #             print("FUCK")
# #             params()
# #             return real_func(*args, **kwargs)
# #         return handler
# #     return wrapper
# #
# # @user_login_data    #没有传递参数，num1 = user_login_data(num1) -->wrapper
# # def num1():
# #     print("aaa")
# #
# # def k():
# #     print("KKKKK")
# #
# # @user_login_data(k) #传参，@user_login_data(k) --> @wrapper
# # def num2(*args, **kwargs):         # num2 = wrapper(num2) --> handler
# #     print("NUM2")
#
#
#
# # def validate_serializer(serializer):
# #     """
# #     对传入请求的数据进行反序列化和数据有效性验证
# #     :param serializer: 类序列化器
# #     :return:
# #     """
# #     @functools.wraps(serializer)
# #     def validate(view_method):
# #         def handle(*args, **kwargs):
# #             self = args[0]
# #             request = args[1]
# #             s = serializer(data=request.data)
# #             if s.is_valid():
# #                 request.data = s.data
# #                 request.serializer = s
# #                 return view_method(*args, **kwargs)
# #             else:
# #                 return self.invalid_seria(s)
# #
# #         return handle
# #
# #     return validate
#
#
# # if __name__ == "__main__":
# #     print(num1.__name__)
# #     print(num2.__name__)
# #     print("*************")
# #     a, b = 1, 2
# #     num2(a, b=3)
#
# # import functools
# #
# #
# # def validate_serializer(serializer):
# #     @functools.wraps(serializer)
# #     def validate(view_method):
# #         @functools.wraps(view_method)
# #         def handler(*args, **kwargs):
# #             self = args[0]
# #             request = args[1]
# #             s = serializer(data=request.data)
# #             if s.is_valid():
# #                 request.data = s.data
# #                 request.serializer = s
# #                 return view_method(*args, **kwargs)
# #             else:
# #                 return self.invalid_seria(s)
# #
# #         return handler
# #
# #     return validate
# #
# # def outter(param):
# #     def wrapper(f):
# #         # @functools.wraps(f)
# #         def inner(*args, **kwargs):
# #             param()
# #             print("Hello world")
# #             return f(*args, **kwargs)
# #         return inner
# #     return wrapper
# #
# #
# # def k():
# #     print("KKKK")
# #
# # # @outter # func1 = outter(func1)     func1 --> wrapper
# # @outter(k)  #@outter(k) --> @wrapper    func1 = @wrapper(func1)  func1--> inner
# # def func1(*args, **kwargs):
# #     print("FUNC1")
# # import functools
# #
# # # class A:
# # #     def __init__(self, n):
# # #         self.n = n
# # #     def __call__(self, *args, **kwargs):
# # #         print(self.n)
# # #
# # # class Wrapper:
# # #     def __init__(self, param):
# # #         self.parma = param
# # #     def __call__(self, func):
# # #         print("HELLO WORLD")
# # #         self.parma()
# # #         return 10
# # #
# # # def k():
# # #     print("KKKKKKK")
# # #
# # # @Wrapper(k)     #func = Wrapper(k)(func)
# # # def func(*args, **kwargs):
# # #     print("FUNC")
# # #     print(args, kwargs)
# # #
# # #
# # #
# # # if __name__ == "__main__":
# # #     print("L1")
# # #     print(func)
# # #     print("L2")
# # #     a, b = 1, 2
# # #     func(a, b=3)
# # #     # print(k)
# #
# # class Wrapper:
# #     def __init__(self, param):
# #         self.param = param
# #     def __call__(self, view_method):
# #         # @functools.wraps(self.param)
# #         @functools.wraps(view_method)
# #         def handler(*args, **kwargs):
# #             self.param(*args, **kwargs)
# #             return view_method(*args, **kwargs)
# #         return handler
# #
# #
# # @Wrapper    # func = Wrapper(func)  func()  --> __call__()的返回值 --> handler
# # def func(*args, **kwargs):
# #     print("FUNC")
# #
# #
# # def k():
# #     print("KKKKK")
# #
# # @Wrapper(k)     #func2 = Wrapper(k)(func2) func2 --> __call__(func2) --> handler
# # def func2(*args, **kwargs):
# #     print("FUNC2")
# #     return 10
# #
# # # func2 = Wrapper(k)(func2) --> __call__(func2) --> handler
# #
# # if __name__ == "__main__":
# #     print(func2.__name__)
# #     k = func2()
# #     print(k)
# # import functools
# #
# #
# # def outter(param):
# #     @functools.wraps(param)
# #     def wrapper(f):
# #         @functools.wraps(f)
# #         def inner(*args, **kwargs):
# #             param()
# #             print("HELLO WORLD")
# #             return f(*args, **kwargs)
# #         return inner
# #     return wrapper
# #
# #
# def k():
#     print("KKKKKKK")
#
#
# import functools
#
#
# import functools
#
# def pa(name):
#     print(name)
#
#
# class Wrapper:
#     def __init__(self, param):
#         self.param = param
#
#     def __call__(self, f):
#         @functools.wraps(f)
#         def handler(*args, **kwargs):
#             self.param(args[0])
#             print("HELLO")
#             return f(*args, **kwargs)
#         return handler
#
# @Wrapper(pa)    #f = @Wrapper(pa)(f) ==>    f = __call__(f) ==> f = handler
# def f(name):
#     print(name)
#
#
# f("luhu")
#
#
#
#
#
#
#
#
#
#
# # import functools
# #
# #
# # class Wrapper:
# #     def __init__(self, param):
# #         self.param = param
# #     def __call__(self, f):
# #         @functools.wraps(f)
# #         def inner(*args, **kwargs):
# #             self.param()
# #             print("HELLO WORLD")
# #             return f(*args, **kwargs)
# #         return inner
#
#
# @Wrapper(k)     # func2 = Wrapper(k)(func2)     func2 --> __call__(func2)   func2 --> inner
# def func2(*args, **kwargs):
#     print("FUNC2")
#
# @Wrapper    # func3 = Wrapper(fun3)     func3() 调用 __call__()方法
# def func3(*args, **kwagrs):
#     print("FUNC3")
#
#
# if __name__ == "__main__":
#     # print(func2.__name__)
#     #
#     # a, b = 1, 2
#     # func2(a, b)
#
#     print("***********")
#     print(id(func3), id(Wrapper))
#
#

