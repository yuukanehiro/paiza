package main_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/assert"
)

// TestWriteToFile XMLファイル書き込みのテーブル形式テスト
func TestWriteToFile(t *testing.T) {
	type args struct {
		filename    string
		libraryName string
		bookCount   int
	}

	type want struct {
		expectedFile bool
		cleanup      bool
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
			name: "基本的なXMLファイル書き込み",
			input: args{
				filename:    "test_library.xml",
				libraryName: "テスト図書館",
				bookCount:   2,
			},
			want: want{
				expectedFile: true,
				cleanup:      true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "日本語ファイル名でのXML書き込み",
			input: args{
				filename:    "テスト図書館.xml",
				libraryName: "日本語図書館",
				bookCount:   1,
			},
			want: want{
				expectedFile: true,
				cleanup:      true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "書籍なしのXMLファイル書き込み",
			input: args{
				filename:    "empty_library.xml",
				libraryName: "空の図書館",
				bookCount:   0,
			},
			want: want{
				expectedFile: true,
				cleanup:      true,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テスト用XMLドキュメントを作成
			doc := createTestLibraryXML(tt.input.libraryName, tt.input.bookCount)

			// ファイルに書き込み
			doc.Indent(2)
			err := doc.WriteToFile(tt.input.filename)

			// エラーが期待される場合のチェック
			if tt.wantErr {
				assert.Error(t, err, "エラーが期待されましたが、エラーが発生しませんでした")
				if err != nil {
					assert.Contains(t, err.Error(), tt.wantErrMessage)
				}
				return
			}

			assert.NoError(t, err, "ファイル書き込みエラー")

			// ファイルが存在するかチェック
			if tt.want.expectedFile {
				_, err := os.Stat(tt.input.filename)
				assert.False(t, os.IsNotExist(err), "ファイルが作成されていません: %s", tt.input.filename)
			}

			// クリーンアップ
			if tt.want.cleanup {
				defer func() {
					if err := os.Remove(tt.input.filename); err != nil {
						t.Logf("クリーンアップエラー: %v", err)
					}
				}()
			}
		})
	}
}

// TestReadFromFile XMLファイル読み込みのテーブル形式テスト
func TestReadFromFile(t *testing.T) {
	tests := []struct {
		name              string
		setupFilename     string
		readFilename      string
		libraryName       string
		bookCount         int
		expectedBookCount int
		expectError       bool
	}{
		{
			name:              "正常なXMLファイル読み込み",
			setupFilename:     "setup_test.xml",
			readFilename:      "setup_test.xml",
			libraryName:       "読み込みテスト図書館",
			bookCount:         3,
			expectedBookCount: 3,
			expectError:       false,
		},
		{
			name:              "書籍なしのXMLファイル読み込み",
			setupFilename:     "empty_setup_test.xml",
			readFilename:      "empty_setup_test.xml",
			libraryName:       "空の図書館",
			bookCount:         0,
			expectedBookCount: 0,
			expectError:       false,
		},
		{
			name:              "存在しないファイルの読み込み",
			setupFilename:     "",
			readFilename:      "nonexistent.xml",
			libraryName:       "",
			bookCount:         0,
			expectedBookCount: 0,
			expectError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// セットアップ: テストファイルを作成（存在しないファイルテスト以外）
			if tt.setupFilename != "" {
				setupDoc := createTestLibraryXML(tt.libraryName, tt.bookCount)
				setupDoc.Indent(2)
				err := setupDoc.WriteToFile(tt.setupFilename)
				if err != nil {
					t.Fatalf("セットアップファイル作成エラー: %v", err)
				}

				// クリーンアップ
				defer func() {
					if err := os.Remove(tt.setupFilename); err != nil {
						t.Logf("セットアップファイルクリーンアップエラー: %v", err)
					}
				}()
			}

			// ファイルから読み込み
			doc := etree.NewDocument()
			err := doc.ReadFromFile(tt.readFilename)

			// エラーの期待値チェック
			if tt.expectError {
				if err == nil {
					t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				}
				return
			}

			if err != nil {
				t.Fatalf("ファイル読み込みエラー: %v", err)
			}

			// 読み込んだデータの検証
			books := doc.FindElements("//book")
			if len(books) != tt.expectedBookCount {
				t.Errorf("読み込んだ書籍数が正しくありません。期待値: %d, 実際値: %d", tt.expectedBookCount, len(books))
			}

			// ルート要素の検証
			if tt.libraryName != "" {
				library := doc.SelectElement("library")
				if library == nil {
					t.Fatal("library要素が見つかりません")
				}
				actualName := library.SelectAttrValue("name", "")
				if actualName != tt.libraryName {
					t.Errorf("図書館名が正しくありません。期待値: %s, 実際値: %s", tt.libraryName, actualName)
				}
			}
		})
	}
}

