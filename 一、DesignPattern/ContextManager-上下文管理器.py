"""
自定义文件管理器的两种方法
"""
# 方法一
class File(object):
    def __init__(self, file_name, file_model):
        self.file_name = file_name
        self.file_mode = file_model

    def __enter__(self):
        self.f = open(self.file_name, self.model)
        return self.f

    def __exit__(self, exc_type, exc_val, exc_tb):
        self.f.close()


# 方法二
from contextlib import contextmanager


@contextmanager
def file(file_name, file_model):
    f = open(file_name, file_model)
    yield f
    f.close()