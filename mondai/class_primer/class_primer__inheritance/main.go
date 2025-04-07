package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type Customer struct {
	ID  int
	Old int
}

func NewCustomer(id int, old int) *Customer {
	return &Customer{
		ID:  id,
		Old: old,
	}
}

type OrderType string

const (
	Food    OrderType = "food"
	Alcohol OrderType = "alcohol"
)

type Order struct {
	Customer  *Customer
	OrderType OrderType
	Price     int
	Discount  int
}

func (o Order) totalPrice() int {
	return o.Price + o.Discount
}

// エラー定義
var ErrorChildOrderAlcohol error = errors.New("customer is not old enough to order alcohol")
var ErrorCustomerNil error = errors.New("customer is nil")

var isCustomerAlcoholOrder map[int]bool

const discountFoodAfterAlcoholOrder = -200

func NewOrder(customer *Customer, orderType OrderType, price int) (*Order, error) {
	if customer == nil {
		return nil, ErrorCustomerNil
	}

	if orderType == Alcohol && customer.Old < 20 {
		return nil, ErrorChildOrderAlcohol
	}

	if orderType == Alcohol {
		isAlcohol := true
		isCustomerAlcoholOrder[customer.ID] = isAlcohol
	}

	// アルコールの注文後は
	// 食べ物に対して割引を適用
	discount := 0
	if isCustomerAlcoholOrder[customer.ID] && orderType == Food {
		discount = discountFoodAfterAlcoholOrder
	}

	return &Order{
		Customer:  customer,
		OrderType: orderType,
		Price:     price,
		Discount:  discount,
	}, nil
}

func main() {
	isCustomerAlcoholOrder = make(map[int]bool)
	infoArray := nextLineWithoutEmptySpace("int").([]int)
	customerCount := infoArray[0]
	orderCount := infoArray[1]

	customers := map[int]*Customer{}
	for i := 1; i <= customerCount; i++ {
		customerInfo := nextLineWithoutEmptySpace("int").([]int)
		customers[i] = &Customer{}
		customers[i] = NewCustomer(i, customerInfo[0])
	}

	sales := map[string]int{}
	// 初期化
	for i := 1; i <= customerCount; i++ {
		sales[itoa(i)] = 0
	}

	for i := 1; i <= orderCount; i++ {
		orderInfo := nextLineWithoutEmptySpace("string").([]string)
		customerID, err := strconv.Atoi(orderInfo[0])
		if err != nil {
			fmt.Println("input error", err)
			return
		}

		orderType := OrderType(orderInfo[1])
		price, err := strconv.Atoi(orderInfo[2])
		if err != nil {
			fmt.Println("input error", err)
			return
		}

		order, err := NewOrder(customers[customerID], orderType, price)
		if err != nil {
			if errors.Is(err, ErrorChildOrderAlcohol) {
				continue
			} else {
				fmt.Println("error", err)
				return
			}
		}

		sales[itoa(customerID)] += order.totalPrice()
	}

	for i := 1; i <= customerCount; i++ {
		fmt.Println(sales[itoa(i)])
	}
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

// 行を取得してstringで返却
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 空白を除いた行を取得してinterface{}で返却
func nextLineWithoutEmptySpace(elementType string) interface{} {
	line := nextLine()

	var numberArray []string
	numberArray = strings.Fields(line) // 空白を除去して分割

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
// paiza 国の大衆居酒屋で働きながらクラスの勉強をしていたあなたは、お客さんをクラスに見立てることで店内の情報を管理できることに気付きました。
// 全てのお客さんは、ソフトドリンクと食事を頼むことができます。
// paiza 国の法律では、 20 歳以上のお客さんは成人とみなされ、お酒を頼むことができます。
// 20 歳未満のお客さんは未成年とみなされ、お酒を頼もうとした場合はその注文は取り消されます。
// また、お酒を頼んだ場合、以降の全ての食事の注文 が毎回 200 円引きになります.

// 店内の全てのお客さんの数と注文の回数、各注文をしたお客さんの番号とその内容が与えられるので、各お客さんの会計を求めてください。

// ヒント

// 注文について、20 歳未満のお客さんにできて、 20 歳以上のお客さんにできないことはないので、20歳未満のお客さんのクラスを作成して、それを継承して 20歳以上のお客さんのクラスを作成することで効率よく実装することができます。

// ▼　下記解答欄にコードを記入してみよう

// 入力される値
// N K
// a_1
// ...
// a_N
// n_1 s_1 m_1
// ...
// n_K s_K m_K

// ・ 1 行目では、お客さんの人数 N と注文の回数 K が与えられます。
// ・ 続く N 行のうち i 行目(1 ≦ i ≦ N)では、i 番目のお客さんの年齢が与えられます。
// ・ 続く K 行では、頼んだお客さんの番号 n_i , 注文の種類 s_i , 値段 m_i (1 ≦ i ≦ K) が与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。

// input
// 2 5
// 59
// 5
// 2 food 1223
// 1 alcohol 4461
// 1 alcohol 4573
// 1 alcohol 1438
// 2 softdrink 1581

// output
// 10472
// 2804

// input2
// 7 7
// 62
// 91
// 29
// 33
// 79
// 15
// 91
// 2 food 3134
// 7 alcohol 2181
// 6 softdrink 4631
// 3 softdrink 3120
// 4 softdrink 4004
// 6 alcohol 1468
// 6 alcohol 1245

// output2
// 0
// 3134
// 3120
// 4004
// 0
// 4631
// 2181
