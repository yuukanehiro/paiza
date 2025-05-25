# ----------------
# 文字列をソート
# ----------------
strings1 = ["orange", "apple", "banana", "grape"]
# 辞書順にソート
strings1.sort()
print(strings1) # Output: ['apple', 'banana', 'grape', 'orange']

# 任意の関数でソート
strings2 = ["orange", "apple", "banana", "grape"]
strings2.sort(key=lambda x: len(x))
print(strings2) # Output: ['apple', 'grape', 'orange', 'banana']

# lambda x: len(x)
# は下記と同じ
# def f(x):
#     return len(x)
