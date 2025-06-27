

# =========
# iterator
# 1つずつ要素数を取り出すことが出来る
# =========
nums = [10, 20, 30]
it = iter(nums)  # イテレーターを取得

print(next(it)) # Output: 10
print(next(it)) # Output: 20
print(next(it)) # Output: 30
# 要素数を超えて取得しようとすると StopIteration エラー
print(next(it)) # Output: Traceback (most recent call last):
#   File "/Users/kanehiroyuu/Documents/GitHub/paiza/python/playGround/iterator/main.py", line 9, in <module>
#     print(next(it))
# StopIteration
