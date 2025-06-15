import sys
import unittest
from collections import defaultdict
from typing import List, Tuple


def get_sorted_dict(queries: List[Tuple[int, str]]) -> List[str]:
    # 年, 名前でソート
    sorted_events = sorted(queries, key=lambda x: (x[0], x[1]))
    return [name for _, name in sorted_events]


class TestItemPriceMap(unittest.TestCase):
    def test_get_sorted_dict(self):
        expected = [
            'hiro',
            'nao',
            'yuki',
            'nao',
            'nao'
        ]
        input_data = [
            (645, "nao"),
            (593, "hiro"),
            (2058, "yuki"),
            (29484, "nao"),
            (374759, "nao"),
        ]
        actual = get_sorted_dict(input_data)
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    _ = [input().strip() for _ in range(item_count)]
    queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    for v in get_sorted_dict(queries):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 西暦 1,000,000,000 年に行われた歴史の授業のグループワークで、歴史上のいくつかの出来事についての記事を年代順に並べて歴史年表を作成することになりました。
# ところが、歴史年表は 1 枚の紙にまとめる必要があるため、古い出来事を担当する人から順番に歴史年表を書くことにしました。
# グループの人数 N とそのメンバー S_1 ... S_N が与えられます。
# 続けて、歴史年表に載せる出来事の数 K , 各出来事の起こった年 Y_i , その出来事の記事を担当する生徒の名前 C_i が与えられるので、歴史年表を書く担当者の順番を出力してください。
# なお、 1 人の生徒が複数の出来事の記事を担当することがある点に注意してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# S_1
# ...
# S_N
# Y_1 C_1
# ...
# Y_K C_K


# ・1 行目では、グループの人数 N と歴史年表に載せる出来事の数 K が与えられます。
# ・続く N 行のうち i 行目では、 i 人目のメンバーの名前 S_i が与えられます。
# ・続く K 行のうち i 行目では、 i 個目の出来事の起こった年 Y_i と、その記事を担当する生徒の名前 C_i が先頭から順に与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ・歴史年表を書く担当者の名前を順番に出力してください。
# ・ただし、同じ年に複数の出来事があった場合は、名前が辞書順になるように出力してください。
# ・また、出力の末尾には改行を入れてください。

# 条件
# ・1 ≦ N ≦ 1,000
# ・1 ≦ K ≦ 100,000
# ・S_i は 20 文字以下の英数字列 (1 ≦ i ≦ N)
# ・S_i ≠ S_j (i ≠ j)
# ・1 ≦ Y_i ≦ 1,000,000,000 (1 ≦ i ≦ K)
# ・C_i は S_j のいずれか (1 ≦ i ≦ K , 1 ≦ j ≦ N)

# 入力例1
# 3 5
# nao
# hiro
# yuki
# 645 nao
# 593 hiro
# 2058 yuki
# 29484 nao
# 374759 nao

# 出力例1
# hiro
# nao
# yuki
# nao
# nao

# 入力例2
# 5 10
# aoi
# ikoka
# en
# ron
# nana
# 463 nana
# 7583 nana
# 5839 nana
# 17274 nana
# 3773 nana
# 264 nana
# 7485 nana
# 24855 nana
# 395385 nana
# 5355 nana

# 出力例2
# nana
# nana
# nana
# nana
# nana
# nana
# nana
# nana
# nana
# nana
