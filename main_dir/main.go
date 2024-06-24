package main

import (
	"common"
	"fmt"
	"plugin"
)

var testCb func()

func main() {
	var item = common.GetItem()
	fmt.Printf("1111item point is %p\n", item)
	fmt.Printf("1111item is %+v\n", item)
	//for {
	//	run()
	//}

	testCb = item.TestCb()

	run()
}

func run() {
	//time.Sleep(time.Second * 5)

	testGoPlugin1()

	testGoPlugin2()
}

func testGoPlugin1() {
	p, err := plugin.Open("test1.so")
	if err != nil {
		fmt.Println("Error opening plugin:", err)
		return
	}

	//fc, err := p.Lookup("HcalcItem")
	//if err != nil {
	//	fmt.Println("Error looking up Func1:", err)
	//	return
	//}
	//
	//cI, ok := fc.(func())
	//if !ok {
	//	fmt.Println("Error type assertion")
	//}
	//cI()

	//fc, err := p.Lookup("GetItem")
	//if err != nil {
	//	fmt.Println("Error looking up Func1:", err)
	//	return
	//}
	//
	//cI1, ok := fc.(func() *data.Item)
	//if !ok {
	//	fmt.Println("Error type assertion")
	//}
	//
	//item = cI1()

	fc, err := p.Lookup("ItemModify")
	if err != nil {
		fmt.Println("Error looking up Func1:", err)
		return
	}

	cI1, ok := fc.(func())
	if !ok {
		fmt.Println("Error type assertion")
	}

	cI1()
	var item = common.GetItem()
	item.APlus()

	fmt.Printf("item point is %p\n", item)
	fmt.Printf("item is %+v\n", item)
}

func testGoPlugin2() {
	p, err := plugin.Open("test2.so")
	if err != nil {
		fmt.Println("Error opening plugin:", err)
		return
	}

	fc, err := p.Lookup("ModifyItem")
	if err != nil {
		fmt.Println("Error looking up Func1:", err)
		return
	}

	cI1, ok := fc.(func())
	if !ok {
		fmt.Println("Error type assertion")
	}

	cI1()
	var item = common.GetItem()
	item.APlus()

	testCb()

	fmt.Printf("item point is %p\n", item)
	fmt.Printf("item is %+v\n", item)
}
