package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

const Food OrderType = "food"
const Alcohol OrderType = "alcohol"
const SoftDrink OrderType = "softdrink"
const Beer OrderType = "0"
const Amount OrderType = "A"

type Order struct {
	ID        int
	Customer  Customer
	OrderType OrderType
	Price     int
	Discount  int
}

func NewOrder(id int, customer Customer, orderType OrderType, price int) *Order {
	var discount int
	if orderType == Alcohol || orderType == Beer {
		if customer.Old < 20 {
			return nil
		}

		if !isCustomerAlcoholOrder[customer.ID] {
			isCustomerAlcoholOrder[customer.ID] = true
		}
	}

	if orderType == Food && isCustomerAlcoholOrder[customer.ID] {
		discount = -200
	}

	return &Order{
		ID:        id,
		Customer:  customer,
		OrderType: orderType,
		Price:     price,
		Discount:  discount,
	}
}

var isCustomerAlcoholOrder map[int]bool

func main() {
	isCustomerAlcoholOrder = make(map[int]bool)

	infoArray := nextLineInts()
	customerCount := infoArray[0]
	orderCount := infoArray[1]

	customersMap := make(map[int]*Customer)
	for i := 1; i <= customerCount; i++ {
		customerInfo := nextLineInts()
		c := NewCustomer(i, customerInfo[0])
		if c == nil {
			continue
		}
		customersMap[i] = &Customer{}
		customersMap[i] = c
	}

	ordersMap := make(map[int]*Order)
	price := 0
	sales := map[int]int{}
	amountCount := 0
	for i := 1; i <= orderCount; i++ {
		orderInfo := nextLineStrings()
		customer := customersMap[atoi(orderInfo[0])]
		orderType := OrderType(orderInfo[1])

		// 価格はビールの場合は500円
		if orderType == Amount {
			// 退店時の売上を出力
			fmt.Println(sales[atoi(orderInfo[0])])
			amountCount++
			continue
		} else if orderType == Beer {
			price = 500
		} else {
			price = atoi(orderInfo[2])
		}

		order := NewOrder(i, *customer, orderType, price)
		if order == nil {
			continue
		}

		ordersMap[i] = &Order{}
		ordersMap[i] = order

		sales[customer.ID] += order.Price + order.Discount
	}

	// 退店人数を出力
	fmt.Println(amountCount)
}

func sortAsc(numbers []int) {
	sort.Ints(numbers)
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}

	return i
}

// 行を取得してstringで返却
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextLineInts() []int {
	line := nextLine()
	parts := strings.Fields(line)
	var nums []int
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func nextLineStrings() []string {
	line := nextLine()
	return strings.Fields(line)
}

// Q
// 居酒屋で働きながらクラスの勉強をしていたあなたは、お客さんをクラスに見立てることで勤務時間中の店内の人数や注文の情報を管理できることに気付きました。
// 全てのお客さんは、ソフトドリンクと食事を頼むことができます。加えて 20 歳以上のお客さんはお酒を頼むことができます。
// 20 歳未満のお客さんがお酒を頼もうとした場合はその注文は取り消されます。
// また、お酒（ビールを含む）を頼んだ場合、以降の全ての食事の注文 が毎回 200 円引きになります。

// 今回、この居酒屋でビールフェスをやることになり、ビールの注文が相次いだため、いちいちビールの値段である 500 円を書くのをやめ、注文の種類と値段を書く代わりに 0 とだけを書くことになりました。

// 勤務時間の初めに店内にいるお客さんの人数と与えられる入力の回数、各注文をしたお客さんの番号とその内容、または退店したお客さんの番号が与えられます。
// お客さんが退店する場合はそのお客さんの会計を出力してください。勤務時間中に退店した全てのお客さんの会計を出力したのち、勤務時間中に退店した客の人数を出力してください。

// Input
// 2 3
// 20
// 30
// 1 0
// 2 0
// 1 A

// Output
// 500
// 1

// Input2
// 7 12
// 68
// 85
// 57
// 32
// 90
// 74
// 7
// 2 0
// 4 A
// 3 0
// 1 A
// 4 softdrink 3781
// 6 softdrink 3010
// 4 0
// 5 alcohol 1018
// 1 0
// 1 softdrink 376
// 1 softdrink 797
// 2 alcohol 4284

// Output2
// 0
// 0
// 2
