# 模拟Reis分布式锁支持可重入性的代码
import redis
import threading


locks = threading.local()
locks.redis = dict()


class RedisDCSLocks:
    def key_for(self, user_id):
        return "account_{}".format(user_id)

    def _lock(self, client, key):
        return bool(client.set(key, 1, nx=True, ex=5))

    def _unlock(self, client, key):
        client.delete(key)

    def lock(self, client, user_id):
        key = self.key_for(user_id)
        if key in locks.redis:
            locks.redis[key] += 1
            return True
        ok = self._lock(client, key)
        if not ok:
            return False
        locks.redis[key] = 1
        return True

    def unlock(self, client, user_id):
        key = self.key_for(user_id)
        if key in locks.redis:
            locks.redis[key] -= 1
            if locks.redis[key] <= 0:
                del locks.redis[key]
            return True
        return False


if __name__ == "__main__":
    client = redis.StrictRedis(host='localhost', port=6379, password=124541)
    s = RedisDCSLocks()
    print("lock", s.lock(client, "codehole"))
    print("lock", s.lock(client, "codehole"))
    print("unlock", s.unlock(client, "codehole"))
    print("unlock", s.unlock(client, "codehole"))