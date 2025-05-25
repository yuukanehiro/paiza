
s = "hello"

# 指定したindexの文字を取得
print(s[1]) # Output: e
print(s[2:]) # Output: llo
print(s[1:3]) # Output: el
print(s[:3]) # Output: hel
# step1
print(s[::1]) # Output: hello
# step2
print(s[::2]) # Output: hlo
# 逆順
print(s[::-1]) # Output: olleh

# 文字数取得
print(len(s)) # Output: 5

# 文字列結合
l = s + " " + "Yuu"
print(l) # Output: hello Yuu

# 変数展開
l2 = f"{s} world"
print(l2)  # Output: hello world

# 文字列を数値に変換できる
print(int("123") + int("123")) # Output: 246

# 数値を文字列に変換できる
print(str(123) + str(123)) # Output: 123123
