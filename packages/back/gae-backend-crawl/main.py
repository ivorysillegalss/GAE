# main.py
import logging

from config.logging_config import setup_logging
from kafka_producer.producer import KafkaProducerService


def main():
    # 初始化日志配置
    setup_logging()
    logging.info("日志系统初始化成功")

    kafka_producer = KafkaProducerService()
    message = {}
    kafka_producer.send_message(message)
    kafka_producer.close()

    # 启动其他程序
    # 例如：启动爬虫、Kafka 生产者等


if __name__ == "__main__":
    main()
