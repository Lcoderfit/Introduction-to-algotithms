from multiprocessing import Pool
import os, time, random


def long_time_task(name):
    print('Run task %s(%s)...' % (name, os.getpid()))


if __name__ == '__main__':
    print('Parent process %s.' % os.getpid())
    # 如果Pool不加参数，则默认大小是CPU核数
    # Pool(5)表示最多可以同时跑4个进程
    p = Pool(4)
    for i in range(5):
        # sync是同步，async是异步
        # 异步执行多进程
        p.apply_async(long_time_task, args=(i,))
    print('Waiting for all subprocesses done...')
    # 关闭进程池，不能再添加新的进程
    p.close()
    # 等待所有进程执行完毕
    p.join()
    print('All subprocesses done.')
