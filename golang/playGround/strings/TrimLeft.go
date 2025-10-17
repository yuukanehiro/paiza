package main

import (
	"fmt"
	"strings"
)

func main() {
	// 基本的な使い方
	fmt.Println("=== 基本的な使い方 ===")
	str := "###Hello, World!###"
	trimmed := strings.TrimLeft(str, "#")
	fmt.Printf("元の文字列: '%s'\n", str)
	fmt.Printf("TrimLeft後: '%s'\n", trimmed)

	// 空白文字の削除
	fmt.Println("\n=== 空白文字の削除 ===")
	spaceStr := "   Hello, World!   "
	trimmedSpace := strings.TrimLeft(spaceStr, " ")
	fmt.Printf("元の文字列: '%s'\n", spaceStr)
	fmt.Printf("TrimLeft後: '%s'\n", trimmedSpace)

	// 複数文字の削除
	fmt.Println("\n=== 複数文字の削除 ===")
	multiChar := "aaabbbcccHello"
	trimmedMulti := strings.TrimLeft(multiChar, "abc")
	fmt.Printf("元の文字列: '%s'\n", multiChar)
	fmt.Printf("TrimLeft('abc')後: '%s'\n", trimmedMulti)

	// 数字の削除
	fmt.Println("\n=== 数字の削除 ===")
	numStr := "123456Hello789"
	trimmedNum := strings.TrimLeft(numStr, "0123456789")
	fmt.Printf("元の文字列: '%s'\n", numStr)
	fmt.Printf("TrimLeft(数字)後: '%s'\n", trimmedNum)

	// 左側のみ削除（右側は残る）
	fmt.Println("\n=== 左側のみ削除（右側は残る） ===")
	leftRight := "***Hello***"
	trimmedLeft := strings.TrimLeft(leftRight, "*")
	fmt.Printf("元の文字列: '%s'\n", leftRight)
	fmt.Printf("TrimLeft後: '%s' (右側の*は残る)\n", trimmedLeft)

	// 該当する文字がない場合
	fmt.Println("\n=== 該当する文字がない場合 ===")
	noMatch := "Hello, World!"
	trimmedNoMatch := strings.TrimLeft(noMatch, "#")
	fmt.Printf("元の文字列: '%s'\n", noMatch)
	fmt.Printf("TrimLeft('#')後: '%s' (変化なし)\n", trimmedNoMatch)

	// 空文字列
	fmt.Println("\n=== 空文字列 ===")
	empty := ""
	trimmedEmpty := strings.TrimLeft(empty, "#")
	fmt.Printf("元の文字列: '%s'\n", empty)
	fmt.Printf("TrimLeft後: '%s'\n", trimmedEmpty)

	// TrimLeft vs TrimPrefix の違い
	fmt.Println("\n=== TrimLeft vs TrimPrefix の違い ===")
	str2 := "ababHello"
	trimLeft := strings.TrimLeft(str2, "ab")
	trimPrefix := strings.TrimPrefix(str2, "ab")
	fmt.Printf("元の文字列: '%s'\n", str2)
	fmt.Printf("TrimLeft('ab'): '%s' (左側のa,bを全て削除)\n", trimLeft)
	fmt.Printf("TrimPrefix('ab'): '%s' (先頭の'ab'のみ削除)\n", trimPrefix)

	// 実用例: URLのスキーマ削除
	fmt.Println("\n=== 実用例: プロトコルの削除 ===")
	urls := []string{
		"https://example.com",
		"http://example.com",
		"ftp://example.com",
	}
	for _, url := range urls {
		cleaned := strings.TrimLeft(url, "htps:/")
		fmt.Printf("元: %s -> 削除後: %s\n", url, cleaned)
	}

	// 実用例: CSVデータの前処理
	fmt.Println("\n=== 実用例: CSVデータの前処理 ===")
	csvData := "   ,,,John,Doe,30"
	cleaned := strings.TrimLeft(csvData, " ,")
	fmt.Printf("元のデータ: '%s'\n", csvData)
	fmt.Printf("クリーンアップ後: '%s'\n", cleaned)

	// 実用例: ファイルパスの処理
	fmt.Println("\n=== 実用例: ファイルパスの処理 ===")
	path := "///path/to/file.txt"
	cleanPath := strings.TrimLeft(path, "/")
	fmt.Printf("元のパス: '%s'\n", path)
	fmt.Printf("スラッシュ削除後: '%s'\n", cleanPath)
}

// === 基本的な使い方 ===
// 元の文字列: '###Hello, World!###'
// TrimLeft後: 'Hello, World!###'

// === 空白文字の削除 ===
// 元の文字列: '   Hello, World!   '
// TrimLeft後: 'Hello, World!   '

// === 複数文字の削除 ===
// 元の文字列: 'aaabbbcccHello'
// TrimLeft('abc')後: 'Hello'

// === 数字の削除 ===
// 元の文字列: '123456Hello789'
// TrimLeft(数字)後: 'Hello789'

// === 左側のみ削除（右側は残る） ===
// 元の文字列: '***Hello***'
// TrimLeft後: 'Hello***' (右側の*は残る)

// === 該当する文字がない場合 ===
// 元の文字列: 'Hello, World!'
// TrimLeft('#')後: 'Hello, World!' (変化なし)

// === 空文字列 ===
// 元の文字列: ''
// TrimLeft後: ''

// === TrimLeft vs TrimPrefix の違い ===
// 元の文字列: 'ababHello'
// TrimLeft('ab'): 'Hello' (左側のa,bを全て削除)
// TrimPrefix('ab'): 'abHello' (先頭の'ab'のみ削除)

// === 実用例: プロトコルの削除 ===
// 元: https://example.com -> 削除後: example.com
// 元: http://example.com -> 削除後: example.com
// 元: ftp://example.com -> 削除後: ftp://example.com

// === 実用例: CSVデータの前処理 ===
// 元のデータ: '   ,,,John,Doe,30'
// クリーンアップ後: 'John,Doe,30'

// === 実用例: ファイルパスの処理 ===
// 元のパス: '///path/to/file.txt'
// スラッシュ削除後: 'path/to/file.txt'