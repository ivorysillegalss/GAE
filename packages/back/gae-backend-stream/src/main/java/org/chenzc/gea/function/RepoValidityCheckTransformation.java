package org.chenzc.gea.function;

import com.alibaba.fastjson.JSON;
import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.util.Collector;
import org.chenzc.gea.entity.RepoEntity;

public class RepoValidityCheckTransformation implements FlatMapFunction<String, RepoEntity> {
    @Override
    public void flatMap(String s, Collector<RepoEntity> collector) throws Exception {
        RepoEntity repoEntity = JSON.parseObject(s, RepoEntity.class);
//        TODO 数据清洗其他逻辑
        collector.collect(repoEntity);
    }
}
