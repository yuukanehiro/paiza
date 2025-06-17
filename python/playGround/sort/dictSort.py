d = {
    "b": 1,
    "a": 2,
    "c": 3,
    "a2": 5,
    "b2": 5
}

# 値で降順ソート、同じ値の場合はキーを辞書降順ソート
# sorted()の返却値は[]なのでdict()で戻す
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
