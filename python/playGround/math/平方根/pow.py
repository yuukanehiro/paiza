
from decimal import Decimal, getcontext
import math

getcontext().prec = 30

x = Decimal("10000000000000001")

sqrt_math = Decimal(math.sqrt(float(x)))
sqrt_pow = x ** Decimal("0.5")

print(f"math.sqrt : {sqrt_math}")
print(f"pow       : {sqrt_pow}")
print(f"diff      : {sqrt_math - sqrt_pow}")


# % python3 pow.py
# math.sqrt : 100000000
# pow       : 100000000.000000005000000000000
# diff      : -5.000000000000E-9
