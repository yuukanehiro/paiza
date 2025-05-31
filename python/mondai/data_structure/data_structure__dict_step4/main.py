import unittest
import sys
from collections import defaultdict
from typing import List, Dict

def get_item_table(lines: List[str]) -> Dict[str, int]:
    hashTable = defaultdict(str)

    for line in lines:
        itemName, price = line.split()
        hashTable[itemName] = int(price)

    return hashTable


class TestFruitCounter(unittest.TestCase):
    def test_get_item_table(self):
        want = {"eraser": 50, "pencil": 30}
        got = get_item_table(["eraser 50", "pencil 30"])
        got = self.assertEqual(want, got)

def main():
    itemCount, orderCount = map(int, input().split())
    lines = [input() for _ in range(itemCount)]
    itemTable = get_item_table(lines)

    orderLines = [input() for _ in range(orderCount)]

    for order in orderLines:
        if order in itemTable:
            print(itemTable[order])
        else:
            print(-1)

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
            unittest.main(argv=sys.argv[:1])
    else:
        main()

# テストを実行する場合
# % python3 template.py test
# .
# ----------------------------------------------------------------------
# Ran 1 test in 0.000s

# OK

# 通常の実行
# % python3 template.py

# Input:
# eraser 50
# pencil 30
# book 100
# book
# eraser
# pencil
# margaret

# Output:
# 100
# 50
# 30
# -1

# Q
# 下記の問題をプログラミングしてみよう！
# paiza 商店では N 個の商品が売られており、i 番目の商品の名前は A_i で、価格は B_i です。
# あなたは M 個の商品名が書かれたお買い物リスト S を持っています。リストに書かれているそれぞれの商品について、paiza 商店での価格を出力してください。リストには paiza 商店が扱っていない商品も書かれている可能性がありますが、その場合は価格の代わりに -1 を出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N M
# A_1 B_1
# A_2 B_2
# ...
# A_N B_N
# S_1
# S_2
# ...
# S_M

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# M 行出力してください。i 行目には、商品 S_i の価格 T_i を出力してください。paiza 商店に商品 S_i が売られていない場合は価格の代わりに -1 を出力してください。


# T_1
# T_2
# ...
# T_M


# 末尾に改行を入れ、余計な文字、空行を含んではいけません。
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ N ≦ 100
# ・ 1 ≦ M ≦ 100
# ・ A_i は英子文字からなる1文字以上10文字以下の文字列 (1 ≦ i ≦ N)
# ・ i ≠ j ならば A_i ≠ A_j
# ・ 1 ≦ B_i ≦ 100 (1 ≦ i ≦ N)
# ・ S_i は英子文字からなる1文字以上10文字以下の文字列 (1 ≦ i ≦ N)
# ・ i ≠ j ならば S_i ≠ S_j

# 入力例1
# 3 4
# eraser 50
# pencil 30
# book 100
# book
# eraser
# pencil
# margaret

# 出力例1
# 100
# 50
# 30
# -1
