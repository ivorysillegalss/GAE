
from kafka import KafkaProducer
import json
import logging
from kafka_producer.config import KAFKA_BROKER, KAFKA_TOPIC

# 初始化日志记录器
logger = logging.getLogger(__name__)


class KafkaProducerService:
    def __init__(self):
        # 创建 Kafka 生产者实例，指定序列化器（这里使用 JSON 格式）
        self.producer = KafkaProducer(
            bootstrap_servers=KAFKA_BROKER,
            value_serializer=lambda v: json.dumps(v).encode('utf-8')
        )
        logger.info(f"Kafka Producer 已连接到 {KAFKA_BROKER}")

    def send_message(self, message):
        """
        发送消息到指定 Kafka 主题
        :param message: dict, 要发送的数据，需为 JSON 序列化格式
        """
        try:
            # 发送消息到 Kafka 主题
            future = self.producer.send(KAFKA_TOPIC, value=message)
            # 处理消息发送后的回调
            result = future.get(timeout=10)  # 可设置超时时间
            logger.info(f"消息已发送到 Kafka 主题 {KAFKA_TOPIC}: {message}")
            return result
        except Exception as e:
            logger.error(f"发送消息到 Kafka 失败: {e}")
            raise e

    def close(self):
        """
        关闭生产者实例
        """
        self.producer.close()
        logger.info("Kafka Producer 已关闭")
