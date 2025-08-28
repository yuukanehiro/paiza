package main

import (
	"errors"
	"fmt"
)

// 関数名	用途
// errors.Is	エラーが特定のエラーと一致しているか（チェーンの中にあるか）を確認
// errors.As	特定のエラー型にキャストできるかどうかを確認し、その型のポインタに代入

type MyCustomError struct {
	Msg string
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("MyCustomError: %s", e.Msg)
}

var (
	ErrNotFound = errors.New("not found")
)

func returnWrappedError() error {
	// ErrNotFound を wrap したエラーを返す
	return fmt.Errorf("additional context: %w", ErrNotFound)
}

func returnCustomError() error {
	// 独自エラー型を返す
	return &MyCustomError{Msg: "something went wrong"}
}

func returnSimpleError() error {
	return errors.New("simple error")
}

func main() {
	// ----- errors.Is の使用例 -----
	err1 := returnWrappedError()
	if errors.Is(err1, ErrNotFound) {
		fmt.Println("errors.Is: err1 is ErrNotFound") // ✅ ここは true
	} else {
		fmt.Println("errors.Is: err1 is not ErrNotFound")
	}

	// ----- errors.As の使用例 -----
	err2 := returnCustomError()
	var myErr *MyCustomError
	if errors.As(err2, &myErr) {
		fmt.Println("errors.As: err2 is of type MyCustomError") // ✅ ここは true
		fmt.Println("Message:", myErr.Msg)
	} else {
		fmt.Println("errors.As: err2 is not of type MyCustomError")
	}

	// ----- 通常のエラー比較 -----
	err3 := returnSimpleError()
	var myErr2 *MyCustomError
	if errors.As(err3, &myErr2) {
		fmt.Println("errors.As: err3 is of type MyCustomError")
	} else {
		fmt.Println("errors.As: err3 is not of type MyCustomError") // ✅ ここは false
	}
}
