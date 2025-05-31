import unittest
import sys
from collections import defaultdict



class TestHashTable(unittest.TestCase):
    def test_get_hash_table(self):
         got = get_hash_table(["hij", "abc", "abc", "bcd"])
         want = {"abc": 2, "bcd": 1, "hij": 1}
         self.assertEqual(got, want)

def get_hash_table(lines):
    hashTable = defaultdict(int)
    for keyString in lines:
        hashTable[keyString] += 1

    return hashTable

def main():
    max_line_count = int(input())
    lines = [input() for _ in range(max_line_count)]
    hashTable = get_hash_table(lines)
    for k in sorted(hashTable):
        print(f"{k} {hashTable[k]}")

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
            unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 文字列が N 個与えられます。各文字列の出現回数を文字列の辞書順に出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N
# S_1
# S_2
# ...
# S_N

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# 文字列 S とその出現回数 A を、文字列 S の辞書順に、改行区切りで出力してください。


# S A
# ...


# 末尾に改行を入れ、余計な文字、空行を含んではいけません。
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ N ≦ 100
# ・ S_i は英小文字「a」,「b」, ... ,「z」からなる1文字以上3文字以下の文字列 (1 ≦ i ≦ N)

# 入力例1
# 5
# bcd
# abc
# bcd
# bcd
# bcd

# 出力例1
# abc 1
# bcd 4

# 入力例2
# 7
# p
# p
# pa
# pa
# p
# pai
# pai

# 出力例2
# p 3
# pa 2
# pai 2
