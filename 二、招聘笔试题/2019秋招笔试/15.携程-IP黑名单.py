#!/bin/python
# -*- coding: utf8 -*-
import sys
import os
import re


# 请完成下面这个函数，实现题目要求的功能
# 当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^
# ******************************开始写代码******************************


def CheckBlackList(userIP, blackIP):
    blackip_list = blackIP.split("/")
    black_parentIP = blackip_list[0]
    if len(blackip_list) < 2:
        if userIP == black_parentIP:
            return 1
        return 0
    number = int(blackip_list[1])
    pos = number // 8
    y = number % 8
    p = int(black_parentIP[pos])
    common_list = black_parentIP[:p]
    common = "".join(common_list)

    x = 8 - y
    t = 1
    for i in range(x):
        if p & t:
            p = p - t
        t = t << 1
    if pos < 3:
        min_p = str(p)
        max_p = str(p + 2^x - 1)
        m_w = '0' * (3 - pos - 1) + '1'
        a_w = '255' * (3 - pos - 1) + '254'
        min_ret = common + min_p + m_w
        max_ret = common + max_p + a_w
    else:
        min_p = str(p + 1)
        max_p = str(p + 2^x - 1)
        min_ret = common + min_p
        max_ret = common + max_p

    user_ip = int("".join(userIP.split(".")))
    min_ret = int(min_ret)
    max_ret = int(max_ret)

    if user_ip >= min_ret and user_ip <= max_ret:
        return 1
    return 0


# ******************************结束写代码******************************


try:
    _userIP = input()
except:
    _userIP = None

try:
    _blackIP = input()
except:
    _blackIP = None

res = CheckBlackList(_userIP, _blackIP)

print(str(int(res)) + "\n")