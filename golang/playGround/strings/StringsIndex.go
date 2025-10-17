package main

import (
	"fmt"
	"strings"
)

func main() {
	// 基本的な使い方
	fmt.Println("=== 基本的な使い方 ===")
	str := "Hello, World!"
	index := strings.Index(str, "World")
	fmt.Printf("'%s' の中の 'World' の位置: %d\n", str, index)

	// 見つからない場合
	fmt.Println("\n=== 見つからない場合 ===")
	notFound := strings.Index(str, "Golang")
	fmt.Printf("'%s' の中の 'Golang' の位置: %d (見つからない場合は -1)\n", str, notFound)

	// 最初の一致を返す
	fmt.Println("\n=== 最初の一致を返す ===")
	repeated := "banana"
	firstA := strings.Index(repeated, "a")
	fmt.Printf("'%s' の中の最初の 'a' の位置: %d\n", repeated, firstA)

	// 複数文字の検索
	fmt.Println("\n=== 複数文字の検索 ===")
	sentence := "Go is a programming language"
	progIndex := strings.Index(sentence, "programming")
	fmt.Printf("'%s' の中の 'programming' の位置: %d\n", sentence, progIndex)

	// 日本語での使用
	fmt.Println("\n=== 日本語での使用 ===")
	jpStr := "こんにちは、世界！"
	worldIndex := strings.Index(jpStr, "世界")
	fmt.Printf("'%s' の中の '世界' の位置: %d (バイト単位)\n", jpStr, worldIndex)

	// 空文字列の検索
	fmt.Println("\n=== 空文字列の検索 ===")
	emptyIndex := strings.Index(str, "")
	fmt.Printf("空文字列の位置: %d (常に0を返す)\n", emptyIndex)

	// 実用例: 部分文字列の抽出
	fmt.Println("\n=== 実用例: 部分文字列の抽出 ===")
	email := "user@example.com"
	atIndex := strings.Index(email, "@")
	if atIndex != -1 {
		username := email[:atIndex]
		domain := email[atIndex+1:]
		fmt.Printf("Email: %s\n", email)
		fmt.Printf("ユーザー名: %s\n", username)
		fmt.Printf("ドメイン: %s\n", domain)
	}

	// 実用例: 文字列の存在チェック
	fmt.Println("\n=== 実用例: 文字列の存在チェック ===")
	text := "This is a sample text for demonstration"
	searchWord := "sample"
	if strings.Index(text, searchWord) != -1 {
		fmt.Printf("'%s' は '%s' に含まれています\n", searchWord, text)
	} else {
		fmt.Printf("'%s' は '%s' に含まれていません\n", searchWord, text)
	}
}


// % go run main.go  
// === 基本的な使い方 ===
// 'Hello, World!' の中の 'World' の位置: 7

// === 見つからない場合 ===
// 'Hello, World!' の中の 'Golang' の位置: -1 (見つからない場合は -1)

// === 最初の一致を返す ===
// 'banana' の中の最初の 'a' の位置: 1

// === 複数文字の検索 ===
// 'Go is a programming language' の中の 'programming' の位置: 8

// === 日本語での使用 ===
// 'こんにちは、世界！' の中の '世界' の位置: 18 (バイト単位)

// === 空文字列の検索 ===
// 空文字列の位置: 0 (常に0を返す)

// === 実用例: 部分文字列の抽出 ===
// Email: user@example.com
// ユーザー名: user
// ドメイン: example.com

// === 実用例: 文字列の存在チェック ===
// 'sample' は 'This is a sample text for demonstration' に含まれています