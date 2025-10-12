package main

import (
	"fmt"

	"github.com/beevik/etree"
)

func main_namespaces() {
	fmt.Println("6. XML名前空間のサンプル")
	fmt.Println("========================")

	fmt.Println("=== 名前空間付きXMLの作成 ===")

	// 名前空間付きXMLドキュメントを作成
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// ルート要素に名前空間を定義
	catalog := etree.NewElement("catalog")
	catalog.CreateAttr("xmlns", "http://example.com/library")           // デフォルト名前空間
	catalog.CreateAttr("xmlns:book", "http://example.com/book")         // book名前空間
	catalog.CreateAttr("xmlns:author", "http://example.com/author")     // author名前空間
	catalog.CreateAttr("xmlns:price", "http://example.com/price")       // price名前空間
	doc.SetRoot(catalog)

	// 名前空間付きのコレクション要素
	collection := catalog.CreateElement("book:collection")
	collection.CreateAttr("name", "技術書コレクション")
	collection.CreateAttr("version", "1.0")

	// 1冊目の書籍（名前空間付き要素）
	item1 := collection.CreateElement("book:item")
	item1.CreateAttr("id", "ns1")

	title1 := item1.CreateElement("book:title")
	title1.CreateAttr("lang", "ja")
	title1.SetText("名前空間を使ったXML設計")

	authorInfo1 := item1.CreateElement("author:info")
	authorName1 := authorInfo1.CreateElement("author:name")
	authorName1.SetText("名前空間 太郎")
	authorCountry1 := authorInfo1.CreateElement("author:country")
	authorCountry1.SetText("Japan")

	priceInfo1 := item1.CreateElement("price:details")
	amount1 := priceInfo1.CreateElement("price:amount")
	amount1.CreateAttr("currency", "JPY")
	amount1.SetText("3800")

	// 2冊目の書籍
	item2 := collection.CreateElement("book:item")
	item2.CreateAttr("id", "ns2")

	title2 := item2.CreateElement("book:title")
	title2.CreateAttr("lang", "en")
	title2.SetText("XML Namespaces Guide")

	authorInfo2 := item2.CreateElement("author:info")
	authorName2 := authorInfo2.CreateElement("author:name")
	authorName2.SetText("John Namespace")
	authorCountry2 := authorInfo2.CreateElement("author:country")
	authorCountry2.SetText("USA")

	priceInfo2 := item2.CreateElement("price:details")
	amount2 := priceInfo2.CreateElement("price:amount")
	amount2.CreateAttr("currency", "USD")
	amount2.SetText("45.99")

	// 作成したXMLを表示
	fmt.Println("\n作成した名前空間付きXML:")
	doc.Indent(2)
	xmlString, _ := doc.WriteToString()
	fmt.Println(xmlString)

	fmt.Println("\n=== 名前空間を考慮した検索 ===")

	// 名前空間付き要素の検索
	bookItems := doc.FindElements("//book:item")
	fmt.Printf("book:item要素の数: %d\n", len(bookItems))

	// 名前空間付きタイトルの検索
	bookTitles := doc.FindElements("//book:title")
	fmt.Println("\n全ての書籍タイトル:")
	for i, title := range bookTitles {
		lang := title.SelectAttrValue("lang", "")
		fmt.Printf("  %d. %s (%s)\n", i+1, title.Text(), lang)
	}

	// 著者情報の検索
	authorNames := doc.FindElements("//author:name")
	fmt.Println("\n全ての著者名:")
	for i, name := range authorNames {
		fmt.Printf("  %d. %s\n", i+1, name.Text())
	}

	// 価格情報の検索
	priceAmounts := doc.FindElements("//price:amount")
	fmt.Println("\n全ての価格:")
	for i, amount := range priceAmounts {
		currency := amount.SelectAttrValue("currency", "")
		fmt.Printf("  %d. %s %s\n", i+1, amount.Text(), currency)
	}

	fmt.Println("\n=== 特定名前空間の操作 ===")

	// 特定の名前空間の要素を修正
	firstItem := doc.FindElement("//book:item[@id='ns1']")
	if firstItem != nil {
		// 新しい名前空間付き要素を追加
		isbn := firstItem.CreateElement("book:isbn")
		isbn.SetText("978-4-12345-678-9")
		fmt.Println("最初の書籍にISBNを追加")

		// 価格を変更
		priceAmount := firstItem.SelectElement("price:details/price:amount")
		if priceAmount != nil {
			oldPrice := priceAmount.Text()
			priceAmount.SetText("4200")
			fmt.Printf("価格を %s円 から 4200円 に変更\n", oldPrice)
		}
	}

	fmt.Println("\n=== 異なる名前空間のXMLファイル作成 ===")

	// 別の名前空間を使ったXMLも作成
	doc2 := etree.NewDocument()
	doc2.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	library := etree.NewElement("lib:library")
	library.CreateAttr("xmlns:lib", "http://library.example.com")
	library.CreateAttr("xmlns:meta", "http://metadata.example.com")
	doc2.SetRoot(library)

	section := library.CreateElement("lib:section")
	section.CreateAttr("category", "technical")

	document := section.CreateElement("lib:document")
	document.CreateAttr("id", "doc1")

	metadata := document.CreateElement("meta:metadata")
	created := metadata.CreateElement("meta:created")
	created.SetText("2024-01-01")
	modified := metadata.CreateElement("meta:modified")
	modified.SetText("2024-10-12")

	content := document.CreateElement("lib:content")
	content.SetText("名前空間を使ったXMLドキュメントの例")

	fmt.Println("別の名前空間を使ったXML:")
	doc2.Indent(2)
	xmlString2, _ := doc2.WriteToString()
	fmt.Println(xmlString2)

	// ファイルに保存
	fmt.Println("\n=== ファイルに保存 ===")
	err := doc.WriteToFile("namespace_catalog.xml")
	if err != nil {
		fmt.Printf("ファイル保存エラー: %v\n", err)
	} else {
		fmt.Println("名前空間付きXMLをnamespace_catalog.xmlに保存")
	}

	err = doc2.WriteToFile("namespace_library.xml")
	if err != nil {
		fmt.Printf("ファイル保存エラー: %v\n", err)
	} else {
		fmt.Println("別の名前空間XMLをnamespace_library.xmlに保存")
	}

	fmt.Println("\n名前空間操作完了")
}