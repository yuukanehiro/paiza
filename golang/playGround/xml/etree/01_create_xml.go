package main

import (
	"fmt"

	"github.com/beevik/etree"
)

func main_create_xml() {
	fmt.Println("1. 基本的なXML作成のサンプル")
	fmt.Println("===========================")

	// 新しいXMLドキュメントを作成
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// ルート要素を作成
	library := etree.NewElement("library")
	library.CreateAttr("name", "技術書図書館")
	library.CreateAttr("location", "東京")
	doc.SetRoot(library)

	// 1冊目の書籍
	book1 := library.CreateElement("book")
	book1.CreateAttr("id", "1")
	book1.CreateAttr("category", "programming")

	title1 := book1.CreateElement("title")
	title1.CreateAttr("lang", "ja")
	title1.SetText("Go言語プログラミング入門")

	author1 := book1.CreateElement("author")
	author1.SetText("山田太郎")

	price1 := book1.CreateElement("price")
	price1.CreateAttr("currency", "JPY")
	price1.SetText("3200")

	// 2冊目の書籍
	book2 := library.CreateElement("book")
	book2.CreateAttr("id", "2")
	book2.CreateAttr("category", "web")

	title2 := book2.CreateElement("title")
	title2.CreateAttr("lang", "en")
	title2.SetText("JavaScript Essentials")

	author2 := book2.CreateElement("author")
	author2.SetText("John Doe")

	price2 := book2.CreateElement("price")
	price2.CreateAttr("currency", "USD")
	price2.SetText("29.99")

	// XMLを整形して出力
	doc.Indent(2)
	xmlString, err := doc.WriteToString()
	if err != nil {
		fmt.Printf("XMLの文字列化エラー: %v\n", err)
		return
	}

	fmt.Println("\n作成したXML:")
	fmt.Println(xmlString)

	fmt.Println("\n基本的なXML作成完了")
}