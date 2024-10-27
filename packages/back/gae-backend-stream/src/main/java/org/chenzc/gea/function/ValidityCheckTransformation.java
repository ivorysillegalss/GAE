package org.chenzc.gea.function;

import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.util.Collector;
import org.chenzc.gea.entity.TrackEventEntity;

public class ValidityCheckTransformation implements FlatMapFunction<TrackEventEntity,TrackEventEntity> {
    @Override
    public void flatMap(TrackEventEntity trackEventEntity, Collector<TrackEventEntity> collector) throws Exception {
//        TODO 数据清洗逻辑
        collector.collect(trackEventEntity);
    }
}
