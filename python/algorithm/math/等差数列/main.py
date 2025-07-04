def main():
    x, d = map(int, input().split())
    query_count = int(input())
    queries = [int(input().strip()) for _ in range(query_count)]

    for k in queries:
        print(x + (k - 1) * d) # x ... 初項, d ... 交差, k ... 指定する項