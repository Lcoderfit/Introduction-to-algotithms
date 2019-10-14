#!/bin/python
# -*- coding: utf8 -*-
import sys
import os
import re


# 请完成下面这个函数，实现题目要求的功能
# 当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^
# ******************************开始写代码******************************


def DynamicPathPlanning(matrixGrid):
    if not matrixGrid:
        return 0
    rows = len(matrixGrid)
    cols = len(matrixGrid[0])
    if rows <= 0 or cols <= 0:
        return 0
    dp = [[0]*cols for _ in range(rows)]
    for i in range(rows):
        for j in range(cols):
            if matrixGrid[i][j] == 1:
                dp[i][j] = -1

    for i in range(rows):
        if dp[i][0] == -1:
            break
        dp[i][0] = 1
    for j in range(cols):
        if dp[0][j] == -1:
            break
        dp[0][j] = 1

    for i in range(1, rows):
        for j in range(1, cols):
            if dp[i][j] == -1:
                dp[i][j] = 0
                continue
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
    return dp[rows - 1][cols - 1]


# ******************************结束写代码******************************


_matrixGrid_rows = 0
_matrixGrid_cols = 0
_matrixGrid_rows = int(input())
_matrixGrid_cols = int(input())

_matrixGrid = []
for _matrixGrid_i in range(_matrixGrid_rows):
    _matrixGrid_temp = list(map(int, re.split(r'\s+', input().strip())))
    _matrixGrid.append(_matrixGrid_temp)

res = DynamicPathPlanning(_matrixGrid)

print(str(res) + "\n")