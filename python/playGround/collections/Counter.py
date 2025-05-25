from collections import Counter

# リスト内の要素の出現回数をカウント
fruits = ["apple", "banana", "apple", "orange", "banana", "apple"]
fruitCounter = Counter(fruits)
print(fruitCounter) # Output: Counter({'apple': 3, 'banana': 2, 'orange': 1})

print(fruitCounter["apple"]) # Output: 3
print(fruitCounter["cherry"]) # Output: 0 // 存在しない要素は0を返却

# 出現頻度が多い要素を取得
print(fruitCounter.most_common(1)) # Output: [('apple', 3)]
print(fruitCounter.most_common(2)) # Output: [('apple', 3), ('banana', 2)]
print(fruitCounter.most_common(3)) # Output: [('apple', 3), ('banana', 2), ('orange', 1)]
