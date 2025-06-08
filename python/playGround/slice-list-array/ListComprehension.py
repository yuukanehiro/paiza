
squares1 = []
for i in range(10):
    squares1.append(i ** 2)
print(squares1) # Output: [0, 1, 4, 9, 16, 25, 36, 49, 64, 81]


# リスト内包表記（List comprehension）
# mapやfilterなどと同じような使い方ができる
squares2 = [i ** 2 for i in range(10)]
print(squares2) # Output: [0, 1, 4, 9, 16, 25, 36, 49, 64, 81]

squares3 = [i ** 2 for i in range(10) if i % 2 == 0]
print(squares3) # Output:[0, 4, 16, 36, 64]
