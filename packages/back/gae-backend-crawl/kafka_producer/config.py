
import os
from dotenv import load_dotenv

# 加载环境变量
load_dotenv()

# Kafka 配置
KAFKA_BROKER = os.getenv("KAFKA_BROKER", "127.0.0.1:9092")
KAFKA_TOPIC = os.getenv("KAFKA_TOPIC", "gaeMessageTopic")
KAFKA_GROUP_ID = os.getenv("KAFKA_GROUP_ID", "gaeMessageGroup")
