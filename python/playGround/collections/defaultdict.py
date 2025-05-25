from collections import defaultdict

hashTable = defaultdict(int)
hashTable["apple"] += 1  # "apple"がなければ、デフォルトの0に1を足して1になる
hashTable["banana"] += 1
hashTable["apple"] += 1  # "apple"の値が更新され、2になる

print(hashTable["apple"]) # Output: 2
