package org.chenzc.gea.constant;


/**
 * @author chenz
 * @date 2024/10/26
 */
public class FlinkConstant {

    private FlinkConstant(){

    }

    /**
     * kafka消息清洗相关配置信息
     */

    public static final String MESSAGE_GROUP_ID = "gaeMessageGroup";
    public static final String MESSAGE_TOPIC_NAME = "gaeMessageTopic";
    public static final String MESSAGE_BROKER = "gae-kafka:9092";


    /**
     * redis相关配置
     */
    public static final String REDIS_IP = "127.0.0.1";
    public static final String REDIS_PORT = "6379";
    public static final String REDIS_PASSWORD = "";

    public static final String SOURCE_NAME = "gae_kafka_source";
    public static final String TRACK_EVENT_TYPE_FUNCTION_NAME = "gae_transfer";
    public static final String SINK_NAME = "gae_sink";
    public static final String JOB_NAME = "GaeBootStrap";

    public static final Integer MONTH = 30;
}
