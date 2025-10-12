package main_test

import (
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/assert"
)

// TestCreateXMLDocument XMLドキュメント作成のテーブル形式テスト
func TestCreateXMLDocument(t *testing.T) {
	type args struct {
		libraryName    string
		libraryLocation string
	}

	type want struct {
		expectedTag      string
		expectedName     string
		expectedLocation string
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
			name: "基本的な図書館XML作成",
			input: args{
				libraryName:     "技術書図書館",
				libraryLocation: "東京",
			},
			want: want{
				expectedTag:      "library",
				expectedName:     "技術書図書館",
				expectedLocation: "東京",
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "英語名の図書館XML作成",
			input: args{
				libraryName:     "Tech Library",
				libraryLocation: "Tokyo",
			},
			want: want{
				expectedTag:      "library",
				expectedName:     "Tech Library",
				expectedLocation: "Tokyo",
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "空文字列属性のテスト",
			input: args{
				libraryName:     "",
				libraryLocation: "",
			},
			want: want{
				expectedTag:      "library",
				expectedName:     "",
				expectedLocation: "",
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// XMLドキュメントを作成
			doc := etree.NewDocument()
			doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

			// ルート要素を作成
			library := etree.NewElement("library")
			library.CreateAttr("name", tt.input.libraryName)
			library.CreateAttr("location", tt.input.libraryLocation)
			doc.SetRoot(library)

			// エラーが期待される場合のチェック
			if tt.wantErr {
				t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				return
			}

			// テスト: ルート要素の確認
			assert.Equal(t, tt.want.expectedTag, doc.Root().Tag, "ルート要素のタグが正しくありません")

			// テスト: 属性の確認
			actualName := library.SelectAttrValue("name", "")
			assert.Equal(t, tt.want.expectedName, actualName, "name属性が正しくありません")

			actualLocation := library.SelectAttrValue("location", "")
			assert.Equal(t, tt.want.expectedLocation, actualLocation, "location属性が正しくありません")
		})
	}
}

// TestCreateBookElements 書籍要素作成のテーブル形式テスト
func TestCreateBookElements(t *testing.T) {
	type args struct {
		id        string
		category  string
		title     string
		titleLang string
		author    string
		price     string
		currency  string
	}

	type want struct {
		hasBook    bool
		bookCount  int
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
			name: "日本語プログラミング書籍",
			input: args{
				id:        "1",
				category:  "programming",
				title:     "Go言語プログラミング入門",
				titleLang: "ja",
				author:    "山田太郎",
				price:     "3200",
				currency:  "JPY",
			},
			want: want{
				hasBook:   true,
				bookCount: 1,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "英語Web開発書籍",
			input: args{
				id:        "2",
				category:  "web",
				title:     "JavaScript Essentials",
				titleLang: "en",
				author:    "John Doe",
				price:     "29.99",
				currency:  "USD",
			},
			want: want{
				hasBook:   true,
				bookCount: 1,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "空文字列テスト",
			input: args{
				id:        "",
				category:  "",
				title:     "",
				titleLang: "",
				author:    "",
				price:     "",
				currency:  "",
			},
			want: want{
				hasBook:   true,
				bookCount: 1,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// XMLドキュメントを作成
			doc := etree.NewDocument()
			library := etree.NewElement("library")
			doc.SetRoot(library)

			// 書籍要素を作成
			book := library.CreateElement("book")
			book.CreateAttr("id", tt.input.id)
			book.CreateAttr("category", tt.input.category)

			title := book.CreateElement("title")
			title.CreateAttr("lang", tt.input.titleLang)
			title.SetText(tt.input.title)

			author := book.CreateElement("author")
			author.SetText(tt.input.author)

			price := book.CreateElement("price")
			price.CreateAttr("currency", tt.input.currency)
			price.SetText(tt.input.price)

			// エラーが期待される場合のチェック
			if tt.wantErr {
				t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				return
			}

			// テスト: 書籍の存在確認
			books := doc.FindElements("//book")
			if tt.want.hasBook {
				assert.Equal(t, tt.want.bookCount, len(books), "書籍数が正しくありません")
			}

			// テスト: 書籍属性の確認
			actualID := book.SelectAttrValue("id", "")
			assert.Equal(t, tt.input.id, actualID, "ID属性が正しくありません")

			actualCategory := book.SelectAttrValue("category", "")
			assert.Equal(t, tt.input.category, actualCategory, "category属性が正しくありません")

			// テスト: タイトル要素の確認
			titleElem := book.SelectElement("title")
			assert.NotNil(t, titleElem, "title要素が見つかりません")
			assert.Equal(t, tt.input.title, titleElem.Text(), "タイトルが正しくありません")

			actualTitleLang := titleElem.SelectAttrValue("lang", "")
			assert.Equal(t, tt.input.titleLang, actualTitleLang, "title要素のlang属性が正しくありません")

			// テスト: 著者要素の確認
			authorElem := book.SelectElement("author")
			assert.NotNil(t, authorElem, "author要素が見つかりません")
			assert.Equal(t, tt.input.author, authorElem.Text(), "著者が正しくありません")

			// テスト: 価格要素の確認
			priceElem := book.SelectElement("price")
			assert.NotNil(t, priceElem, "price要素が見つかりません")
			assert.Equal(t, tt.input.price, priceElem.Text(), "価格が正しくありません")

			actualCurrency := priceElem.SelectAttrValue("currency", "")
			assert.Equal(t, tt.input.currency, actualCurrency, "currency属性が正しくありません")
		})
	}
}

// TestXMLOutput XML出力形式のテーブル形式テスト
func TestXMLOutput(t *testing.T) {
	type args struct {
		rootTag        string
		rootAttr       map[string]string
		childTag       string
		childText      string
		indent         int
	}

	type want struct {
		expectedContains []string
		hasNewlines      bool
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
			name: "基本的なXML出力テスト",
			input: args{
				rootTag: "test",
				rootAttr: map[string]string{
					"version": "1.0",
				},
				childTag:  "child",
				childText: "テストテキスト",
				indent:    2,
			},
			want: want{
				expectedContains: []string{
					`<?xml version="1.0" encoding="UTF-8"?>`,
					`<test version="1.0">`,
					`<child>テストテキスト</child>`,
					`</test>`,
				},
				hasNewlines: true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "複数属性のXML出力テスト",
			input: args{
				rootTag: "library",
				rootAttr: map[string]string{
					"name":     "テスト図書館",
					"location": "テスト市",
				},
				childTag:  "description",
				childText: "図書館の説明",
				indent:    4,
			},
			want: want{
				expectedContains: []string{
					`<library`,
					`name="テスト図書館"`,
					`location="テスト市"`,
					`<description>図書館の説明</description>`,
				},
				hasNewlines: true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "インデントなしのXML出力テスト",
			input: args{
				rootTag:   "simple",
				rootAttr:  map[string]string{},
				childTag:  "data",
				childText: "シンプルなデータ",
				indent:    0,
			},
			want: want{
				expectedContains: []string{
					`<simple>`,
					`<data>シンプルなデータ</data>`,
					`</simple>`,
				},
				hasNewlines: false,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := etree.NewDocument()
			doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

			root := etree.NewElement(tt.input.rootTag)
			for key, value := range tt.input.rootAttr {
				root.CreateAttr(key, value)
			}
			doc.SetRoot(root)

			if tt.input.childTag != "" {
				child := root.CreateElement(tt.input.childTag)
				child.SetText(tt.input.childText)
			}

			// インデントを設定
			if tt.input.indent > 0 {
				doc.Indent(tt.input.indent)
			}

			xmlString, err := doc.WriteToString()

			// エラーが期待される場合のチェック
			if tt.wantErr {
				assert.Error(t, err, "エラーが期待されましたが、エラーが発生しませんでした")
				if err != nil {
					assert.Contains(t, err.Error(), tt.wantErrMessage)
				}
				return
			}

			assert.NoError(t, err, "XML文字列化エラー")

			// 期待される文字列が含まれているかチェック
			for _, expected := range tt.want.expectedContains {
				assert.Contains(t, xmlString, expected, "期待される文字列が含まれていません")
			}

			// インデントのチェック
			if tt.want.hasNewlines {
				assert.Contains(t, xmlString, "\n", "インデントが正しく設定されていません（改行が含まれていない）")
			}
		})
	}
}