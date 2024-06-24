package main

import (
	"common"
	"fmt"
)

var count int32

func HcalcItem() {
	fmt.Println(222)

	count++
	fmt.Printf("count is %d\n", count)
}

func ItemModify() {
	item := common.GetItem()
	item.(*common.Item).A = 10000
}

func init() {
	fmt.Println("init item")
}
