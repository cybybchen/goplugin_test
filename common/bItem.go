package common

import (
	"fmt"
)

type BItem struct {
	*Item
}

func (this *BItem) APlus() {
	this.A += 100

	fmt.Println("version is bitem")
}
