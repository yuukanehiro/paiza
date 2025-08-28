package main

import (
	"errors"
	"fmt"
)

// errors.Is() と errors.As() の違い：一言でまとめると
// 関数	比較するもの	用途
// errors.Is	値（エラー定数）	エラーが特定のエラー（sentinel error）に一致するかを確認
// errors.As	型（エラー構造体）	エラーが特定の型に変換できるかを確認（フィールドにアクセスしたいとき）

// 定数エラー（sentinel error）
var ErrNotFound = errors.New("not found")

// 構造体エラー（型としてのエラー）
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Msg)
}

func returnWrappedSentinel() error {
	return fmt.Errorf("wrap: %w", ErrNotFound)
}

func returnWrappedStructError() error {
	return fmt.Errorf("wrap: %w", &MyError{Code: 404, Msg: "resource missing"})
}

func main() {
	// ------- Isの例（定数比較）-------
	err1 := returnWrappedSentinel()
	if errors.Is(err1, ErrNotFound) {
		fmt.Println("errors.Is: err1 is ErrNotFound ✅") // ここは true
	}

	// ------- Asの例（型キャスト）-------
	err2 := returnWrappedStructError()
	var myErr *MyError
	if errors.As(err2, &myErr) {
		fmt.Println("errors.As: err2 is MyError ✅") // ここは true
		fmt.Printf("Code=%d, Msg=%s\n", myErr.Code, myErr.Msg)
	}

	// ------- AsではErrNotFoundはダメ -------
	var notFound *MyError
	if errors.As(ErrNotFound, &notFound) {
		fmt.Println("これは出ない ❌")
	} else {
		fmt.Println("errors.As: ErrNotFound is NOT MyError ❌") // true
	}
}
