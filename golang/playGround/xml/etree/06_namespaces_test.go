package main_test

import (
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/assert"
)

// TestNamespaceCreation 名前空間付きXML作成のテーブル形式テスト
func TestNamespaceCreation(t *testing.T) {
	type args struct {
		namespaceURI string
		prefix       string
		elementName  string
	}

	type want struct {
		expectedQName     string
		expectedNamespace string
		elementExists     bool
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
			name: "book名前空間の要素作成",
			input: args{
				namespaceURI: "http://example.com/book",
				prefix:       "book",
				elementName:  "item",
			},
			want: want{
				expectedQName:     "book:item",
				expectedNamespace: "http://example.com/book",
				elementExists:     true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "author名前空間の要素作成",
			input: args{
				namespaceURI: "http://example.com/author",
				prefix:       "author",
				elementName:  "info",
			},
			want: want{
				expectedQName:     "author:info",
				expectedNamespace: "http://example.com/author",
				elementExists:     true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "デフォルト名前空間の要素作成",
			input: args{
				namespaceURI: "http://example.com/default",
				prefix:       "",
				elementName:  "catalog",
			},
			want: want{
				expectedQName:     "catalog",
				expectedNamespace: "http://example.com/default",
				elementExists:     true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := createNamespaceTestXML()

			// 名前空間付き要素を検索
			var element *etree.Element
			if tt.input.prefix != "" {
				element = doc.FindElement("//" + tt.input.prefix + ":" + tt.input.elementName)
			} else {
				element = doc.FindElement("//" + tt.input.elementName)
			}

			// エラーが期待される場合のチェック
			if tt.wantErr {
				t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				return
			}

			if tt.want.elementExists {
				assert.NotNil(t, element, "名前空間付き要素が見つかりません: %s", tt.want.expectedQName)
				if element != nil {
					assert.Equal(t, tt.want.expectedQName, element.FullTag(), "要素の完全名が正しくありません")
				}
			} else {
				assert.Nil(t, element, "要素が見つかるべきではありませんでした")
			}
		})
	}
}

// TestNamespaceSearch 名前空間を考慮した検索のテーブル形式テスト
func TestNamespaceSearch(t *testing.T) {
	type args struct {
		xpath       string
		description string
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
			name: "book名前空間の要素検索",
			input: args{
				xpath:       "//book:item",
				description: "book:item要素の検索",
			},
			want: want{
				expectedCount: 2,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "author名前空間の要素検索",
			input: args{
				xpath:       "//author:name",
				description: "author:name要素の検索",
			},
			want: want{
				expectedCount: 2,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "特定属性を持つ名前空間要素の検索",
			input: args{
				xpath:       "//book:item[@id='ns1']",
				description: "特定IDのbook:item検索",
			},
			want: want{
				expectedCount: 1,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "存在しない名前空間要素の検索",
			input: args{
				xpath:       "//nonexistent:element",
				description: "存在しない名前空間要素の検索",
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
			doc := createNamespaceTestXML()
			elements := doc.FindElements(tt.input.xpath)

			// エラーが期待される場合のチェック
			if tt.wantErr {
				t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				return
			}

			assert.Equal(t, tt.want.expectedCount, len(elements), "検索結果の数が正しくありません: %s", tt.input.description)
		})
	}
}

// TestNamespaceAttributes 名前空間宣言のテスト
func TestNamespaceAttributes(t *testing.T) {
	type args struct {
		attrName string
	}

	type want struct {
		expectedValue string
		shouldExist   bool
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
			name: "book名前空間宣言の確認",
			input: args{
				attrName: "xmlns:book",
			},
			want: want{
				expectedValue: "http://example.com/book",
				shouldExist:   true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "author名前空間宣言の確認",
			input: args{
				attrName: "xmlns:author",
			},
			want: want{
				expectedValue: "http://example.com/author",
				shouldExist:   true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "存在しない名前空間の確認",
			input: args{
				attrName: "xmlns:nonexistent",
			},
			want: want{
				expectedValue: "",
				shouldExist:   false,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := createNamespaceTestXML()
			root := doc.Root()
			value := root.SelectAttrValue(tt.input.attrName, "")

			// エラーが期待される場合のチェック
			if tt.wantErr {
				t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				return
			}

			if tt.want.shouldExist {
				assert.Equal(t, tt.want.expectedValue, value, "名前空間宣言の値が正しくありません")
			} else {
				assert.Empty(t, value, "存在しないはずの名前空間宣言が見つかりました")
			}
		})
	}
}

// createNamespaceTestXML 名前空間テスト用XMLを作成
func createNamespaceTestXML() *etree.Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// ルート要素に名前空間を定義
	catalog := etree.NewElement("catalog")
	catalog.CreateAttr("xmlns:book", "http://example.com/book")
	catalog.CreateAttr("xmlns:author", "http://example.com/author")
	doc.SetRoot(catalog)

	// 名前空間付きのコレクション要素
	collection := catalog.CreateElement("book:collection")
	collection.CreateAttr("name", "テストコレクション")

	// 1冊目の書籍
	item1 := collection.CreateElement("book:item")
	item1.CreateAttr("id", "ns1")

	title1 := item1.CreateElement("book:title")
	title1.SetText("名前空間テスト1")

	authorInfo1 := item1.CreateElement("author:info")
	authorName1 := authorInfo1.CreateElement("author:name")
	authorName1.SetText("テスト著者1")

	// 2冊目の書籍
	item2 := collection.CreateElement("book:item")
	item2.CreateAttr("id", "ns2")

	title2 := item2.CreateElement("book:title")
	title2.SetText("名前空間テスト2")

	authorInfo2 := item2.CreateElement("author:info")
	authorName2 := authorInfo2.CreateElement("author:name")
	authorName2.SetText("テスト著者2")

	return doc
}