// TestRoundTripFileIO ファイル書き込み→読み込みの一連テスト
func TestRoundTripFileIO(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		books       []testBookData
		expectedXML []string
	}{
		{
			name:     "複数書籍の書き込み→読み込みテスト",
			filename: "roundtrip_test.xml",
			books: []testBookData{
				{"1", "programming", "Go入門", "ja", "田中太郎", "3000", "JPY"},
				{"2", "web", "React Guide", "en", "John Smith", "49.99", "USD"},
			},
			expectedXML: []string{
				"Go入門",
				"React Guide",
				"田中太郎",
				"John Smith",
				`id="1"`,
				`id="2"`,
			},
		},
		{
			name:     "日本語文字の書き込み→読み込みテスト",
			filename: "japanese_test.xml",
			books: []testBookData{
				{"1", "プログラミング", "日本語プログラミング入門", "ja", "山田花子", "2500", "JPY"},
			},
			expectedXML: []string{
				"日本語プログラミング入門",
				"山田花子",
				"プログラミング",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 書き込み用ドキュメント作成
			writeDoc := createTestLibraryWithBooks("ラウンドトリップテスト", tt.books)

			// ファイルに書き込み
			writeDoc.Indent(2)
			err := writeDoc.WriteToFile(tt.filename)
			if err != nil {
				t.Fatalf("ファイル書き込みエラー: %v", err)
			}

			// クリーンアップ
			defer func() {
				if err := os.Remove(tt.filename); err != nil {
					t.Logf("クリーンアップエラー: %v", err)
				}
			}()

			// ファイルから読み込み
			readDoc := etree.NewDocument()
			err = readDoc.ReadFromFile(tt.filename)
			if err != nil {
				t.Fatalf("ファイル読み込みエラー: %v", err)
			}

			// XML文字列化して内容チェック
			readDoc.Indent(2)
			xmlString, err := readDoc.WriteToString()
			if err != nil {
				t.Fatalf("XML文字列化エラー: %v", err)
			}

			// 期待される文字列が含まれているかチェック
			for _, expected := range tt.expectedXML {
				if !strings.Contains(xmlString, expected) {
					t.Errorf("期待される文字列が含まれていません: %s\n実際のXML:\n%s", expected, xmlString)
				}
			}

			// 書籍数の確認
			books := readDoc.FindElements("//book")
			if len(books) != len(tt.books) {
				t.Errorf("読み込んだ書籍数が正しくありません。期待値: %d, 実際値: %d", len(tt.books), len(books))
			}
		})
	}
}

// testBookData テスト用書籍データ構造体
type testBookData struct {
	id       string
	category string
	title    string
	lang     string
	author   string
	price    string
	currency string
}

// createTestLibraryXML テスト用図書館XMLを作成（指定された書籍数で）
func createTestLibraryXML(libraryName string, bookCount int) *etree.Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	library := etree.NewElement("library")
	library.CreateAttr("name", libraryName)
	doc.SetRoot(library)

	// 指定された数の書籍を作成
	for i := 0; i < bookCount; i++ {
		book := library.CreateElement("book")
		book.CreateAttr("id", fmt.Sprintf("%d", i+1))
		book.CreateAttr("category", "test")

		title := book.CreateElement("title")
		title.CreateAttr("lang", "ja")
		title.SetText(fmt.Sprintf("テスト書籍%d", i+1))

		author := book.CreateElement("author")
		author.SetText(fmt.Sprintf("テスト著者%d", i+1))

		price := book.CreateElement("price")
		price.CreateAttr("currency", "JPY")
		price.SetText(fmt.Sprintf("%d", 1000*(i+1)))
	}

	return doc
}

// createTestLibraryWithBooks テスト用図書館XMLを作成（具体的な書籍データで）
func createTestLibraryWithBooks(libraryName string, books []testBookData) *etree.Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	library := etree.NewElement("library")
	library.CreateAttr("name", libraryName)
	doc.SetRoot(library)

	// 指定された書籍データで書籍を作成
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