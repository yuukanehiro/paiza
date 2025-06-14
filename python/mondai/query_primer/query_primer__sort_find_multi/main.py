import sys
import unittest
from collections import defaultdict
from typing import List, Dict
from bisect import bisect_left

def get_insert_positions(items: List[int], queries: List[str], target: int) -> List[int]:
    result = []
    for q in queries:
        q_line = q.split()
        if q_line[0] == "join":
            items.append(int(q_line[1]))
        elif q_line[0] == "sorting":
            items = sorted(items)
            result.append((bisect_left(items, target) + 1))

    return result


class TestItemPriceMap(unittest.TestCase):
    def test_get_item_price_map_case1(self):
        expected = [5]
        actual = get_insert_positions([118, 174, 133], ["join 137", "join 177", "sorting"], 176)
        self.assertEqual(expected, actual)
    def test_get_item_price_map_case2(self):
        expected = [3, 3, 5, 5, 6]
        actual = get_insert_positions([169, 164, 162, 112, 191, 168, 168, 199, 176, 146], ["join 196", "join 142", "sorting", "sorting", "join 131", "join 140", "sorting", "sorting", "join 143", "sorting"], 145)
        self.assertEqual(expected, actual)

def main():
    item_count, query_count, targetHeight = map(int, input().split())
    items = [int(input().strip()) for _ in range(item_count)]
    queries = [input().strip() for _ in range(query_count)]

    for v in get_insert_positions(items, queries, targetHeight):
        print(v)

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# paiza 君のクラスには paiza 君を含めて N + 1 人の生徒がいます。paiza 君の身長は P cm で、他の N 人の生徒の身長はそれぞれ A_1 ... A_N です。
# このクラスには次のようなイベントが合計 K 回起こります。
# それぞれのイベントは以下のうちのいずれかです。

# ・転校生がクラスに加入する
# ・全員で背の順に並ぶ

# 全員で背の順で並ぶイベントが起こるたびに、そのとき paiza 君は前から何番目に並ぶことになるかを出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K P
# A_1
# ...
# A_N
# event_1
# ...
# event_K


# ・1 行目では、paiza 君を除いたクラスの人数 N と起こるイベントの回数 K と paiza君の身長 P が与えられます。
# ・続く N 行では、初めにクラスにいる N 人の生徒の身長が与えられます。
# ・続く K 行では、起こるイベントを表す文字列が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ・全員で背の順で並ぶイベントが起こるたびに、paiza 君が前から何番目に並ぶことになるかを出力してください。
# ・また、出力の末尾には改行を入れてください。

# 条件
# ・1 ≦ N , K ≦ 100,000
# ・100 ≦ P ≦ 200
# ・100 ≦ A_i ≦ 200 (1 ≦ i ≦ N)
# ・転校生を含め、クラスの中で P cm の生徒は paiza 君のみであることが保証されている
# ・event_i (1 ≦ i ≦ K) は以下のいずれかの形式で与えられる。



# join num

# 身長 num(cm) の生徒がクラスに加入したことを表す。


# sorting

# 生徒が背の順に並ぶことを表す
# この入力が与えられるたび、paiza 君が背の順で前から何番目に並ぶことになるかを出力してください。


# 入力例1
# 3 3 176
# 118
# 174
# 133
# join 137
# join 177
# sorting

# 出力例1
# 5

# 入力例2
# 10 10 145
# 169
# 164
# 162
# 112
# 191
# 168
# 168
# 199
# 176
# 146
# join 196
# join 142
# sorting
# sorting
# join 131
# join 140
# sorting
# sorting
# join 143
# sorting

# 出力例2
# 3
# 3
# 5
# 5
# 6
