package main_test

import (
	"fmt"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/assert"
)

// TestBasicElementSearch 基本的な要素検索のテーブル形式テスト
func TestBasicElementSearch(t *testing.T) {
	type args struct {
		xpath string
	}

	type want struct {
		expectedCount int
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
			name: "全書籍検索",
			input: args{
				xpath: "//book",
			},
			want: want{
				expectedCount: 3,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "タイトル要素検索",
			input: args{
				xpath: "//title",
			},
			want: want{
				expectedCount: 3,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "著者要素検索",
			input: args{
				xpath: "//author",
			},
			want: want{
				expectedCount: 3,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "価格要素検索",
			input: args{
				xpath: "//price",
			},
			want: want{
				expectedCount: 3,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "存在しない要素検索",
			input: args{
				xpath: "//nonexistent",
			},
			want: want{
				expectedCount: 0,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := createSearchTestXML()
			elements := doc.FindElements(tt.input.xpath)

			// エラーが期待される場合のチェック
			if tt.wantErr {
				t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				return
			}

			assert.Equal(t, tt.want.expectedCount, len(elements), "検索結果の数が正しくありません")
		})
	}
}

// TestAttributeBasedSearch 属性ベースの検索のテーブル形式テスト
func TestAttributeBasedSearch(t *testing.T) {
	tests := []struct {
		name          string
		xpath         string
		expectedCount int
		expectedTexts []string
	}{
		{
			name:          "日本語タイトル検索",
			xpath:         "//title[@lang='ja']",
			expectedCount: 2,
			expectedTexts: []string{"Go言語プログラミング", "SQLデータベース入門"},
		},
		{
			name:          "英語タイトル検索",
			xpath:         "//title[@lang='en']",
			expectedCount: 1,
			expectedTexts: []string{"JavaScript Essentials"},
		},
		{
			name:          "プログラミングカテゴリ検索",
			xpath:         "//book[@category='programming']",
			expectedCount: 1,
			expectedTexts: nil, // 書籍要素なのでテキストチェックなし
		},
		{
			name:          "webカテゴリ検索",
			xpath:         "//book[@category='web']",
			expectedCount: 1,
			expectedTexts: nil,
		},
		{
			name:          "databaseカテゴリ検索",
			xpath:         "//book[@category='database']",
			expectedCount: 1,
			expectedTexts: nil,
		},
		{
			name:          "JPY通貨の価格検索",
			xpath:         "//price[@currency='JPY']",
			expectedCount: 2,
			expectedTexts: []string{"3200", "2800"},
		},
		{
			name:          "USD通貨の価格検索",
			xpath:         "//price[@currency='USD']",
			expectedCount: 1,
			expectedTexts: []string{"29.99"},
		},
		{
			name:          "存在しない属性値検索",
			xpath:         "//book[@category='nonexistent']",
			expectedCount: 0,
			expectedTexts: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := createSearchTestXML()
			elements := doc.FindElements(tt.xpath)

			// 検索結果数の確認
			assert.Equal(t, tt.expectedCount, len(elements), "検索結果の数が正しくありません")

			// 期待されるテキストの確認
			if tt.expectedTexts != nil {
				actualTexts := make([]string, len(elements))
				for i, elem := range elements {
					actualTexts[i] = elem.Text()
				}

				for _, expectedText := range tt.expectedTexts {
					assert.Contains(t, actualTexts, expectedText, "期待されるテキストが見つかりません")
				}
			}
		})
	}
}

// TestSpecificElementSearch 特定要素検索のテーブル形式テスト
func TestSpecificElementSearch(t *testing.T) {
	tests := []struct {
		name         string
		xpath        string
		expectedText string
		expectFound  bool
	}{
		{
			name:         "ID=1の書籍検索",
			xpath:        "//book[@id='1']",
			expectedText: "", // 書籍要素なのでテキストは空
			expectFound:  true,
		},
		{
			name:         "ID=2の書籍検索",
			xpath:        "//book[@id='2']",
			expectedText: "",
			expectFound:  true,
		},
		{
			name:         "ID=3の書籍検索",
			xpath:        "//book[@id='3']",
			expectedText: "",
			expectFound:  true,
		},
		{
			name:         "存在しないID検索",
			xpath:        "//book[@id='999']",
			expectedText: "",
			expectFound:  false,
		},
		{
			name:         "ID=1の書籍のタイトル検索",
			xpath:        "//book[@id='1']/title",
			expectedText: "Go言語プログラミング",
			expectFound:  true,
		},
		{
			name:         "ID=2の書籍の著者検索",
			xpath:        "//book[@id='2']/author",
			expectedText: "John Doe",
			expectFound:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := createSearchTestXML()
			element := doc.FindElement(tt.xpath)

			// 要素の存在確認
			if tt.expectFound {
				assert.NotNil(t, element, "要素が見つかりませんでした: %s", tt.xpath)
				// テキストの確認（期待されるテキストがある場合）
				if tt.expectedText != "" {
					actualText := element.Text()
					assert.Equal(t, tt.expectedText, actualText, "要素のテキストが正しくありません")
				}
			} else {
				assert.Nil(t, element, "要素が見つかるべきではありませんでした: %s", tt.xpath)
			}
		})
	}
}

// TestComplexSearch 複合条件検索のテーブル形式テスト
func TestComplexSearch(t *testing.T) {
	tests := []struct {
		name               string
		searchDescription  string
		validateFunc       func(*etree.Document) (bool, string)
	}{
		{
			name:              "日本語のプログラミング書籍検索",
			searchDescription: "プログラミングカテゴリかつ日本語タイトルの書籍",
			validateFunc: func(doc *etree.Document) (bool, string) {
				books := doc.FindElements("//book[@category='programming']")
				for _, book := range books {
					title := book.SelectElement("title")
					if title != nil && title.SelectAttrValue("lang", "") == "ja" {
						return true, "条件に合う書籍が見つかりました"
					}
				}
				return false, "プログラミングカテゴリの日本語書籍が見つかりません"
			},
		},
		{
			name:              "JPY価格の書籍数確認",
			searchDescription: "JPY通貨で価格が設定されている書籍数が2冊",
			validateFunc: func(doc *etree.Document) (bool, string) {
				jpyPrices := doc.FindElements("//price[@currency='JPY']")
				if len(jpyPrices) == 2 {
					return true, "JPY価格の書籍が2冊見つかりました"
				}
				return false, fmt.Sprintf("JPY価格の書籍数が正しくありません。期待値: 2, 実際値: %d", len(jpyPrices))
			},
		},
		{
			name:              "各書籍の必須要素確認",
			searchDescription: "全ての書籍にtitle、author、price要素が存在する",
			validateFunc: func(doc *etree.Document) (bool, string) {
				books := doc.FindElements("//book")
				for i, book := range books {
					if book.SelectElement("title") == nil {
						return false, fmt.Sprintf("書籍%dにtitle要素がありません", i+1)
					}
					if book.SelectElement("author") == nil {
						return false, fmt.Sprintf("書籍%dにauthor要素がありません", i+1)
					}
					if book.SelectElement("price") == nil {
						return false, fmt.Sprintf("書籍%dにprice要素がありません", i+1)
					}
				}
				return true, "全ての書籍に必須要素が存在します"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := createSearchTestXML()
			success, message := tt.validateFunc(doc)

			assert.True(t, success, "検索テストが失敗しました: %s - %s", tt.searchDescription, message)
			t.Logf("検索テストが成功しました: %s", message)
		})
	}
}

// createSearchTestXML 検索テスト用のXMLドキュメントを作成
func createSearchTestXML() *etree.Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	library := etree.NewElement("library")
	library.CreateAttr("name", "検索テスト図書館")
	doc.SetRoot(library)

	// 書籍データ
	booksData := []struct {
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

	// 書籍要素を作成
	for _, bookData := range booksData {
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