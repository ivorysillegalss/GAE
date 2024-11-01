package org.chenzc.gea;

import com.google.common.base.Throwables;
import lombok.extern.slf4j.Slf4j;
import org.apache.flink.api.common.eventtime.WatermarkStrategy;
import org.apache.flink.connector.kafka.source.KafkaSource;
import org.apache.flink.streaming.api.datastream.DataStreamSource;
import org.apache.flink.streaming.api.datastream.SingleOutputStreamOperator;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.chenzc.gea.constant.FlinkConstant;
import org.chenzc.gea.entity.ContributorEntity;
import org.chenzc.gea.entity.RepoEntity;
import org.chenzc.gea.function.ContributorsValidityCheckTransformation;
import org.chenzc.gea.function.RepoValidityCheckTransformation;
import org.chenzc.gea.sink.ContributorsKafkaSink;
import org.chenzc.gea.sink.RepoKafkaSink;
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
public class GaeBootStrap {

    public static void main(String[] args) {
        StreamExecutionEnvironment executionEnvironment1 = StreamExecutionEnvironment.getExecutionEnvironment();
        StreamExecutionEnvironment executionEnvironment2 = StreamExecutionEnvironment.getExecutionEnvironment();

//        1. 获取source
//        获取kafkaConsumer
        KafkaSource<String> kafkaConsumer1 = MessageQueueUtils.getKafkaConsumer(
                FlinkConstant.UN_CLEANSING_REPO_TOPIC, FlinkConstant.UN_CLEANSING_REPO_GROUP, FlinkConstant.MESSAGE_BROKER);

        KafkaSource<String> kafkaConsumer2 = MessageQueueUtils.getKafkaConsumer(
                FlinkConstant.UN_CLEANSING_USER_TOPIC, FlinkConstant.UN_CLEANSING_USER_GROUP, FlinkConstant.MESSAGE_BROKER);

        DataStreamSource<String> dataSource1 = executionEnvironment1.fromSource(kafkaConsumer1,
                WatermarkStrategy.noWatermarks(), FlinkConstant.SOURCE_NAME);

        DataStreamSource<String> dataSource2 = executionEnvironment2.fromSource(kafkaConsumer2,
                WatermarkStrategy.noWatermarks(), FlinkConstant.SOURCE_NAME);

//        2. 经过Transformation对数据进行操作

//        数据清洗具体执行的操作：
//        数据是否完整合法：
        SingleOutputStreamOperator<RepoEntity> outputDataSource1 = dataSource1.flatMap(new RepoValidityCheckTransformation())
                .name(FlinkConstant.REPO_VALIDLY_CHECK_FUNCTION_NAME);

        SingleOutputStreamOperator<ContributorEntity> outputDataSource2 = dataSource2.flatMap(new ContributorsValidityCheckTransformation())
                .name(FlinkConstant.CONTRIBUTORS_VALIDLY_CHECK_FUNCTION_NAME);

//        3. 将处理后的数据发送MQ
        outputDataSource1.addSink(new RepoKafkaSink()).name(FlinkConstant.SINK_NAME);

        outputDataSource2.addSink(new ContributorsKafkaSink()).name(FlinkConstant.SINK_NAME);

        try {
            executionEnvironment1.execute(FlinkConstant.REPO_JOB_NAME);
        } catch (Exception e) {
            Throwables.getStackTraceAsString(e);
        }

        try {
            executionEnvironment2.execute(FlinkConstant.CONTRIBUTORS_JOB_NAME);
        } catch (Exception e) {
            Throwables.getStackTraceAsString(e);
        }
    }

}
