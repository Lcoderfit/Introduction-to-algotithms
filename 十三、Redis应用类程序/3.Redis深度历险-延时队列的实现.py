"""
延时队列的实现
"""
import json
import redis
import time
import uuid


class Solution:
    def delay(self, msg):
        msg.id = str(uuid.uuid4())
        value = json.dumps(msg.id)
        retry_ts = time.time() + 5
        # zset: key score1 value1 score2 value2
        # score: 到期处理时间 value: 请求的ID
        redis.zadd("delay-queue", retry_ts, value)

    def loop(self):
        while True:
            # 返回score位于0到time.time()之间的value（默认从小到大排序）, 返回索引位于start和num之间的集合,左闭右开
            values = redis.zrangebyscore("delay-queue", 0, time.time(), start=0, num=1)
            if not values:
                time.sleep(1)
                continue
            value = values[0]
            success = redis.zrem("delay-queue", value)
            if success:
                msg = json.loads(value)
                handle_msg(msg)
