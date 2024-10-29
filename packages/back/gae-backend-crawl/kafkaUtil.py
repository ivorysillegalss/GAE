from kafka import KafkaProducer
import json

# Kafka 生产者配置
producer = KafkaProducer(
    bootstrap_servers='127.0.0.1:9092',  # Java消费者中指定的 broker
    value_serializer=lambda v: json.dumps(v).encode('utf-8')  # 消息序列化为JSON格式
)

# 发送消息到Java消费者指定的主题
message = {"example_key": "example_value"}
producer.send('gaeMessageTopic', value=message)  # 主题名称应与Java消费者配置一致
producer.flush()  # 确保消息成功发送
print("消息已发送到Kafka")
