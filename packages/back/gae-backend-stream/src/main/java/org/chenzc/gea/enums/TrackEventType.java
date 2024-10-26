package org.chenzc.gea.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

/**
 * 时间追踪埋点 （记录消息发送的进度）
 * @author chenz
 * @date 2024/05/28
 */
@Getter
@AllArgsConstructor
public enum TrackEventType implements PowerfulEnums{

    /**
     * 成功处理消息
     */
    RECEIVE("10","消息接受成功"),

    /**
     * 责任链中丢弃消费 （删除模板）
     */
    DISCARD("20","消息被丢弃"),

    /**
     * 夜间屏蔽消息
     */
    NIGHT_SHIELD("22","夜间屏蔽"),

    /**
     * 夜间屏蔽 次日发送
     */
    NIGHT_SHIELD_NEXT_SEND("24","夜间屏蔽 次日发送"),

    /**
     * 内容去重
     */
    DEDUPLICATION_CONTENT("40","消息正在内容去重"),

    /**
     * 频次去重
     */
    DEDUPLICATION_RULE("42","消息正在频次去重"),

    /**
     * 下发成功
     */
    SUCCESS_SEND("60","消息下发成功"),

    /**
     * 下发失败
     */
    FAIL_SEND("62","消息下发失败");

    private final String code;
    private final String message;
}
