
versions = [(1, 2, 3),(1, 2, 1),(1, 1, 4)]
versions.sort()
print(versions) # Output: [(1, 1, 4), (1, 2, 1), (1, 2, 3)]


countChars = [(2, "b"), (2, "a"), (1, "c")]
countChars.sort()
# 1つ目の要素を整数でソート
# 2つめの要素を辞書順でソート
print(countChars) # Output: [(1, 'c'), (2, 'a'), (2, 'b')]


countChars2 = [(1, "b"), (2, "a"), (3, "c"), (5, "a"), (5, "b")]
# 数値で降順, 同じ数値の場合は辞書順に昇順
countChars2.sort(key=lambda x: (-x[0], x[1])) # -x[0]としてマイナスをかけて昇順ソートをすると結果として降順ソートになる
print(countChars2) # Output: [(5, 'a'), (5, 'b'), (3, 'c'), (2, 'a'), (1, 'b')]

# 数値で降順, 同じ数値の場合は辞書順に降順
countChars3 = [(1, "b"), (2, "a"), (3, "c"), (5, "a"), (5, "b")]
countChars3.sort(key=lambda x: (-x[0], tuple(-ord(c) for c in x[1]))) # ord(c) は各文字の文字コード（例：'a' → 97, 'z' → 122）
print(countChars3) # Output: [(5, 'b'), (5, 'a'), (3, 'c'), (2, 'a'), (1, 'b')]
