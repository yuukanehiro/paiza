def merge_sort(arr, depth=0):
    indent = "  " * depth  # 再帰の深さに応じてインデントを調整
    print(f"{indent}merge_sort({arr})")

    if len(arr) <= 1:
        print(f"{indent}=> return {arr}")
        return arr

    mid = len(arr) // 2
    left = merge_sort(arr[:mid], depth + 1)
    right = merge_sort(arr[mid:], depth + 1)

    merged = merge(left, right, depth + 1)
    print(f"{indent}=> merged {left} + {right} = {merged}")
    return merged

def merge(left, right, depth):
    indent = "  " * depth
    print(f"{indent}merge({left}, {right})")

    result = []
    i = j = 0

    # 両方のリストが残っている間、比較して小さい方を追加
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            print(f"{indent}→ {left[i]} <= {right[j]} → append {left[i]}")
            result.append(left[i])
            i += 1
        else:
            print(f"{indent}→ {left[i]} > {right[j]} → append {right[j]}")
            result.append(right[j])
            j += 1

    # 残った要素を全部追加
    if i < len(left):
        print(f"{indent}→ extend {left[i:]}")
    if j < len(right):
        print(f"{indent}→ extend {right[j:]}")
    result.extend(left[i:])
    result.extend(right[j:])
    return result

# テスト用のリスト
arr = [5, 2, 4, 1, 3]
sorted_arr = merge_sort(arr)
print("\n 最終結果:", sorted_arr)

# merge_sort([5, 2, 4, 1, 3])
#   merge_sort([5, 2])
#     merge_sort([5])
#     => return [5]
#     merge_sort([2])
#     => return [2]
#     merge([5], [2])
#     → 5 > 2 → append 2
#     → extend [5]
#   => merged [5] + [2] = [2, 5]
#   merge_sort([4, 1, 3])
#     merge_sort([4])
#     => return [4]
#     merge_sort([1, 3])
#       merge_sort([1])
#       => return [1]
#       merge_sort([3])
#       => return [3]
#       merge([1], [3])
#       → 1 <= 3 → append 1
#       → extend [3]
#     => merged [1] + [3] = [1, 3]
#     merge([4], [1, 3])
#     → 4 > 1 → append 1
#     → 4 > 3 → append 3
#     → extend [4]
#   => merged [4] + [1, 3] = [1, 3, 4]
#   merge([2, 5], [1, 3, 4])
#   → 2 > 1 → append 1
#   → 2 <= 3 → append 2
#   → 5 > 3 → append 3
#   → 5 > 4 → append 4
#   → extend [5]
# => merged [2, 5] + [1, 3, 4] = [1, 2, 3, 4, 5]

# 最終結果: [1, 2, 3, 4, 5]