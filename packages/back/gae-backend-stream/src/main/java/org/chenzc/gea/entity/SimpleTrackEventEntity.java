package org.chenzc.gea.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * 记录简单的埋点信息（sink持久化）
 * @author chenz
 * @date 2024/06/06
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class SimpleTrackEventEntity {
    /**
     * 具体点位
     */
    private String state;

    /**
     * 业务Id(数据追踪使用)
     * 生成逻辑参考 TaskInfoUtils
     */
    private Long businessId;

    /**
     * 生成时间
     */
    private long timestamp;
}
