package main

import (
	"fmt"
	"log"

	"github.com/beevik/etree"
)

func main_file_io() {
	fmt.Println("2. XMLファイル読み書きのサンプル")
	fmt.Println("==============================")

	// まずサンプルXMLを作成
	doc := createSampleLibrary()

	// XMLファイルに保存
	filename := "library.xml"
	fmt.Printf("\n%sに保存中...\n", filename)
	doc.Indent(2) // インデントを設定してから保存
	err := doc.WriteToFile(filename)
	if err != nil {
		log.Fatalf("ファイル保存エラー: %v", err)
	}
	fmt.Printf("%sに保存完了\n", filename)

	// XMLファイルから読み込み
	fmt.Printf("\n%sから読み込み中...\n", filename)
	newDoc := etree.NewDocument()
	err = newDoc.ReadFromFile(filename)
	if err != nil {
		log.Fatalf("ファイル読み込みエラー: %v", err)
	}
	fmt.Printf("%sから読み込み完了\n", filename)

	// 読み込んだ内容を確認
	fmt.Println("\n読み込んだXMLの内容:")
	newDoc.Indent(2)
	xmlString, _ := newDoc.WriteToString()
	fmt.Println(xmlString)

	// 読み込んだXMLの統計情報
	books := newDoc.SelectElements("//book")
	fmt.Printf("\n読み込んだ書籍数: %d冊\n", len(books))

	for i, book := range books {
		title := book.SelectElement("title")
		if title != nil {
			fmt.Printf("  %d. %s\n", i+1, title.Text())
		}
	}

	fmt.Println("\nファイル読み書き完了")
}

// createSampleLibrary サンプルの図書館XMLを作成
func createSampleLibrary() *etree.Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	library := etree.NewElement("library")
	library.CreateAttr("name", "技術書図書館")
	doc.SetRoot(library)

	// 書籍データ
	books := []struct {
		id       string
		category string
		title    string
		lang     string
		author   string
		price    string
		currency string
	}{
		{"1", "programming", "Go言語プログラミング", "ja", "山田太郎", "3200", "JPY"},
		{"2", "web", "JavaScript Essentials", "en", "John Doe", "29.99", "USD"},
		{"3", "database", "SQLデータベース入門", "ja", "田中花子", "2800", "JPY"},
	}

	for _, bookData := range books {
		book := library.CreateElement("book")
		book.CreateAttr("id", bookData.id)
		book.CreateAttr("category", bookData.category)

		title := book.CreateElement("title")
		title.CreateAttr("lang", bookData.lang)
		title.SetText(bookData.title)

		author := book.CreateElement("author")
		author.SetText(bookData.author)

		price := book.CreateElement("price")
		price.CreateAttr("currency", bookData.currency)
		price.SetText(bookData.price)
	}

	return doc
}