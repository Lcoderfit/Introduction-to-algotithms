# _*_ coding: utf-8 _*_
import json
import time
import sys

from kafka import KafkaProducer
from kafka import KafkaConsumer
from kafka.errors import KafkaError

KAFKA_HOST = "localhost"
KAFKA_PORT = 9092
KAFKA_TOPIC = "topic-demo"


class KafkaProduceData:
    def __init__(self, kafka_host, kafka_port, kafka_topic):
        self.kafka_host = kafka_host
        self.kafka_port = kafka_port
        self.kafka_topic = kafka_topic
        bootstrap_servers = f"localhost:9092"
        self.producer = KafkaProducer(bootstrap_servers=bootstrap_servers)

    def send_json_data(self, params):
        try:
            params_message = json.dumps(params, ensure_ascii=False)
            producer = self.producer
            print(params_message)
        except KafkaError as e:
            print(e)


class KafkaConsumeData:
    def __init__(self, kafka_host, kafka_port, kafka_topic):
        self.kafka_host = kafka_host
        self.kafka_port = kafka_port
        self.kafka_topic = kafka_topic
        bootstrap_servers = "localhost:9092"
        self.consumer = KafkaConsumer(self.kafka_topic, bootstrap_servers=bootstrap_servers)

    def consume_data(self):
        try:
            for message in self.consumer:
                yield message
        except KeyboardInterrupt as e:
            print(e)


def main(xtype):
    if xtype == "p":
        producer = KafkaProduceData(KAFKA_HOST, KAFKA_PORT, KAFKA_TOPIC)
        print("===========> producer:", producer)
        for _id in range(100):
            params = {"msg": _id}
            producer.send_json_data(params)
            time.sleep(1)
    elif xtype == "c":
        consumer = KafkaConsumeData(KAFKA_HOST, KAFKA_PORT, KAFKA_TOPIC)
        print("===========> consumer:", consumer)
        message = consumer.consume_data()
        for msg in message:
            print(msg)


if __name__ == "__main__":
    xtype = sys.argv[1]
    main(xtype)


#
#     if xtype == 'c':
#         # 消费模块
#         consumer = Kafka_consumer(KAFAKA_HOST, KAFAKA_PORT, KAFAKA_TOPIC, group)
#         print("===========> consumer:", consumer)
#         message = consumer.consume_data()
#         for msg in message:
#             print('msg---------------->k,v', msg.key, msg.value)
#             print('offset---------------->', msg.offset)
#
#
# if __name__ == '__main__':
#     xtype = sys.argv[1]
#     group = sys.argv[2]
#     key = sys.argv[3]
#     main(xtype, group, key)
