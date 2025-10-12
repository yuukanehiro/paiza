package main

import (
	"fmt"

	"github.com/beevik/etree"
)

func main_modify() {
	fmt.Println("4. XML要素修正のサンプル")
	fmt.Println("========================")

	// XMLファイルを読み込み
	doc := etree.NewDocument()
	err := doc.ReadFromFile("library.xml")
	if err != nil {
		fmt.Printf("ファイル読み込みエラー: %v\n", err)
		fmt.Println("library.xmlファイルが必要です。先に02_file_io.goを実行してください。")
		return
	}

	fmt.Println("=== 修正前の状態 ===")
	printAllBooks(doc)

	fmt.Println("\n=== 要素の修正 ===")

	// 1. 要素のテキストを変更
	fmt.Println("\n1. ID=1の書籍の価格を変更")
	book1 := doc.FindElement("//book[@id='1']")
	if book1 != nil {
		priceElem := book1.SelectElement("price")
		if priceElem != nil {
			oldPrice := priceElem.Text()
			priceElem.SetText("3500")
			titleElem := book1.SelectElement("title")
			if titleElem != nil {
				fmt.Printf("「%s」の価格を %s円 から 3500円 に変更\n", titleElem.Text(), oldPrice)
			}
		}
	}

	// 2. 新しい要素を追加
	fmt.Println("\n2. 新しい書籍を追加")
	library := doc.SelectElement("library")
	if library != nil {
		newBook := library.CreateElement("book")
		newBook.CreateAttr("id", "4")
		newBook.CreateAttr("category", "ai")

		title := newBook.CreateElement("title")
		title.CreateAttr("lang", "ja")
		title.SetText("機械学習実践ガイド")

		author := newBook.CreateElement("author")
		author.SetText("佐藤智子")

		price := newBook.CreateElement("price")
		price.CreateAttr("currency", "JPY")
		price.SetText("4200")

		fmt.Println("新しい書籍「機械学習実践ガイド」を追加")
	}

	// 3. 要素の削除
	fmt.Println("\n3. ID=2の書籍を削除")
	book2 := doc.FindElement("//book[@id='2']")
	if book2 != nil && library != nil {
		titleElem := book2.SelectElement("title")
		titleText := ""
		if titleElem != nil {
			titleText = titleElem.Text()
		}
		library.RemoveChild(book2)
		fmt.Printf("書籍「%s」を削除\n", titleText)
	}

	fmt.Println("\n=== 修正後の状態 ===")
	printAllBooks(doc)

	// 修正したXMLを新しいファイルに保存
	fmt.Println("\n=== ファイルに保存 ===")
	doc.Indent(2)
	err = doc.WriteToFile("modified_library.xml")
	if err != nil {
		fmt.Printf("ファイル保存エラー: %v\n", err)
	} else {
		fmt.Println("修正したXMLをmodified_library.xmlに保存")
	}

	fmt.Println("\n要素修正完了")
}

// printAllBooks 全ての書籍情報を表示
func printAllBooks(doc *etree.Document) {
	books := doc.FindElements("//book")
	fmt.Printf("書籍数: %d冊\n", len(books))

	for i, book := range books {
		id := book.SelectAttrValue("id", "")

		titleElem := book.SelectElement("title")
		title := ""
		if titleElem != nil {
			title = titleElem.Text()
		}

		authorElem := book.SelectElement("author")
		author := ""
		if authorElem != nil {
			author = authorElem.Text()
		}

		priceElem := book.SelectElement("price")
		price := ""
		currency := ""
		if priceElem != nil {
			price = priceElem.Text()
			currency = priceElem.SelectAttrValue("currency", "")
		}

		fmt.Printf("  %d. [ID:%s] %s - %s (%s %s)\n", i+1, id, title, author, price, currency)
	}
}