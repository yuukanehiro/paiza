import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# を返却
# ------------------------
def get_result_dp(n: int) -> List[int]:
    # 0段の時は登りきっていると考えて1通りと考える
    dp = [1] * (n + 1)

    for i in range(2, n + 1):
        dp[i] = dp[i - 1] + dp[i - 2]

    return dp

# 4段目にいく方法
# 1 → 1 → 1 → 1
# 1 → 2 → 1
# 2 → 1 → 1
# 1 → 1 → 2
# 2 → 2
# → 5通り

# 3段目にいく方法
# 1 → 1 → 1
# 1 → 2
# 2 → 1
# → 3通り

# 2段目にいく方法
# 1 → 1
# 2
# → 2通り

# 3段目から4段目にいく方法
# 1 → 1 → 1 → 1
# 1 → 2 → 1
# 2 → 1 → 1
# → 3通り

# 2段目から3段目で止まらずに4段目にいく方法
# 1 → 1 → 2
# 2 → 2
# ❌ 2 → 1 → 1 ... 3段目に着地してしまっているので重複する
# → 2通り

# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestItemPriceMap(unittest.TestCase):
    def test_get_result_dp_3(self):
        expected = [1, 1, 2, 3]
        actual = get_result_dp(3)
        self.assertEqual(expected, actual)
    def test_get_result_dp_4(self):
        expected = [1, 1, 2, 3, 5]
        actual = get_result_dp(4)
        self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    # item_count, query_count = map(int, input().split())
    n = int(input()) # 1つのintの場合
    # inputList = list(map(int, input().split()))
    # item_count, query_count = inputList[0], inputList[1]
    # items = [input().strip() for _ in range(item_count)]
    # 2次元配列
    # items = [list(map(int, input().split())) for _ in range(item_count)]
    # 1 - indexed
    # items = [0] + [int(input().strip()) for _ in range(item_count)]
    # queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    dp = (get_result_dp(n))
    print(dp[n])


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# https://paiza.jp/works/mondai/dp_primer/dp_primer_stairs_step0/edit?language_uid=python3
# 下記の問題をプログラミングしてみよう！
# 整数 n が与えられます。
# 階段を上るのに、1 歩で 1 段または 2 段を上ることができるとき、n 段の階段を上る方法は何通りあるでしょうか。

# (ヒント)
# これまでは問題文中に具体的な漸化式が書かれていましたが、この問題にはありません。自分で漸化式を立てる必要があります。

# 部分問題として、1 ~ n-1 段の階段を上る方法が何通りあるか、という問題を考えてみましょう。この部分問題の答えが分かっているとして、n 段の階段を上る方法が何通りあるかを考えてみましょう。n 段目に到達するには、n-1 段目から1段上る方法と、n-2 段目から2段上る方法の2種類が考えられます。dp[n] を n 段の階段を上る方法の数とすれば、この関係は dp[n] = dp[n-1] + dp[n-2] で表すことが出来ます。よって、0段の階段を上る方法が1通り (何もしない) であることを踏まえると、以下のようにして答えを求めることが出来ます。

# dp[0] <- 1

# for i = 1 to n
#     dp[i] <- 0
#     if i >= 1 then
#         dp[i] <- dp[i] + dp[i-1]    // i-1 段目から1段上って i 段へ到達
#     if i >= 2 then
#         dp[i] <- dp[i] + dp[i-2]    // i-2 段目から2段上って i 段へ到達

# print dp[n]
# このような場合分けをすると上で考察した漸化式を満たす配列が実現できます (ピンとこなければ、i に具体的な値を入れて dp[i] がどのように計算されるのか、その処理を追ってみましょう) 。この場合分けは今のところ冗長に見えますが、次の問題を解くときに活きてきます。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# n

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# n 段の階段を上る方法の数を1行に出力してください。

# また、末尾に改行を入れ、余計な文字、空行を含んではいけません。

# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ n ≦ 40

# 入力例1
# 3

# 出力例1
# 3
