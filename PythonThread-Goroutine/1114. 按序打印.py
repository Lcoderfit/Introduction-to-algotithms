# from threading import Lock
#
# class Foo:
#     def __init__(self):
#         self.first_job_done = Lock()
#         self.second_job_done = Lock()
#         self.first_job_done.acquire()
#         self.second_job_done.acquire()
#
#
#     def first(self, printFirst: 'Callable[[], None]') -> None:
#         # printFirst() outputs "first". Do not change or remove this line.
#         printFirst()
#         self.first_job_done.release()
#
#
#     def second(self, printSecond: 'Callable[[], None]') -> None:
#         # with语句可以自动加锁与解锁
#         with self.first_job_done:
#             # printSecond() outputs "second". Do not change or remove this line.
#             printSecond()
#             # 任务二运行完毕，释放锁
#             self.second_job_done.release()
#
#
#     def third(self, printThird: 'Callable[[], None]') -> None:
#         with self.second_job_done:
#             # printThird() outputs "third". Do not change or remove this line.
#             printThird()

# from threading import Lock
#
# class Foo:
#     def __init__(self):
#         self.first_job_done = Lock()
#         self.second_job_done = Lock()
#         self.first_job_done.acquire()
#         self.second_job_done.acquire()
#
#
#     def first(self, printFirst: 'Callable[[], None]') -> None:
#         # printFirst() outputs "first". Do not change or remove this line.
#         printFirst()
#         self.first_job_done.release()
#
#
#     def second(self, printSecond: 'Callable[[], None]') -> None:
#         if self.first_job_done.acquire():
#             # printSecond() outputs "second". Do not change or remove this line.
#             printSecond()
#             # 任务二运行完毕，释放锁
#             self.second_job_done.release()
#
#
#     def third(self, printThird: 'Callable[[], None]') -> None:
#         if self.second_job_done.acquire():
#             # printThird() outputs "third". Do not change or remove this line.
#             printThird()

from threading import Lock


if "__name__" == "__main__":
    a = Lock()
    b = Lock()
    a.acquire()
    b.acquire()

    if a.acquire():
        print("alsdkjf")
        k.release()
