# import numpy as np
#
# a = np.array([1, 2, 3, 4])
# print(a)
#
# import time
#
# a = np.random.rand(1000000)
# b = np.random.rand(1000000)
#
# tic = time.time()
# c = np.dot(a, b)
# toc = time.time()
#
# print(c)
# print("Vectorized version:" + str(1000*(toc-tic)) + "ms")
#
# tic = time.time()
# for i in range(1000000):
#     c += a[i]*b[i]
#
# toc = time.time()
# print(c)
# print("for loop version:" + str(1000*(toc-tic)) + "ms")


# 第二个示例
import numpy as np

n = 10
u = np.zeros((n, 2))
print(type(u[0][0]))
