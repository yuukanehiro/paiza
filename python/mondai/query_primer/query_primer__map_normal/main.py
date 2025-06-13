import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_query_merge_map(items: Dict[str, str], queries: List[str]) -> List[str]:
    result = []
    for q in queries:
        q_items = q.split()
        if q_items[0] == "call":
            result.append(items[q_items[1]])
        elif q_items[0] == "leave":
            items.pop(q_items[1])
        elif q_items[0] == "join":
            items[q_items[1]] = q_items[2]

    return result


class TestItemPriceMap(unittest.TestCase):
    def test_get_query_merge_map(self):
        expected = ["Yui", "Sakura"]
        actual = get_query_merge_map({"1": "Sin", "2": "Sakura", "3": "Kayo", "4": "Yui"}, ["call 4", "leave 2", "join 2 Sakura", "call 2"])
        self.assertEqual(expected, actual)

def main():
    item_count, query_count = map(int, input().split())
    items: Dict[str, str] = dict(input().split() for _ in range(item_count))
    queries = [input().strip() for _ in range(query_count)]

    for v in get_query_merge_map(items, queries):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()


# Q
# 下記の問題をプログラミングしてみよう！
# 3xxx年、ロボット学校の先生である paiza 君は、新しく担当するクラスの生徒一人一人の出席番号と識別 ID を覚えて、出席番号が与えられたら、その生徒の識別 ID を言えるようになる必要があります。
# paiza 君の務める学校は転校が多く、頻繁に生徒が増減します。

# 覚えるべき生徒の出席番号と識別 ID が与えられたのち、いくつかのイベントを表す文字列が与えられるので、与えられた順に各イベントに応じて次のような処理をおこなってください。

# ・join num id
# 生徒番号 num , 識別ID id の生徒を新たに覚える

# ・leave num
# 生徒番号 num の生徒を忘れる

# ・call num
# 生徒番号 num の生徒の識別 ID を出力する

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# num_1 ID_1
# ...
# num_N ID_N
# S_1
# ...
# S_K


# ・1 行目では、初めに覚える生徒の人数 N と与えられるイベントの回数 K が与えられます。
# ・続く N 行のうち i 行目 (1 ≦ i ≦ N) では、i 番目の生徒の出席番号と識別 ID の組 num_i , ID_i が半角スペース区切りで与えられます。
# ・続く K 行では、起きるイベントを表す文字列 S_i (1 ≦ i ≦ K) が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ・イベントに従って出力してください。
# ・また、出力の末尾には改行を入れてください。

# 条件
# ・1 ≦ N , K ≦ 100,000
# ・1 ≦ num_i ≦ 1,000,000 (1 ≦ i ≦ N)
# ・num_i ≠ num_j (i ≠ j)
# ・ID_i は アルファベット大文字小文字と数字から成る 20 文字以下の文字列 (1 ≦ i ≦ N)
# ・S_i は次のいずれかの形式である。


# ・join num id
# 生徒番号 num , 識別 ID id の生徒を新たに覚える。

# ・leave num
# 生徒番号 num の生徒を忘れる。

# ・call num
# 生徒番号 num の生徒の識別 ID を出力する。
# この時点で生徒番号 num の生徒がいることは保証されている。

# 1 ≦ num ≦ 1,000,000
# id は 20 文字以下の文字列
# 入力例1
# 4 4
# 1 Sin
# 2 Sakura
# 3 Kayo
# 4 Yui
# call 4
# leave 2
# join 2 Sakuya
# call 2

# 出力例1
# Yui
# Sakuya

# 入力例2
# 5 10
# 696042 pieF4
# 162082 Geig1
# 43482 Ich7D
# 647458 foh8C
# 71317 Aiv4g
# call 43482
# call 696042
# call 696042
# leave 696042
# call 647458
# call 647458
# call 162082
# join 591845 Ue7wo
# call 591845
# leave 647458

# 出力例2
# Ich7D
# pieF4
# pieF4
# foh8C
# foh8C
# Geig1
# Ue7wo
