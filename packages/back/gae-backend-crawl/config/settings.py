# config/settings.py
import os
from dotenv import load_dotenv

# 加载 .env 文件中的环境变量
load_dotenv()

# 配置变量
KAFKA_BROKER = os.getenv("KAFKA_BROKER", "localhost:9092")
KAFKA_TOPIC = os.getenv("KAFKA_TOPIC", "gaeMessageTopic")
DB_HOST = os.getenv("DB_HOST", "localhost")
DB_USER = os.getenv("DB_USER", "user")
DB_PASS = os.getenv("DB_PASS", "password")
