package org.chenzc.gea;

import com.google.common.base.Throwables;
import lombok.extern.slf4j.Slf4j;
import org.apache.flink.api.common.eventtime.WatermarkStrategy;
import org.apache.flink.connector.kafka.source.KafkaSource;
import org.apache.flink.streaming.api.datastream.DataStreamSource;
import org.apache.flink.streaming.api.datastream.SingleOutputStreamOperator;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.chenzc.gea.constant.FlinkConstant;
import org.chenzc.gea.entity.TrackEventEntity;
import org.chenzc.gea.function.EventTypeTransformation;
import org.chenzc.gea.sink.KafkaSink;
import org.chenzc.gea.utils.MessageQueueUtils;

/**
 * flink启动类
 * <p></>
 * 在flink中 数据处理的流程可分为三个步骤：
 * 创建数据源 Source 从外部系统读取数据
 * 业务处理逻辑 Transformation 对读取进来的数据进行各种转换和处理操作
 * 数据持久化 Sink 将处理后的数据写入外部系统
 *
 * @author chenz
 * @date 2024/06/06
 */


@Slf4j
public class CommuniBootStrap {

    public static void main(String[] args) {
        StreamExecutionEnvironment executionEnvironment = StreamExecutionEnvironment.getExecutionEnvironment();

//        1. 获取source
//        获取kafkaConsumer
        KafkaSource<String> kafkaConsumer = MessageQueueUtils.getKafkaConsumer(
                FlinkConstant.MESSAGE_TOPIC_NAME, FlinkConstant.MESSAGE_GROUP_ID, FlinkConstant.MESSAGE_BROKER);

        DataStreamSource<String> dataSource = executionEnvironment.fromSource(kafkaConsumer,
                WatermarkStrategy.noWatermarks(), FlinkConstant.SOURCE_NAME);


//        2. 经过Transformation对数据进行操作

//        数据清洗具体执行的操作：
//        数据是否完整合法：

        SingleOutputStreamOperator<TrackEventEntity> outputDataSource = dataSource.flatMap(new EventTypeTransformation())
                .name(FlinkConstant.TRACK_EVENT_TYPE_FUNCTION_NAME);

//        3. 将处理后的数据发送MQ
        outputDataSource.addSink(new KafkaSink()).name(FlinkConstant.SINK_NAME);

        try {
            executionEnvironment.execute(FlinkConstant.JOB_NAME);
        } catch (Exception e) {
//            log.error("flink word log error!");.
            Throwables.getStackTraceAsString(e);
        }
    }

}
