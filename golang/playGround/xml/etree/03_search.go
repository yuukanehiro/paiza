package main

import (
	"fmt"

	"github.com/beevik/etree"
)

func main_search() {
	fmt.Println("3. XML要素検索のサンプル")
	fmt.Println("========================")

	// サンプルXMLファイルを読み込み
	doc := etree.NewDocument()
	err := doc.ReadFromFile("library.xml")
	if err != nil {
		fmt.Printf("ファイル読み込みエラー: %v\n", err)
		fmt.Println("library.xmlファイルが必要です。先に02_file_io.goを実行してください。")
		return
	}

	fmt.Println("\n=== 基本的な検索 ===")

	// 全ての書籍を検索
	books := doc.FindElements("//book")
	fmt.Printf("全書籍数: %d冊\n", len(books))

	// 各書籍の情報を表示
	for i, book := range books {
		id := book.SelectAttrValue("id", "")
		category := book.SelectAttrValue("category", "")

		titleElem := book.SelectElement("title")
		title := ""
		lang := ""
		if titleElem != nil {
			title = titleElem.Text()
			lang = titleElem.SelectAttrValue("lang", "")
		}

		authorElem := book.SelectElement("author")
		author := ""
		if authorElem != nil {
			author = authorElem.Text()
		}

		fmt.Printf("\n書籍 %d:\n", i+1)
		fmt.Printf("  ID: %s\n", id)
		fmt.Printf("  カテゴリ: %s\n", category)
		fmt.Printf("  タイトル: %s (%s)\n", title, lang)
		fmt.Printf("  著者: %s\n", author)
	}

	fmt.Println("\n=== 条件付き検索 ===")

	// 日本語の書籍のみ
	japaneseTitles := doc.FindElements("//title[@lang='ja']")
	fmt.Println("\n日本語の書籍:")
	for _, title := range japaneseTitles {
		fmt.Printf("  - %s\n", title.Text())
	}

	// プログラミングカテゴリの書籍
	programmingBooks := doc.FindElements("//book[@category='programming']")
	fmt.Println("\nプログラミングカテゴリの書籍:")
	for _, book := range programmingBooks {
		titleElem := book.SelectElement("title")
		authorElem := book.SelectElement("author")
		if titleElem != nil && authorElem != nil {
			fmt.Printf("  - %s (著者: %s)\n", titleElem.Text(), authorElem.Text())
		}
	}

	// 特定のIDの書籍
	book1 := doc.FindElement("//book[@id='1']")
	if book1 != nil {
		fmt.Println("\nID=1の書籍:")
		titleElem := book1.SelectElement("title")
		if titleElem != nil {
			fmt.Printf("  タイトル: %s\n", titleElem.Text())
		}
	}

	fmt.Println("\n=== XPath検索のパターン ===")

	// 複数条件の検索例
	webBooks := doc.FindElements("//book[@category='web']")
	fmt.Printf("\nWebカテゴリの書籍数: %d\n", len(webBooks))

	// 属性値で絞り込み
	jpyPrices := doc.FindElements("//price[@currency='JPY']")
	fmt.Printf("JPY価格の書籍数: %d\n", len(jpyPrices))

	fmt.Println("\n要素検索完了")
}