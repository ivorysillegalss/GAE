'''
根据评价算法得到排行榜和图数据
'''

import pandas as pd 
import pyarrow.parquet as pq
import pyarrow as pa
import pandas as pd
import os 
import pickle
os.chdir(os.path.dirname(os.path.abspath(__file__)))
from conflag import *
from typing import List

dfs = dict() 
for i in ['Forks数量','网络数量','开放问题数量','星标数量','订阅者数量','观看者数量']:
    dfs[i]=pq.ParquetDataset(f"data/R_{i}.parquet").read().to_pandas()  
for i in ['公开仓库数量', '公开 Gist 数量', '粉丝数量', '关注数量']:
    dfs[i]=pq.ParquetDataset(f"data/R_{i}.parquet").read().to_pandas()

def R_F(va: float, ty: str) -> float:  
    global dfs 
    t = dfs[ty]['u'].tolist()
    n = len(t) 
    l = 0
    r = n - 1 
    while l < r: 
        mid = l + r >> 1 
        if t[mid] >= va: 
            r = mid 
        else: 
            l = mid + 1 
    res = 1 - (l + 1) / n 
    if len(t) == 1: res = 1 
    return res 

# project_use_list = ['Forks数量','网络数量','开放问题数量','星标数量','订阅者数量','观看者数量']
def calc_project(L: List[float]) -> float: 
    res = L[0] * 0.3 + L[1] * 0.1 + L[2] * 0.15 + L[3] * 0.4 + L[4] * 0.025 + L[5] * 0.025 
    return res 

# worker_use_list = ['公开仓库数量', '公开 Gist 数量', '粉丝数量', '关注数量']
def calc_worker(L: List[float]) -> float: 
    res = L[0] * 0.4 + L[1] * 0.3 + L[2] * 0.15 + L[3] * 0.15  
    return res 


if __name__ == "__main__": 
    # 个人知名度综合
    ones, score = [], []
    df = pq.ParquetDataset("data/contributor_output.parquet").read().to_pandas()
    for i in df.index:
        ones.append(df.loc[i,'ID']) 
        L = [] 
        for j in ['公开仓库数量', '公开 Gist 数量', '粉丝数量', '关注数量']:
            L.append(R_F((df.loc[i,j]), j))
        score.append(calc_worker(L))    

    who = dict() 
    for i in ones:
        who[i] = [] 
    o=pd.DataFrame({'ID':ones, 'score':score})
    o.to_parquet("result/Woker.parquet",engine="pyarrow")

    # 项目排行榜
    df = pq.ParquetDataset("data/repo_output.parquet").read().to_pandas()
    ones = []
    val = []

    for i in df.index:
        L = [] 
        for j in ['Forks数量','网络数量','开放问题数量','星标数量','订阅者数量','观看者数量']:
            L.append(R_F((df.loc[i,j]), j))
        va = calc_project(L)
        di = df.loc[i,'语言']
        su = 0 
        for j in range(len(di)): 
            su += di[j]['amount'] 
        dic={}
        if len(di) == 0 or su == 0: 
            pass  
        else: 
            for j in range(len(di)): 
                w = di[j]['amount'] / su 
                if di[j]['name'] not in languages_dict: continue 
                for s in languages_dict[di[j]['name']]: 
                    dic[s]=0 
                for s in languages_dict[di[j]['name']]: 
                    dic[s]+=va*w 
            dic = dict(sorted(dic.items(), key=lambda item: item[1], reverse=True))
        val.append(dic) 
        ones.append(df.loc[i,'仓库ID']) 
        Len = len(df.loc[i,'贡献者ID'])
        c = 1 
        for j in df.loc[i,'贡献者ID']:
            dj = dic.copy() 
            if int(j) in who: 
                po = 1 - c / Len
                if Len == 1: po = 1 
                dj = {key: value * po for key, value in dj.items()}
                who[int(j)].append(dj) 
                # print(dj)
            c += 1 

    # print(who)
    
    with open('data/project.pkl', 'wb') as f:
        pickle.dump(who, f)

    # pd.DataFrame({'ID':list(who.keys()), 'score': list(who.values())}).to_parquet("result/project.parquet")
    
