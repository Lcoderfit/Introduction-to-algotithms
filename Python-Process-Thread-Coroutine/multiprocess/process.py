from multiprocessing import Process
import os


def run_proc(name):
    print('Run child process %s(%s)' % (name, os.getpid()))


if __name__ == '__main__':
    print('Parent process %s.' % os.getpid())
    # 两个参数，target表示子进程要执行的代码，args是一个元组，表示传入target函数的参数
    p = Process(target=run_proc, args=('test',))
    # 启动进程
    p.start()
    # 用于进程同步，等待子进程结束
    p.join()
    print('Child process end.')
