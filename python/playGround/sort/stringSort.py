# ----------------
# 文字列をソート
# ----------------
strings1 = ["orange", "apple", "banana", "grape"]
# 辞書順に昇順ソート
strings1.sort()
print(strings1) # Output: ['apple', 'banana', 'grape', 'orange']

# 辞書順に降順ソート
strings2 = ["orange", "apple", "banana", "grape"]
strings2.sort(key=lambda x: tuple(-ord(c) for c in x)) # ord(c) は各文字の文字コード（例：'a' → 97, 'z' → 122）
print(strings2) # Output: ['orange', 'grape', 'banana', 'apple']

# 任意の関数でソート
strings2 = ["orange", "apple", "banana", "grape"]
strings2.sort(key=lambda x: len(x))
print(strings2) # Output: ['apple', 'grape', 'orange', 'banana']

# lambda x: len(x)
# は下記と同じ
# def f(x):
#     return len(x)
