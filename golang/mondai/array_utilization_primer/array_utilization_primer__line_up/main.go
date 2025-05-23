package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	infoArray := nextLineBySeparator(" ", "int").([]int)
	sheetCount := infoArray[1]
	delSheetCount := infoArray[2]

	var sheetArray []int
	for i := 0; i < sheetCount; i++ {
		str := nextLine()
		n, _ := strconv.Atoi(str)
		sheetArray = append(sheetArray, n)
	}

	if len(sheetArray) < delSheetCount {
		fmt.Println("Error: sheetArray length is less than delSheetCount")
		return
	}

	deledSheetArray := sheetArray[delSheetCount:]

	if len(deledSheetArray) == 0 {
		fmt.Printf("%d\n", 0)
		return
	}

	seen := make(map[int]bool)
	var res []int
	for _, v := range deledSheetArray {
		if !seen[v] {
			seen[v] = true
			res = append(res, v)
		}
	}

	for _, v := range res {
		fmt.Printf("%d\n", v)
	}
}

// 行を取得してstringで返却
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 行を取得してinterface{}で返却
// 注意: interface{}で返却されるので、キャストを忘れないこと
// 利用例
// ・array := nextLineBySeparator(" ", "string").([]string) // []stringで取得
// ・array := nextLineBySeparator(" ", "int").([]int)　// []intで取得
func nextLineBySeparator(separator string, elementType string) interface{} {
	line := nextLine()

	var numberArray []string
	numberArray = strings.Split(line, separator)

	if len(numberArray) == 0 {
		return nil
	}

    if elementType == "string" {
        return numberArray
    } else if elementType == "int" {

		var numbers []int
		for _, v := range numberArray {
			number, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("input error", err)
				return nil
			}

			numbers = append(numbers, number)
		}

		return numbers
	} else {
		fmt.Println("elementType error")
		return nil
	}
}

// Q
// 開店直後に店に入るために、番号 1 〜 N の N 人が開店前に店の前にブルーシートを合計 K 枚置きました。ブルーシートの先頭は A_1 , 最後尾は A_K です。しかし、店側は先頭から F 枚のブルーシートを撤去しました。

// 1 〜 N 番の人々は次のルールに従って店に並びます。
// ・自分のブルーシートが 1 枚以上残っているとき、自分のブルーシートのうち、最も先頭に近いブルーシートの位置に並ぶ。
// ・自分のブルーシートが 1 枚も残っていないとき、並ぶことなく帰宅する。

// 全員が並び終わった後に待機列にいる人の番号を先頭から順に答えてください。

// 例
// ・N = 3, K = 6, F = 3, A = [3, 2, 1, 2, 3, 2] のとき
// 撤去が行われる前と行われた後のブルーシートの様子は次の通りになります。

// 番号 1 の人のブルーシートは 1 枚も残っていないので並ぶことなく帰宅します。
// 番号 2 の人のブルーシートは 2 枚残っているので、最も先頭に近いブルーシートの位置に並びます。
// 番号 3 の人のブルーシートは 1 枚残っているので、そのブルーシートの位置に並びます。

// 最終的な待機列は次の通りになるので、期待する出力は

// 2
// 3

// Input
// 5 10 1
// 1
// 4
// 4
// 3
// 5
// 4
// 2
// 4
// 1
// 1

// Output
// 4
// 3
// 5
// 2
// 1
