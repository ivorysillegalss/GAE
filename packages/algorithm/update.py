'''
用于更新数据
'''

import os

import pandas as pd
import pyarrow.parquet as pq

os.chdir(os.path.dirname(os.path.abspath(__file__)))
from conflag import *


class Update():
    def __init__(self) -> None:
        pass

    def update_R1(self) -> None:
        # 命名为：data/R_特征.parquet 
        df = pq.ParquetDataset("data/contributor_output.parquet").read().to_pandas()
        print(df)
        for i in worker_use_list:
            temp = df[i].tolist()
            temp = list(set(temp)).sort()
            temp = pd.DataFrame({'u': temp})
            temp.to_parquet(f"data\R_{i}.parquet")

            # temp = pq.ParquetDataset("data/").read().to_pandas()

    def update_R2(self) -> None:
        df = pq.ParquetDataset("data/repo_output.parquet")

        for i in project_use_list:
            temp = df[i].tolist()
            temp = list(set(temp)).sort()
            temp = pd.DataFrame({'u': temp})
            temp.to_parquet(f"data\R_{i}.parquet")

    def check_and_update(self) -> bool:
        flag = False
        df1 = pd.DataFrame()
        try:
            df1 = pq.ParquetDataset("update_data/contributor_output.parquet").read().to_pandas()
            df1.columns = df1.columns.map(info_worker)
        except:
            pass

        if len(df1) != 0:
            # 更新contributor 
            tmp = pq.ParquetDataset("data/contributor_output.parquet").read().to_pandas()
            # 按行合并
            merged_df = pd.concat([tmp, df1], ignore_index=True)
            merged_df = merged_df.drop_duplicates(subset='ID', keep='last')
            merged_df.to_parquet('data/contributor_output.parquet')
            self.update_R1()
            os.remove("update_data/contributor_output.parquet")
            flag = True

        df2 = pd.DataFrame()
        try:
            df2 = pq.ParquetDataset("update_data/repo_output.parquet").read().to_pandas()
            df2.columns = df2.columns.map(info_project)

        except:
            pass
        if len(df2) != 0:
            # 更新contributor 
            tmp = pq.ParquetDataset("data/repo_output.parquet").read().to_pandas()
            # 按行合并
            merged_df = pd.concat([tmp, df2], ignore_index=True)
            merged_df = merged_df.drop_duplicates(subset='仓库ID', keep='last')
            merged_df.to_parquet('data/repo_output.parquet')
            self.update_R2()
            os.remove("update_data/repo_output.parquet")
            flag = True
        return flag


Update_machine = Update()
flag = Update_machine.check_and_update()
if flag:
    pass
