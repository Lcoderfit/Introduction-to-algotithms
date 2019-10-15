# 文章投票应用
import redis
import time


client = redis.Redis(host='localhost', port=6379, password=124541)


class ArticleVote:
    def __init__(self):
        self.ONE_WEEK_IN_SECONDS = 7 * 86400
        self.VOTE_SCORE = 432

    def article_vote(self, client, user, article):
        """文章投票"""
        # 截止时间
        coutoff = time.time() - self.ONE_WEEK_IN_SECONDS
        if client.zscore('time:', article) < cutoff:
            return

        article_id = article.partition(':')[-1]
        if client.sadd('voted:', article_id, user):
            client.zincrby('score:', article, self.VOTE_SCORE)
            client.hincrby(article, 'votes', 1)

    def post_article(self, client, user, title, link):
        """发布文章"""
        article_id = str(client.incr('article:'))

        voted = 'voted' + article_id
        client.sadd(voted, user)
        client.expire(voted, self.ONE_WEEK_IN_SECONDS)
        post_time = time.time()

        article = 'article:' + article_id
        client.hmset(article, {
            'title': title,
            'link': link,
            'time': post_time,
            'votes': 1
        })

        client.zadd('score:', article, post_time + self.VOTE_SCORE)
        client.zadd('time:', article, post_time)

        return article_id

