package main_test

import (
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/assert"
)

// TestAttributeOperations 属性操作のテーブル形式テスト
func TestAttributeOperations(t *testing.T) {
	type args struct {
		operation string
		attrName  string
		attrValue string
	}

	type want struct {
		expectedFunc func(*etree.Element) bool
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
			name: "属性の追加",
			input: args{
				operation: "add",
				attrName:  "rating",
				attrValue: "5",
			},
			want: want{
				expectedFunc: func(elem *etree.Element) bool {
					return elem.SelectAttrValue("rating", "") == "5"
				},
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "属性の修正",
			input: args{
				operation: "modify",
				attrName:  "category",
				attrValue: "advanced",
			},
			want: want{
				expectedFunc: func(elem *etree.Element) bool {
					return elem.SelectAttrValue("category", "") == "advanced"
				},
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "属性の削除",
			input: args{
				operation: "remove",
				attrName:  "category",
				attrValue: "",
			},
			want: want{
				expectedFunc: func(elem *etree.Element) bool {
					return elem.SelectAttrValue("category", "not_found") == "not_found"
				},
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := createAttributeTestXML()
			book := doc.FindElement("//book[@id='1']")
			assert.NotNil(t, book, "テスト対象の書籍が見つかりません")

			switch tt.input.operation {
			case "add":
				book.CreateAttr(tt.input.attrName, tt.input.attrValue)
			case "modify":
				attr := book.SelectAttr(tt.input.attrName)
				if attr != nil {
					attr.Value = tt.input.attrValue
				}
			case "remove":
				book.RemoveAttr(tt.input.attrName)
			}

			// エラーが期待される場合のチェック
			if tt.wantErr {
				t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				return
			}

			result := tt.want.expectedFunc(book)
			assert.True(t, result, "属性操作の結果が期待通りではありません")
		})
	}
}

// TestAttributeSearch 属性による検索のテーブル形式テスト
func TestAttributeSearch(t *testing.T) {
	tests := []struct {
		name          string
		xpath         string
		expectedCount int
	}{
		{
			name:          "特定属性値での検索",
			xpath:         "//book[@category='programming']",
			expectedCount: 1,
		},
		{
			name:          "属性の存在チェック",
			xpath:         "//book[@id]",
			expectedCount: 2,
		},
		{
			name:          "存在しない属性値での検索",
			xpath:         "//book[@category='nonexistent']",
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := createAttributeTestXML()
			elements := doc.FindElements(tt.xpath)
			assert.Equal(t, tt.expectedCount, len(elements), "検索結果の数が正しくありません")
		})
	}
}

// createAttributeTestXML 属性テスト用XMLを作成
func createAttributeTestXML() *etree.Document {
	doc := etree.NewDocument()
	library := etree.NewElement("library")
	doc.SetRoot(library)

	// 書籍1
	book1 := library.CreateElement("book")
	book1.CreateAttr("id", "1")
	book1.CreateAttr("category", "programming")

	title1 := book1.CreateElement("title")
	title1.SetText("Go入門")

	// 書籍2
	book2 := library.CreateElement("book")
	book2.CreateAttr("id", "2")
	book2.CreateAttr("category", "web")

	title2 := book2.CreateElement("title")
	title2.SetText("JavaScript基礎")

	return doc
}