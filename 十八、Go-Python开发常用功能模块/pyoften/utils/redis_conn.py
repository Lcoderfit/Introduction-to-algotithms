import redis


class RedisConn:
    def single_conn(self):
        conn = redis.Redis(host="localhost", port=6379, password=124541, db=0)
        conn.set("a", 10)
        a = conn.get("a")
        print(a)

    def multi_conn(self):
        pool = redis.ConnectionPool(host="localhost", port=6379, password=124541, db=0, max_connections=100)
        conn2 = redis.Redis(connection_pool=pool)
        conn2.set("b", 5)
        b = conn2.get("b")
        print(b)


if __name__ == "__main__":
    redisConn = RedisConn()
    redisConn.single_conn()
    redisConn.multi_conn()