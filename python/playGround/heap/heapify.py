import heapq

# リストをヒープに変換
arr = [2, 3, 1, 6, 4, 8]
heapq.heapify(arr)

while arr:
    print(heapq.heappop(arr))

# Output
# 1
# 2
# 3
# 4
# 6
# 8
