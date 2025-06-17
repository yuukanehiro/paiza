
versions = [(1, 2, 3),(1, 2, 1),(1, 1, 4)]
versions.sort()
print(versions) # Output: [(1, 1, 4), (1, 2, 1), (1, 2, 3)]


countChars = [(2, "b"), (2, "a"), (1, "c")]
countChars.sort()
# 1つ目の要素を整数でソート
# 2つめの要素を辞書順でソート
print(countChars) # Output: [(1, 'c'), (2, 'a'), (2, 'b')]


countChars2 = [(1, "b"), (2, "a"), (3, "c"), (5, "a"), (5, "b")]
# 数値で降順, 辞書順
countChars2.sort(key=lambda x: (-x[0], x[1])) # -x[0]としてマイナスをかけて昇順ソートをすると結果として降順ソートになる
print(countChars2) # Output: [(5, 'a'), (5, 'b'), (3, 'c'), (2, 'a'), (1, 'b')]
