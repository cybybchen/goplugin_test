package common

var intance ItemInf = &Item{
	A: 1111,
}

func SetItem(item ItemInf) {
	intance = item
}

func GetItem() ItemInf {
	return intance
}
