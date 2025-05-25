
hashSet = set()
hashSet.add(1)
hashSet.add(2)
print(hashSet) # Output: {1, 2}
print(len(hashSet)) # Output: 2

print(1 in hashSet) # Output: True
print(2 in hashSet) # Output: True
print(3 in hashSet) # Output: False

# 要素の削除
hashSet.remove(2)
print(2 in hashSet) # Output: False

# 存在しない要素を削除すると例外エラー
# hashSet.remove(3) # Error!

# 内包表記
hashSet2 = { i for i in range(3) }
print(hashSet2) # Output: {0, 1, 2}

hashSet2.add(1)
print(hashSet2) # Output: {0, 1, 2} ... 変化しないことを確認
hashSet2.add(3)
print(hashSet2)  # Output: {0, 1, 2, 3} ... 存在しない要素は追加される
