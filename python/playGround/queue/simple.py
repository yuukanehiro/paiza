from collections import deque

queue = deque()
queue.append(1)
queue.append(2)
print(queue) # Output: deque([1, 2])

# popleft() ... キューの先頭から取り出す
queue.popleft()
print(queue) # Output: deque([2])
