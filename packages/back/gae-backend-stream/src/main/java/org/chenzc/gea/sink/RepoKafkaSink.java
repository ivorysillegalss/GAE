package org.chenzc.gea.sink;

import com.alibaba.fastjson2.JSON;
import lombok.extern.slf4j.Slf4j;
import org.apache.flink.streaming.api.functions.sink.SinkFunction;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.chenzc.gea.constant.FlinkConstant;
import org.chenzc.gea.entity.RepoEntity;
import org.chenzc.gea.utils.MessageQueueUtils;

@Slf4j
public class RepoKafkaSink implements SinkFunction<RepoEntity> {

    private final KafkaProducer<String, String> producer;

    public RepoKafkaSink() {
        // 初始化配置加载器和Kafka主题
        this.producer = MessageQueueUtils.getKafkaProducer(FlinkConstant.MESSAGE_BROKER);
    }

    @Override
    public void invoke(RepoEntity value, Context context) {
        try {

            String jsonString = JSON.toJSONString(value);
            // 创建 Kafka 消息
            ProducerRecord<String, String> record = new ProducerRecord<>(FlinkConstant.CLEANSING_REPO_NAME, jsonString);

            // 发送消息到 Kafka
            producer.send(record, (RecordMetadata metadata, Exception exception) -> {
                if (exception != null) {
                    System.out.println("Failed to send message to Kafka");
                } else {
                    System.out.println("Message sent to Kafka");
                }
            });
        } catch (Exception e) {
            System.out.println("RepoKafkaSink invoke error");
        }
    }
}
