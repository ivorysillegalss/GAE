package org.chenzc.gea.constant;


/**
 * @author chenz
 * @date 2024/10/26
 */
public class FlinkConstant {

    private FlinkConstant() {

    }

    /**
     * kafka消息清洗相关配置信息
     */

    public static final String CLEANSING_REPO_NAME = "gaeCleansingRepoTopic";
    public static final String CLEANSING_CONTRIBUTORS_NAME = "gaeCleansingContributorsTopic";

    //    TODO kafkaip配置
//    public static final String MESSAGE_BROKER = "gae-kafka:9092";
    public static final String MESSAGE_BROKER = "127.0.0.1:9092";


    public static final String UN_CLEANSING_REPO_GROUP = "gaeUnCleansingRepoGroup";
    public static final String UN_CLEANSING_REPO_TOPIC = "gaeUnCleansingRepoTopic";
    public static final String CONTRIBUTORS_VALIDLY_CHECK_FUNCTION_NAME = "contributors_check";

    public static final String UN_CLEANSING_USER_GROUP = "gaeUnCleansingUserGroup";
    public static final String UN_CLEANSING_USER_TOPIC = "gaeUnCleansingUserTopic";
    public static final String REPO_VALIDLY_CHECK_FUNCTION_NAME = "repo_check";


    public static final String SOURCE_NAME = "gae_kafka_source";
    public static final String SINK_NAME = "gae_sink";
    public static final String REPO_JOB_NAME = "RepoBootStrap";
    public static final String CONTRIBUTORS_JOB_NAME = "ContributorsBootStrap";

    public static final Integer MONTH = 30;
}
