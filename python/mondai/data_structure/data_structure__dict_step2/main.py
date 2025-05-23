import string

maxCount = int(input())
l = list(input())

alphabetDict = {char: 0 for char in string.ascii_lowercase}
for v in l:
    alphabetDict[v] += 1


print(*alphabetDict.values())

# Q
# 下記の問題をプログラミングしてみよう！
# 長さ N の文字列 S が与えられます。S に含まれている各文字の出現回数をそれぞれ求め、「a」の出現回数、「b」の出現回数、...、「z」の出現回数をこの順に半角スペース区切りで1行に出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N
# S

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# 「a」の出現回数 count_a、「b」の出現回数 count_b、...、「z」の出現回数 count_z をこの順に半角スペース区切りで1行に出力してください。


# count_a count_b ... count_z


# 末尾に改行を入れ、余計な文字、空行を含んではいけません。
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ N ≦ 100
# ・ S は英小文字「a」,「b」, ... ,「z」からなる長さ N の文字列

# 入力例1
# 13
# aaabbbccdddde

# 出力例1
# 3 3 2 4 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

# 入力例2
# 26
# ahgektndrmypqlfsjiouwzxcbv

# 出力例2
# 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
