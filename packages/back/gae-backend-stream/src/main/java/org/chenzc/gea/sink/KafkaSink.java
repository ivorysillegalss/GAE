package org.chenzc.gea.sink;

import com.alibaba.fastjson2.JSON;
import com.google.common.base.Throwables;
import lombok.extern.slf4j.Slf4j;
import org.apache.flink.streaming.api.functions.sink.SinkFunction;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.chenzc.gea.constant.FlinkConstant;
import org.chenzc.gea.entity.TrackEventEntity;
import org.chenzc.gea.utils.MessageQueueUtils;

@Slf4j
public class KafkaSink implements SinkFunction<TrackEventEntity> {

    private final KafkaProducer<String, String> producer;

    public KafkaSink() {
        // 初始化配置加载器和Kafka主题
        this.producer = MessageQueueUtils.getKafkaProducer(FlinkConstant.MESSAGE_BROKER);
    }

    @Override
    public void invoke(TrackEventEntity value, Context context) {
        try {

            String jsonString = JSON.toJSONString(value);
            // 创建 Kafka 消息
            ProducerRecord<String, String> record = new ProducerRecord<>(FlinkConstant.MESSAGE_TOPIC_NAME, jsonString);

            // 发送消息到 Kafka
            producer.send(record, (RecordMetadata metadata, Exception exception) -> {
                if (exception != null) {
//                    log.error("Failed to send message to Kafka: {}", Throwables.getStackTraceAsString(exception));
                    System.out.println("Failed to send message to Kafka");
                } else {
//                    log.info("Message sent to Kafka: topic={}, partition={}, offset={}", metadata.topic(), metadata.partition(), metadata.offset());
                    System.out.println("Message sent to Kafka");
                }
            });
        } catch (Exception e) {
//            log.error("KafkaSink invoke error : {}", Throwables.getStackTraceAsString(e));
            System.out.println("KafkaSink invoke error");
        }
    }
}
