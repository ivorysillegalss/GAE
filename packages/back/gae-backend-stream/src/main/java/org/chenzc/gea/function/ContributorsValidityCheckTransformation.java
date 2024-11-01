package org.chenzc.gea.function;

import com.alibaba.fastjson.JSON;
import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.util.Collector;
import org.chenzc.gea.entity.ContributorEntity;

public class ContributorsValidityCheckTransformation implements FlatMapFunction<String, ContributorEntity> {
    @Override
    public void flatMap(String s, Collector<ContributorEntity> collector) throws Exception {
        ContributorEntity contributor = JSON.parseObject(s, ContributorEntity.class);
//        TODO 数据清洗其他逻辑
        collector.collect(contributor);
    }
}
