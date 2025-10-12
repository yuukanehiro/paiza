package main_test

import (
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/assert"
)

// TestModifyElements XML要素修正のテーブル形式テスト
func TestModifyElements(t *testing.T) {
	type args struct {
		operation string
		targetID  string
		newValue  string
	}

	type want struct {
		validateFunc   func(*etree.Document) bool
		expectedResult string
	}

	type test struct {
		name           string
		input          args
		want           want
		wantErr        bool
		wantErrMessage string
	}

	tests := []test{
		{
			name: "書籍価格の変更",
			input: args{
				operation: "price_change",
				targetID:  "1",
				newValue:  "3500",
			},
			want: want{
				validateFunc: func(doc *etree.Document) bool {
					book := doc.FindElement("//book[@id='1']")
					if book == nil {
						return false
					}
					price := book.SelectElement("price")
					return price != nil && price.Text() == "3500"
				},
				expectedResult: "価格が3500に変更されている",
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "新しい書籍の追加",
			input: args{
				operation: "add_book",
				targetID:  "4",
				newValue:  "AI入門",
			},
			want: want{
				validateFunc: func(doc *etree.Document) bool {
					books := doc.FindElements("//book")
					return len(books) == 4
				},
				expectedResult: "書籍が4冊になっている",
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "書籍の削除",
			input: args{
				operation: "delete_book",
				targetID:  "2",
				newValue:  "",
			},
			want: want{
				validateFunc: func(doc *etree.Document) bool {
					books := doc.FindElements("//book")
					return len(books) == 2
				},
				expectedResult: "書籍が2冊になっている",
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テスト用XMLを作成
			doc := createModifyTestXML()

			// 修正を実行
			switch tt.input.operation {
			case "price_change":
				book := doc.FindElement("//book[@id='" + tt.input.targetID + "']")
				if book != nil {
					price := book.SelectElement("price")
					if price != nil {
						price.SetText(tt.input.newValue)
					}
				}
			case "add_book":
				library := doc.SelectElement("library")
				if library != nil {
					newBook := library.CreateElement("book")
					newBook.CreateAttr("id", tt.input.targetID)
					newBook.CreateAttr("category", "ai")

					title := newBook.CreateElement("title")
					title.SetText(tt.input.newValue)

					author := newBook.CreateElement("author")
					author.SetText("AI太郎")

					price := newBook.CreateElement("price")
					price.SetText("4000")
				}
			case "delete_book":
				library := doc.SelectElement("library")
				book := doc.FindElement("//book[@id='" + tt.input.targetID + "']")
				if library != nil && book != nil {
					library.RemoveChild(book)
				}
			}

			// エラーが期待される場合のチェック
			if tt.wantErr {
				t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				return
			}

			// 結果を検証
			result := tt.want.validateFunc(doc)
			assert.True(t, result, "修正結果が期待通りではありません: %s", tt.want.expectedResult)
		})
	}
}

// createModifyTestXML 修正テスト用XMLを作成
func createModifyTestXML() *etree.Document {
	doc := etree.NewDocument()
	library := etree.NewElement("library")
	doc.SetRoot(library)

	// 3冊の書籍データ
	booksData := []struct {
		id       string
		category string
		title    string
		author   string
		price    string
	}{
		{"1", "programming", "Go入門", "田中太郎", "3000"},
		{"2", "web", "React基礎", "佐藤花子", "2500"},
		{"3", "database", "SQL実践", "山田次郎", "2800"},
	}

	for _, bookData := range booksData {
		book := library.CreateElement("book")
		book.CreateAttr("id", bookData.id)
		book.CreateAttr("category", bookData.category)

		title := book.CreateElement("title")
		title.SetText(bookData.title)

		author := book.CreateElement("author")
		author.SetText(bookData.author)

		price := book.CreateElement("price")
		price.SetText(bookData.price)
	}

	return doc
}