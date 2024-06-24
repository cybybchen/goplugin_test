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

func ModifyItem() {
	var bItem = &common.BItem{}
	bItem.Item = common.GetItem().(*common.Item)
	common.SetItem(bItem)
}

func init() {
	fmt.Println("init item2")
}
