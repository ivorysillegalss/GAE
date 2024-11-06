import os
import pickle
from collections import deque
from typing import List

import pandas as pd
import pyarrow.parquet as pq

os.chdir(os.path.dirname(os.path.abspath(__file__)))


class Query():
    def __init__(self) -> None:
        super().__init__()
        self.nation_data = self.load_or_create_nation_data()
        self.graph_data = self.load_or_create_relation_data()

    '''
        # 得到排行榜，只会在更新时重置

        # 以项目评分为第一关键字，个人知名度评分为第二关键字的排行榜 
        
          ID | 个人知名度评分 | 项目个人评分 | 
           1        20.5           5.5 
           2        20.1           4.5

    '''

    # 全部人
    def query_Ranks(self, base: List[str]) -> None:
        with open(os.path.join("data", "project.pkl"), 'rb') as f:
            loaded_dict = pickle.load(f)
        df1 = pq.ParquetDataset(r"result\Woker.parquet").read().to_pandas()
        user_list = [int(i) for i in df1['ID'].tolist()]
        df1 = dict(zip(df1['ID'].tolist(), df1['score'].tolist()))
        dic = {}
        for i in user_list:
            dic[i] = (0, 0)
            if i in loaded_dict:
                mx = 0

                if base != []:
                    mx = dict()

                for j in loaded_dict[i]:
                    for k, v in j.items():
                        if base != []:
                            if k in base:
                                if k in mx:
                                    mx[k] = max(mx[k], v)
                                else:
                                    mx[k] = v
                        else:
                            mx = max(mx, v)

                if base == []:
                    dic[i] = (mx, 0)
                else:
                    if len(mx) == 0:
                        dic[i] = (0, 0)
                        continue
                    res = 0
                    for k, v in mx.items():
                        res += v
                    res /= len(mx)
                    dic[i] = (res, 0)

        for i in user_list:
            if i in df1:
                dic[i] = (dic[i][0], df1[i])

        result = sorted(dic.items(), key=lambda item: (item[1][0], item[1][1]), reverse=True)

        id, project_score, user_score = [], [], []
        for i, j in result:
            id.append(i)
            project_score.append(j[0])
            user_score.append(j[1])
        temp = pq.ParquetDataset(r"data\contributor_output.parquet").read().to_pandas()

        temp['ID'] = pd.Categorical(temp['ID'], categories=id, ordered=True)
        temp = temp.sort_values("ID")

        temp_list = ['登录名', '头像 URL', 'HTML 网址', 'Gravatar ID', '姓名', '公司', '博客', '位置', '邮箱',
                     '个人简介']
        df = {'ID': id, 'project_score': project_score, 'user_score': user_score}
        for i in temp_list:
            df[i] = temp[i].tolist()
        df = pd.DataFrame(df)
        df.to_parquet("result/all_rank.parquet")

        # 个人

    def query_Rank(self, id: int, Lis: List[str]) -> None:
        self.query_Ranks(Lis)
        df = pq.ParquetDataset(r"result\all_rank.parquet").read().to_pandas()
        res = pd.DataFrame(df[df['ID'] == id])
        res.to_parquet("result/one_rank.parquet")

    def load_or_create_relation_data(self):
        if os.path.exists('result/graph.pkl'):
            with open(os.path.join("result", "graph.pkl"), 'rb') as f:
                return pickle.load(f)
        else:
            self.query_nation()  # 运行国家推测
            with open(os.path.join("result", "graph.pkl"), 'rb') as f:
                return pickle.load(f)

    def load_or_create_nation_data(self):
        if os.path.exists('result/nation.pkl'):
            with open(os.path.join("result", "nation.pkl"), 'rb') as f:
                return pickle.load(f)
        else:
            self.query_nation()  # 运行国家推测
            with open(os.path.join("result", "nation.pkl"), 'rb') as f:
                return pickle.load(f)

    # 猜测国家
    def query_nation(self) -> None:
        # 存储图信息为：graph.parquet 
        # 存储nation信息为：Nation.parquet 
        df1 = pq.ParquetDataset(r"data\contributor_output.parquet").read().to_pandas()
        n = len(df1["ID"].tolist())
        user_ids = df1["ID"].tolist()

        print(df1.columns)

        usernames = df1["登录名"].tolist()
        avatar_urls = df1["头像 URL"].tolist()

        point_map = dict()
        f_pmap = dict()
        username_map = dict()
        avatar_urls_map = dict()

        c = 0
        for i, username, avatar_url in zip(df1['ID'].tolist(), usernames, avatar_urls):
            point_map[i] = c + 1
            avatar_urls_map[i] = str(avatar_url)
            username_map[i] = str(username)
            c += 1
            f_pmap[c] = i

        e, ne, h, w, idx = [0 for i in range(n * 10 + 1)], [0 for i in range(n * 10 + 1)], [-1 for i in
                                                                                            range(n * 10 + 1)], [0 for i
                                                                                                                 in
                                                                                                                 range(
                                                                                                                     n * 10 + 1)], 0

        def add_edge(a: int, b: int, c: int) -> None:
            nonlocal e, ne, h, w, idx, n
            e[idx] = b
            ne[idx] = h[a]
            w[idx] = c
            h[a] = idx
            idx += 1

        def deque_bfs(a: int) -> List[int]:
            nonlocal e, ne, h, w, idx, n
            dq = deque([a])
            dist = [10 ** 9 for i in range(n + 1)]
            dist[a] = 0
            while len(dq) != 0:
                t = dq.popleft()
                i = h[t]
                while i != -1:
                    j = e[i]
                    if j != a and dist[j] > dist[t] + w[i]:
                        dist[j] = dist[t] + w[i]
                        if w[i] == 1:
                            dq.appendleft(j)
                        else:
                            dq.append(j)
                    i = ne[i]
            return dist
            # g = [[0 for i in range(n + 1)] for j in range(n + 1)]

        for i in df1.index:
            I_d = point_map[df1.loc[i, "ID"]]
            gz = df1.loc[i, '关注列表']
            for j in gz:
                if int(j['id']) in point_map:
                    t_d = point_map[int(j['id'])]
                    add_edge(I_d, t_d, 2)
                    add_edge(t_d, I_d, 2)
        df1 = pq.ParquetDataset(r"data\repo_output.parquet").read().to_pandas()
        for i in df1['贡献者ID'].tolist():
            for j in range(len(i)):
                for k in range(j + 1, len(i)):
                    if i[j] in point_map and i[k] in point_map:
                        cj, ck = point_map[i[j]], point_map[i[k]]
                        add_edge(cj, ck, 1)
                        add_edge(ck, cj, 1)

        df1 = pq.ParquetDataset(r"data\contributor_output.parquet").read().to_pandas()
        countries = [
            "Argentina", "Australia", "Austria", "Azerbaijan", "Bahamas", "Pakistan",
            "Belarus", "Belgium", "Botswana", "Brazil", "Brunei", "Bulgaria",
            "Canada", "Chile", "China", "Colombia", "Croatia", "Cyprus",
            "Czech Republic", "Denmark", "Egypt", "Estonia", "Ethiopia", "Finland",
            "France", "Germany", "Greece", "Jamaica", "Japan", "Jordan", "Kazakhstan",
            "South Korea", "Kuwait", "Latvia", "Lebanon", "Lithuania", "Luxembourg",
            "Malaysia", "Maldives", "Mexico", "Monaco", "Mongolia", "Morocco",
            "Netherlands", "New Zealand", "Nigeria", "Norway", "Pakistan", "Philippines",
            "Poland", "Portugal", "Qatar", "Romania", "Russia", "Singapore",
            "Slovakia", "Slovenia", "South Africa", "Spain", "Sweden", "Switzerland",
            "Taiwan", "Thailand", "Turkey", "Ukraine", "United Arab Emirates", "United Kingdom",
            "United States", "Uzbekistan", "Vietnam"
        ]
        con = {}
        for i in df1.index:
            x = df1.loc[i, '位置']
            I_D = df1.loc[i, 'ID']
            con[I_D] = ""
            for j in countries:
                if j in x:
                    con[I_D] = j
                    break

        end_res = dict()
        for i in user_ids:
            dist = deque_bfs(point_map[i])

            dic_tmp = {}

            if con[i] == "":
                for o in range(1, n + 1):
                    low = f_pmap[o]
                    if con[low] != "":
                        if con[low] in dic_tmp:
                            dic_tmp[con[low]] = min(dist[o], dic_tmp[con[low]])
                        else:
                            dic_tmp[con[low]] = dist[o]
            else:
                dic_tmp[con[i]] = 0
            dic_tmp = dict(sorted(dic_tmp.items(), key=lambda item: item[1], reverse=False))

            result = []
            tl = 1
            for k, v in dic_tmp.items():
                rk = 1 - tl / len(dic_tmp)
                if rk == 0: rk = 1
                tl += 1
                result.append((k, rk))

            end_res[username_map.get(i, '')] = result
        # df = pd.DataFrame({'ID': ones, 'countries': end_res}) 
        # df.to_parquet("result/all_nation.parquet") 
        with open(os.path.join("result", "nation.pkl"), 'wb') as f:
            pickle.dump(end_res, f)

        end_res_1 = dict()
        for i in user_ids:
            Ls = []
            t = point_map[i]
            o = h[t]
            while o != -1:
                j = e[o]
                Ls.append(
                    (
                        username_map.get(f_pmap[j])
                        , int(w[o])
                        , avatar_urls_map.get(i, ''),
                    ))

                o = ne[o]
            end_res_1[username_map.get(i, '')] = Ls
            # print(i, Ls)

        with open(os.path.join("result", "graph.pkl"), 'wb') as f:
            pickle.dump(end_res_1, f)

    def get_nation_for_user(self, username: str) -> str:
        # 从加载的字典中获取指定用户的国家信息
        if username in self.nation_data:
            user_data = self.nation_data[username]
            return user_data
        else:
            return "未知"  # 如果用户不在结果中

    def get_relation_for_user(self, username: str) -> str:
        if username in self.graph_data:
            relation_data = self.graph_data[username]
            return relation_data
        else:
            return "未知"

    # if __name__ == "__main__":
#     rank = Query() 
# rank.query_Ranks(["Web 开发"])
# rank.query_nation()
