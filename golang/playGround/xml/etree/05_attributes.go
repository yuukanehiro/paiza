package main

import (
	"fmt"

	"github.com/beevik/etree"
)

func main_attributes() {
	fmt.Println("5. XML属性操作のサンプル")
	fmt.Println("========================")

	// XMLファイルを読み込み
	doc := etree.NewDocument()
	err := doc.ReadFromFile("library.xml")
	if err != nil {
		fmt.Printf("ファイル読み込みエラー: %v\n", err)
		fmt.Println("library.xmlファイルが必要です。先に02_file_io.goを実行してください。")
		return
	}

	fmt.Println("=== 属性の読み取り ===")

	// 最初の書籍の属性を確認
	book1 := doc.FindElement("//book[@id='1']")
	if book1 != nil {
		fmt.Println("\nID=1の書籍の属性:")

		// 全属性を表示
		fmt.Println("  全ての属性:")
		for _, attr := range book1.Attr {
			fmt.Printf("    %s = \"%s\"\n", attr.Key, attr.Value)
		}

		// 特定の属性値を取得
		id := book1.SelectAttrValue("id", "不明")
		category := book1.SelectAttrValue("category", "不明")
		fmt.Printf("  特定属性: ID=%s, Category=%s\n", id, category)
	}

	fmt.Println("\n=== 属性の追加 ===")

	// 新しい属性を追加
	if book1 != nil {
		book1.CreateAttr("rating", "5")
		book1.CreateAttr("publisher", "技術出版社")
		book1.CreateAttr("available", "true")
		fmt.Println("新しい属性を追加: rating, publisher, available")

		fmt.Println("\n追加後の属性:")
		for _, attr := range book1.Attr {
			fmt.Printf("  %s = \"%s\"\n", attr.Key, attr.Value)
		}
	}

	fmt.Println("\n=== 属性の修正 ===")

	// 属性値を変更
	if book1 != nil {
		// rating属性を変更
		ratingAttr := book1.SelectAttr("rating")
		if ratingAttr != nil {
			oldRating := ratingAttr.Value
			ratingAttr.Value = "4"
			fmt.Printf("rating属性を %s から 4 に変更\n", oldRating)
		}

		// available属性を変更
		availableAttr := book1.SelectAttr("available")
		if availableAttr != nil {
			availableAttr.Value = "false"
			fmt.Println("available属性をfalseに変更")
		}
	}

	fmt.Println("\n=== 属性の削除 ===")

	// 特定の属性を削除
	if book1 != nil {
		fmt.Println("publisher属性を削除前:")
		publisher := book1.SelectAttrValue("publisher", "")
		fmt.Printf("  publisher = \"%s\"\n", publisher)

		book1.RemoveAttr("publisher")
		fmt.Println("publisher属性を削除")

		// 削除後の確認
		publisherAfter := book1.SelectAttrValue("publisher", "属性なし")
		fmt.Printf("削除後: publisher = \"%s\"\n", publisherAfter)
	}

	fmt.Println("\n=== 全書籍の属性一覧 ===")

	// 全書籍の属性を表示
	books := doc.FindElements("//book")
	for i, book := range books {
		id := book.SelectAttrValue("id", "")
		titleElem := book.SelectElement("title")
		title := ""
		if titleElem != nil {
			title = titleElem.Text()
		}

		fmt.Printf("\n書籍 %d: %s (ID:%s)\n", i+1, title, id)
		fmt.Println("  属性:")
		for _, attr := range book.Attr {
			fmt.Printf("    %s = \"%s\"\n", attr.Key, attr.Value)
		}
	}

	fmt.Println("\n=== 属性による検索 ===")

	// 特定の属性値を持つ要素を検索
	programBooks := doc.FindElements("//book[@category='programming']")
	fmt.Printf("\nプログラミングカテゴリの書籍数: %d\n", len(programBooks))

	ratedBooks := doc.FindElements("//book[@rating]")
	fmt.Printf("評価属性を持つ書籍数: %d\n", len(ratedBooks))

	// ファイルに保存
	fmt.Println("\n=== ファイルに保存 ===")
	doc.Indent(2)
	err = doc.WriteToFile("attributes_library.xml")
	if err != nil {
		fmt.Printf("ファイル保存エラー: %v\n", err)
	} else {
		fmt.Println("属性操作したXMLをattributes_library.xmlに保存")
	}

	fmt.Println("\n属性操作完了")
}