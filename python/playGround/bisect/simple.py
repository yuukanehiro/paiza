from bisect import bisect_left, bisect_right, insort_left

nums1 = [1, 3, 3, 4, 4, 4, 4, 4, 19]
# 4以上の数が最初に現れるINDEX
print(bisect_left(nums1, 4)) # Output: 3
# 4より大きい数(19)の位置
print(bisect_right(nums1, 4)) # Output: 8
# ⭐️4の出現数
# bisect_right(nums, x) - bisect_left(nums, x)
print(bisect_right(nums1, 4) - bisect_left(nums1, 4)) # Output: 5


nums2 = [1, 3, 3, 5, 7, 9, 10, 13, 19]
print(bisect_left(nums2, 4)) # Output: 3
print(bisect_right(nums2, 4)) # Output: 3
# 出現数
print(bisect_right(nums2, 4) - bisect_left(nums2, 4)) # Output: 0 ... 4の出現数0


# ⭐️[lower, upper] の範囲にある個数
# bisect_right(nums, upper) - bisect_left(nums, lower)
print(bisect_right(nums2, 13) - bisect_left(nums2, 2)) # Output: 7

# ⭐️xを挿入すべき位置を取得しxを挿入する
print(bisect_right(nums2, 4)) # Output: 3
# ソート順を保った状態で4を挿入する
insort_left(nums2, 4)
print(nums2) # Output: [1, 3, 3, 4, 5, 7, 9, 10, 13, 19]
