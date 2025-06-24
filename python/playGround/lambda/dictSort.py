d = {
    "b": 1,
    "a": 2,
    "c": 3,
    "a2": 5,
    "b2": 5
}

# 値で降順ソート、同じ値の場合はキーを辞書降順ソート
# sorted()の返却値は[]なのでdict()で戻す
# lambda x: (-x[1], tuple(-ord(c) for c in x[0])) は次のような順で並び替えます：
# ・値が大きい順（降順）
# ・値が同じならキーを辞書順で降順
# 　・ord("a") == 97, ord("b") == 98, ord("2") == 50
# 　・"a"     → (-97,)
# 　・"a2"    → (-97, -50)
sorted_dict = dict(sorted(
    d.items(),
    key=lambda x: (-x[1], tuple(-ord(c) for c in x[0]))
))

# 出力
for k, v in sorted_dict.items():
    print(f"{k}: {v}")
# Output:
# b2: 5
# a2: 5
# c: 3
# a: 2
# b: 1
