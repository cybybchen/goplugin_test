package data

import "fmt"

type Item struct {
	A int32
}

func (this *Item) PrintVersion() {
	fmt.Println("version is 1")
}

func (this *Item) APlus() {
	this.A += 1

	fmt.Println("item version is 1")
}

func (this *Item) TestCb() func() {
	var i = 100

	return func() {
		fmt.Printf("i is %d\n", i)
	}
}
