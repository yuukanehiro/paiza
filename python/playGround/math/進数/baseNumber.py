def convert_decimal(number: int):
    binary = bin(number)      # 2進数に変換（プレフィックス '0b' 付き）
    hexadecimal = hex(number) # 16進数に変換（プレフィックス '0x' 付き）

    print(f"10進数: {number}")
    print(f"2進数 : {binary}（プレフィックス付き）")
    print(f"2進数 : {binary[2:]}（プレフィックスなし）")
    print(f"16進数: {hexadecimal}（プレフィックス付き）")
    print(f"16進数: {hexadecimal[2:]}（プレフィックスなし）")


convert_decimal(255)
# Output:
# 10進数: 255
# 2進数 : 0b11111111（プレフィックス付き）
# 2進数 : 11111111（プレフィックスなし）
# 16進数: 0xff（プレフィックス付き）
# 16進数: ff（プレフィックスなし）
