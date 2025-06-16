import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_result_list(items: List[str], queries: List[str]) -> List[str]:
    result = []

    accounting = defaultdict(dict)
    for v in items:
        accounting[v] = {}

    for q in queries:
        name, no, price = q.split()
        if name in accounting:
            accounting[name][no] = int(price)

    for a in accounting:
        result.append(a)
        if a in accounting:
            for no in accounting[a]:
                result.append(f"{no} {accounting[a][no]}")
        result.append("-----")

    return result


class TestItemPriceMap(unittest.TestCase):
    def test_get_item_price_map(self):
        expected = [
            "A",
            "1 100",
            "3 500",
            "6 2685",
            "-----",
            "B",
            "2 100",
            "-----",
            "C",
            "4 895",
            "5 890",
            "-----",
        ]
        input_items = [
            "A",
            "B",
            "C"
        ]
        input_queries = [
            "A 1 100",
            "B 2 100",
            "A 3 500",
            "C 4 895",
            "C 5 890",
            "A 6 2685",
        ]
        actual = get_result_list(input_items, input_queries)
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    items = [input().strip() for _ in range(item_count)]
    queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    for v in get_result_list(items, queries):
        print(v)



if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# paiza には N 個の部署があり、名前はそれぞれ S_1 ... S_N です。
# 経理係となったあなたは、どの部署が何円のどのような買い物をしたかを記録するように言われました。
# どの部署が何円で何を買ったかの領収書が K 枚与えられるので、各部署の会計表を作成しましょう。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# S_1
# ...
# S_N
# A_1 P_1 M_1
# ...
# A_K P_K M_K


# ・1 行目では、部署の数 N と与えられる領収書の枚数 K が与えられます。
# ・続く N 行のうち、 i 行目では、i 番目に登録されている部署名 S_i が与えられます。
# ・続く K 行のうち、 i 行目では、 i 枚目の領収書に書かれていた部署名 A_i , 注文番号 P_i , その金額 M_i が半角スペース区切りで与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# S_1
# P_1 M_1
# ...
# -----
# S_2
# P_1 M_1
# ...
# -----
# ...
# -----
# S_N
# P_1 M_1
# ...
# -----


# ・各部署について次の形式で出力してください。
# ・1 行目では、登録されている部署名 S_i を出力してください。
# ・2 行目以降には、注文番号 P_1 ... とその値段 M_1 ... を入力された順に半角スペース区切りで出力してください。
# ・各部署の出力の後ろに以下の通り区切りを出力してください。
# -----


# ・詳しくは入出力例を参考にしてください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・1 ≦ N ≦ 100,000
# ・1 ≦ K ≦ 100,000
# ・A_i は S_j のいずれか (1 ≦ i ≦ K , 1 ≦ j ≦ N)
# ・S_i は 20 文字以下の文字列 (1 ≦ i ≦ N)
# ・0 ≦ P_i ≦ 10^10 (1 ≦ i ≦ K)
# ・1 ≦ M_i ≦ 10,000 (1 ≦ i ≦ K)

# 入力例1
# 3 6
# A
# B
# C
# A 1 100
# B 2 100
# A 3 500
# C 4 895
# C 5 890
# A 6 2685

# 出力例1
# A
# 1 100
# 3 500
# 6 2685
# -----
# B
# 2 100
# -----
# C
# 4 895
# 5 890
# -----

# 入力例2
# 4 5
# ed
# bjd
# bdkf
# fkoe
# ed 20 2093
# ed 584 3388
# ed 31737 3885
# ed 023748 9300
# fkoe 82928 274

# 出力例2
# ed
# 20 2093
# 584 3388
# 31737 3885
# 023748 9300
# -----
# bjd
# -----
# bdkf
# -----
# fkoe
# 82928 274
# -----
