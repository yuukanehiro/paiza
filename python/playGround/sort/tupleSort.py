
versions = [(1, 2, 3),(1, 2, 1),(1, 1, 4)]
versions.sort()
print(versions) # Output: [(1, 1, 4), (1, 2, 1), (1, 2, 3)]


countChars = [(2, "b"), (2, "a"), (1, "c")]
countChars.sort()
# 1つ目の要素を整数でソート
# 2つめの要素を辞書順でソート
print(countChars) # Output: [(1, 'c'), (2, 'a'), (2, 'b')]
