import heapq

minHeap = []
heapq.heappush(minHeap, 3)
heapq.heappush(minHeap, 2)
heapq.heappush(minHeap, 4)

# 最小値はindex0
print("start minHeap")
print(minHeap[0]) # Output: 2

while minHeap:
    print(heapq.heappop(minHeap))
# Output:
# 2
# 3
# 4


maxHeap = []
heapq.heappush(maxHeap, -3)
heapq.heappush(maxHeap, -2)
heapq.heappush(maxHeap, -4)

print("start maxHeap")
print(-maxHeap[0]) # Output: 4

while maxHeap:
    print(-heapq.heappop(maxHeap))
# Output:
# 4
# 3
# 2
