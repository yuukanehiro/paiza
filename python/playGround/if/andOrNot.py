
n, m = 5, 10

# 'and' 演算子の使用
if n > 0 and m > 0:
    print("Both n and m are positive.") # Output: Both n and m are positive.

# 'or' 演算子の使用
n = -5
if n < 0 or m > 0:
    print("Either n is negative or m is positive.") # Output: Either n is negative or m is positive.

# 'not' 演算子の使用
if not n > 0:
    print("n is not positive.") # Output: n is not positive.

# Pythonでは範囲を条件にする時に以下のようにできます
if 0 < n < 10:
    print("n is between 0 and 10")  # Output: n is between 0 and 10
