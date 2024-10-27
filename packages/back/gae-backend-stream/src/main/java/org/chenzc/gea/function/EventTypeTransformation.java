package org.chenzc.gea.function;

import com.alibaba.fastjson.JSON;
import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.util.Collector;
import org.chenzc.gea.entity.TrackEventEntity;

/**
 *
 * 数据清洗
 * @author chenz
 * @date 2024/06/06
 */
public class EventTypeTransformation implements FlatMapFunction<String, TrackEventEntity> {
    @Override
    public void flatMap(String value, Collector<TrackEventEntity> collector) throws Exception {
        TrackEventEntity trackEventEntity = JSON.parseObject(value, TrackEventEntity.class);
        collector.collect(trackEventEntity);
    }
}
