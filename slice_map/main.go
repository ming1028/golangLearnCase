package main

import (
	"fmt"
	"github.com/samber/lo"
)

type Order struct {
	OrderNo string
	Uid     int64
}

func main() {
	orders := []*Order{
		&Order{"123213", 65},
		&Order{"33222", 55},
	}
	// 筛选
	lo.Filter(orders, func(item *Order, index int) bool {
		return item.OrderNo == "123"
	})

	// 获取某一个列
	orderNos := lo.Map(orders, func(item *Order, index int) string {
		return item.OrderNo
	})
	fmt.Println(orderNos)

	// slice => map
	orderMapUid := lo.SliceToMap(orders, func(item *Order) (int64, *Order) {
		return item.Uid, item
	})
	fmt.Println(orderMapUid)

	// group by uid
	orderGroupByUid := lo.GroupBy(orders, func(item *Order) int64 {
		return item.Uid
	})
	fmt.Println(orderGroupByUid)

}
