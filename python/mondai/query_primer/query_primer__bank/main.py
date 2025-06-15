import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_result_bank(items: List[str], queries: List[str]) -> List[str]:
    result = []
    bank = defaultdict(dict)
    for line in items:
        name, password, amount = line.split()
        bank[name][password] = int(amount)

    for q in queries:
        name, password, pull = q.split()
        if name in bank and password in bank[name]:
            bank[name][password] -= int(pull)

    for line in items:
        name, password, amount = line.split()
        if name in bank and password in bank[name]:
            result.append(f"{name} {bank[name][password]}")

    return result


class TestItemPriceMap(unittest.TestCase):
    def test_get_result_bank(self):
        expected = [
            "A 9800",
            "B 22456",
            "C 8765"
        ]
        input_items = [
            "A 0000 10000",
            "B 1234 23456",
            "C 5678 98765"
        ]
        input_queries = [
            "A 0101 1000",
            "B 1234 1000",
            "C 5678 90000",
            "A 0000 200",
            "B 1233 10000"
        ]
        actual = get_result_bank(input_items, input_queries)
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    items = [input().strip() for _ in range(item_count)]
    queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    for v in get_result_bank(items, queries):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 銀行 Python3編（paizaランク B 相当）
# 問題にチャレンジして、ユーザー同士で解答を教え合ったり、コードを公開してみよう！

# シェア用URL:
# https://paiza.jp/works/mondai/query_primer/query_primer__bank
# 問題文のURLをコピーする
#  下記の問題をプログラミングしてみよう！
# 2xxx年に paiza が設立した paiza 中央銀行に勤務する pai沢直樹は、故障した ATM の対応として、お金を引き出したい会社と電話をして、会社名が銀行の名簿に登録されており、かつ、会社側が会社の口座の暗証番号を正しく言えた場合にのみ現金を支払い、それを記帳するという業務を任されていました。

# 銀行に登録されている会社名とその口座の暗証番号と残高についての情報、また、直樹の電話の情報が与えられるので、全ての取引が終了した後の全ての会社の残高を出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# C_1 P_1 D_1
# ...
# C_N P_N D_N
# G_1 M_1 W_1
# ...
# G_K M_K W_K


# ・1 行目では、銀行に登録されている会社の数 N と行った取引の数 K が与えられます。
# ・続く N 行のうち、i 行目では、i 番目に登録されている会社名 C_i とその口座の暗証番号 P_i と残高 D_i が与えられます。
# ・続く K 行のうち、i 行目では、i 回目の取引を行おうとした会社の名前 G_i と、その人が言った暗証番号 M_i , 引出そうとした金額 W_i が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# C_1 D_1
# ...
# C_N D_N


# ・ i 行目に、i 番目に登録されている会社名 C_i とその会社の残高 D_i を以上の形式で出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・1 ≦ N ≦ 100,000
# ・1 ≦ K ≦ 100,000
# ・C_i は 20 文字以下の文字列 (1 ≦ i ≦ N)
# ・P_i , M_j は 0000 〜 9999 のいずれかの数値列 (1 ≦ i ≦ N , 1 ≦ j ≦ K)
# ・1 ≦ D_i ≦ 1,000,000,000 (1 ≦ i ≦ N)
# ・G_i は C_j のいずれか (1 ≦ i ≦ K , 1 ≦ j ≦ N)
# ・1 ≦ W_i ≦ 1,000,000 (1 ≦ i ≦ K)
# ・C_i ≠ C_j (i ≠ j)
# ・取引の結果、全ての会社の残高が負にならないことが保証されている。

# 入力例1
# 3 5
# A 0000 10000
# B 1234 23456
# C 5678 98765
# A 0101 1000
# B 1234 1000
# C 5678 90000
# A 0000 200
# B 1233 10000

# 出力例1
# A 9800
# B 22456
# C 8765

# 入力例2
# 4 8
# i 1353 758385
# my 9486 46446
# me 3785 38575
# mine 3573 92474
# i 3785 38753
# i 7536 8674
# my 2472 973
# my 2984 385
# me 7537 4757
# me 3785 3757
# mine 3757 3857
# mine 3573 3858

# 出力例2
# i 758385
# my 46446
# me 34818
# mine 88616
