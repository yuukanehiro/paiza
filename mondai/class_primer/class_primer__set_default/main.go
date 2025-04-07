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
	for i := 1; i <= orderCount; i++ {
		orderInfo := nextLineStrings()
		customer := customersMap[atoi(orderInfo[0])]
		orderType := OrderType(orderInfo[1])

		// 価格はビールの場合は500円
		if orderType == Beer {
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
	}

	var customerKeys []int
	for k := range customersMap {
		customerKeys = append(customerKeys, k)
	}

	sortAsc(customerKeys)

	sales := map[int]int{}
	for _, v := range ordersMap {
		sales[v.Customer.ID] += v.Price + v.Discount
	}

	for _, customerID := range customerKeys {
		fmt.Println(sales[customerID])
	}
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
// 居酒屋で働きながらクラスの勉強をしていたあなたは、お客さんをクラスに見立てることで店内の情報を管理できることに気付きました。
// 全てのお客さんは、ソフトドリンクと食事を頼むことができます。加えて 20 歳以上のお客さんはお酒を頼むことができます。
// 20 歳未満のお客さんがお酒を頼もうとした場合はその注文は取り消されます。
// また、お酒（ビールを含む）を頼んだ場合、以降の全ての食事の注文 が毎回 200 円引きになります。

// 今回、この居酒屋でビールフェスをやることになり、ビールの注文が相次いだため、いちいちビールの値段である 500 円を書くのをやめ、伝票に注文の種類と値段を書く代わりに 0 とだけを書くことになりました。

// 店内の全てのお客さんの数と注文の回数、各注文をしたお客さんの番号とその内容が与えられるので、各お客さんの会計を求めてください。

// Input
// 3 5
// 19
// 43
// 22
// 2 0
// 2 food 4333
// 1 0
// 2 0
// 1 food 4606

// Output
// 4606
// 5133
// 0

// Input2
// 5 10
// 1
// 13
// 31
// 74
// 34
// 1 food 1088
// 4 alcohol 3210
// 1 alcohol 599
// 2 alcohol 602
// 2 softdrink 4375
// 4 food 1752
// 2 0
// 5 alcohol 4565
// 3 0
// 2 0

// Output2
// 1088
// 4375
// 500
// 4762
// 4565
