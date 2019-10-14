import redis


client = redis.Redis(host ='localhost', port=6379, password=124541)

dic = dict()
for i in range(10):
    dic[i] = i
client.hmset('a', dic)

print(client.hgetall('a'))
client.delete('a')