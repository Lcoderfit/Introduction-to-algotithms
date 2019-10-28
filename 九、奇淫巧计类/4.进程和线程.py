# from multiprocessing import Process
# import os
#
# # 子进程要执行的代码
# def run_proc(name):
#     print('Run child process %s (%s)...' % (name, os.getpid()))
#
# if __name__=='__main__':
#     print('Parent process %s.' % os.getpid())
#     p = Process(target=run_proc, args=('test',))
#     print('Child process will start.')
#     p.start()
#     p.join()
#     print('Child process end.')


# *****************************
# from multiprocessing import Pool
# import os, time, random
#
# def long_time_task(name):
#     print('Run task %s (%s)' % (name, os.getpid()))
#     start = time.time()
#     time.sleep(random.random()*3)
#     end = time.time()
#     print('Task %s run %0.2f seconds.' % (name, (end - start)))
#
#
# if __name__ == "__main__":
#     print('Parent process %s.' % os.getpid())
#     p = Pool(8)
#     for i in range(5):
#         p.apply_async(long_time_task, args=(i, ))
#     print('Waiting for all subprocesses done')
#     p.close()
#     p.join()
#     print('ALL subprocesses done')


# **********************************
# import subprocess
#
# print('$ nslookup www.python.org')
# r = subprocess.call(['nslookup', 'www.python.org'])
# print('Exit code:', r)


# *****************
# import subprocess
#
# print('$ nslookup')
# p = subprocess.Popen(['nslookup'], stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
# output, err = p.communicate(b'set q=mx\npython.org\nexit\n')
# print(output.decode('utf-8'))
# print('Exit code:', p.returncode)



# ***********
# from multiprocessing import Process, Queue
# import os, time, random
#
#
# def write(q):
#     print("Process to write: %s" % os.getpid())
#     for value in ['A', 'B', 'C']:
#         print('Put %s to queue...' % value)
#         q.put(value)
#         time.sleep(random.random())
#
# def read(q):
#     print('Process to read: %s' % os.getpid())
#     while True:
#         value = q.get(True)
#         print('Get %s from queue' % value)
#
#
# if __name__ == "__main__":
#     q = Queue()
#     pw = Process(target=write, args=(q,))
#     pr = Process(target=read, args=(q,))
#
#     pw.start()
#     pr.start()
#
#     pw.join()
#     pr.terminate()


# *******************************
# import time, threading
#
#
# def loop():
#     print('thread %s is running...' % threading.current_thread().name)
#     n = 0
#     while n < 5:
#         n = n + 1
#         print('thread %s >>> %s' % (threading.current_thread().name, n))
#         time.sleep(1)
#     print('thread %s ended.' % threading.current_thread().name)
#
# print('thread %s is running...' % threading.current_thread().name)
# t = threading.Thread(target=loop, name="LoopThread")
# t.start()
# t.join()
# print('thread %s ended.' % threading.current_thread().name)


# ****************************
# import time, threading
#
#
# blance = 0
#
#
# def change_it(n):
#     global blance
#     blance = blance + n
#     blance = blance - n
#
#
#
# def run_thread(n):
#     for i in range(100000):
#         change_it(n)
#
# t1 = threading.Thread(target=run_thread, args=(5,))
# t2 = threading.Thread(target=run_thread, args=(8,))
# t1.start()
# t2.start()
# t1.join()
# t2.join()
#
# print(blance)
# # balance = 0
# # lock = threading.Lock()
# #
# # def run_thread(n):
# #     for i in range(100000):
# #         # 先要获取锁:
# #         lock.acquire()
# #         try:
# #             # 放心地改吧:
# #             change_it(n)
# #         finally:
# #             # 改完了一定要释放锁:
# #             lock.release()

# ********************8
import threading


local_school = threading.local()


def process_student():
    std = local_school.student
    print("Hello, %s (in %s)" % (std, threading.current_thread().name))

def process_thread(name):
    local_school.student = name
    process_student()


t1 = threading.Thread(target=process_thread, args=('Alice',), name='Thread-A')
t2 = threading.Thread(target=process_thread, args=('Bob',), name='Thread-B')
t1.start()
t2.start()
t1.join()
t2.join()