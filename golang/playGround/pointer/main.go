package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Function returning value
func createPersonValue() Person {
	return Person{Name: "Alice", Age: 25}
}

// Function returning pointer
func createPersonPointer() *Person {
	return &Person{Name: "Bob", Age: 30}
}

// Large struct example
type LargeStruct struct {
	Data [1000]int
	Name string
}

func createLargeValue() LargeStruct {
	return LargeStruct{Name: "Large Value"}
}

func createLargePointer() *LargeStruct {
	return &LargeStruct{Name: "Large Pointer"}
}

// ポインタ返却のメリット実演
func main() {
	fmt.Println("🔍 メリット1: メモリ効率")
	fmt.Println("------------------------")

	// 値渡し（大きなデータをコピー）
	largeValue := createLargeValue()
	fmt.Printf("📊 値渡しの場合: %d バイト分のデータをメモリにコピー\n", 4000+len("Large Value"))
	fmt.Printf("   値のアドレス: %p\n", &largeValue)

	// ポインタ渡し（アドレスのみ8バイト）
	largePtr := createLargePointer()
	fmt.Printf("📊 ポインタの場合: 8 バイト分のアドレスのみ\n")
	fmt.Printf("   実データのアドレス: %p\n", largePtr)
	fmt.Printf("   ポインタ変数のアドレス: %p\n", &largePtr)

	fmt.Println("\n🔗 メリット2: データ共有・変更可能")
	fmt.Println("--------------------------------")

	// 値型の場合（独立したコピー）
	fmt.Println("🔸 値型の場合:")
	person1 := createPersonValue()
	person1Copy := person1
	fmt.Printf("   元のデータ: %+v\n", person1)
	person1Copy.Age = 999
	fmt.Printf("   コピー変更後 - 元データ: %+v（変更されない）\n", person1)
	fmt.Printf("   コピー変更後 - コピー: %+v（変更された）\n", person1Copy)

	// ポインタ型の場合（同じデータを参照）
	fmt.Println("\n🔸 ポインタ型の場合:")
	person2 := createPersonPointer()
	person2Ref := person2
	fmt.Printf("   元のデータ: %+v\n", *person2)
	person2Ref.Age = 999
	fmt.Printf("   参照経由変更後 - 元データ: %+v（変更された！）\n", *person2)
	fmt.Printf("   参照経由変更後 - 参照: %+v（同じデータ）\n", *person2Ref)

	fmt.Println("\n❌ メリット3: nil チェック")
	fmt.Println("------------------------")

	// ポインタはnilを返せる（値型は不可）
	fmt.Println("🔸 値型: 空の値しか返せない")
	fmt.Println("🔸 ポインタ型: nilで「存在しない」を表現可能")

	var nilPerson *Person
	if nilPerson == nil {
		fmt.Println("   結果: データが見つかりません")
	}

	fmt.Println("\n🏭 メリット4: ファクトリーパターン")
	fmt.Println("------------------------------")

	// 条件によってnilまたはオブジェクトを返す
	fmt.Println("🔸 有効なIDで検索:")
	validPerson := findPerson("valid")
	if validPerson != nil {
		fmt.Printf("   結果: 見つかりました %+v\n", *validPerson)
	}

	fmt.Println("🔸 無効なIDで検索:")
	invalidPerson := findPerson("invalid")
	if invalidPerson == nil {
		fmt.Println("   結果: データが見つかりません")
	}

	fmt.Println("\n📝 まとめ:")
	fmt.Println("   ✅ 大きなデータでメモリ効率UP")
	fmt.Println("   ✅ 複数箇所での同じデータ共有")
	fmt.Println("   ✅ 「存在しない」状態を表現可能")
	fmt.Println("   ✅ 条件付きオブジェクト生成に便利")
}

// Return nil or Person based on condition
func findPerson(name string) *Person {
	if name == "valid" {
		return &Person{Name: "Found Person", Age: 35}
	}
	return nil // not found
}

// % go run main.go
// 🔍 メリット1: メモリ効率
// ------------------------
// 📊 値渡しの場合: 4011 バイト分のデータをメモリにコピー
//    値のアドレス: 0x140000c0008
// 📊 ポインタの場合: 8 バイト分のアドレスのみ
//    実データのアドレス: 0x140000c4008
//    ポインタ変数のアドレス: 0x1400009c038

// 🔗 メリット2: データ共有・変更可能
// --------------------------------
// 🔸 値型の場合:
//    元のデータ: {Name:Alice Age:25}
//    コピー変更後 - 元データ: {Name:Alice Age:25}（変更されない）
//    コピー変更後 - コピー: {Name:Alice Age:999}（変更された）

// 🔸 ポインタ型の場合:
//    元のデータ: {Name:Bob Age:30}
//    参照経由変更後 - 元データ: {Name:Bob Age:999}（変更された！）
//    参照経由変更後 - 参照: {Name:Bob Age:999}（同じデータ）

// ❌ メリット3: nil チェック
// ------------------------
// 🔸 値型: 空の値しか返せない
// 🔸 ポインタ型: nilで「存在しない」を表現可能
//    結果: データが見つかりません

// 🏭 メリット4: ファクトリーパターン
// ------------------------------
// 🔸 有効なIDで検索:
//    結果: 見つかりました {Name:Found Person Age:35}
// 🔸 無効なIDで検索:
//    結果: データが見つかりません

// 📝 まとめ:
//    ✅ 大きなデータでメモリ効率UP
//    ✅ 複数箇所での同じデータ共有
//    ✅ 「存在しない」状態を表現可能
//    ✅ 条件付きオブジェクト生成に便利
