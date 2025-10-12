package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Go etree XML処理サンプル")
	fmt.Println("=====================")
	fmt.Println()
	fmt.Println("実行したいサンプルを選択してください:")
	fmt.Println("1. 基本的なXML作成")
	fmt.Println("2. XMLファイルの読み書き")
	fmt.Println("3. XML要素の検索")
	fmt.Println("4. XML要素の修正")
	fmt.Println("5. XML属性の操作")
	fmt.Println("6. XML名前空間の処理")
	fmt.Println("0. 終了")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("選択 (0-6): ")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("\n=== 1. 基本的なXML作成 ===")
			main_create_xml()
		case "2":
			fmt.Println("\n=== 2. XMLファイルの読み書き ===")
			main_file_io()
		case "3":
			fmt.Println("\n=== 3. XML要素の検索 ===")
			main_search()
		case "4":
			fmt.Println("\n=== 4. XML要素の修正 ===")
			main_modify()
		case "5":
			fmt.Println("\n=== 5. XML属性の操作 ===")
			main_attributes()
		case "6":
			fmt.Println("\n=== 6. XML名前空間の処理 ===")
			main_namespaces()
		case "0":
			fmt.Println("プログラムを終了します。")
			return
		default:
			fmt.Println("無効な選択です。0-6の数字を入力してください。")
			continue
		}

		fmt.Println("\n" + strings.Repeat("-", 50))
		fmt.Println("メニューに戻ります...")
		fmt.Println()
	}
}