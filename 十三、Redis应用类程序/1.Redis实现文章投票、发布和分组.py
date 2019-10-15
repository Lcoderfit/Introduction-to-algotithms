# 文章投票应用
import redis
import time


client = redis.Redis(host='localhost', port=6379, password=124541)


class ArticleVote:
    def __init__(self):
        self.ONE_WEEK_IN_SECONDS = 7 * 86400
        self.VOTE_SCORE = 432

    def article_vote(self, client, user, article):
        coutoff = time.time() - self.ONE_WEEK_IN_SECONDS
        if client.zscore('time:', article) < cutoff:
            return

        article_id = article.partition(':')[-1]
        if client.sadd('voted:', article_id, user):
            client.zincrby('score:', article, self.VOTE_SCORE)
            client.hincrby(article, 'votes', 1)