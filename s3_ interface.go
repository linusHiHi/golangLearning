package main

import "fmt"

type Manager interface {
	Checker() int
	ReNewer(int)
	Printer()
}

type DigitalManager interface {
	Manager
	TypePrinter()
	BrandPrinter()
}

type Goods struct {
	name   string
	price  int
	amount int
}

func (g *Goods) Checker() int {
	return g.amount
}

func (g *Goods) ReNewer(newAmount int) {
	g.amount = newAmount
}

func (g *Goods) Printer() {
	fmt.Printf("%v的储量为：%d\n", g.name, g.amount)
}

type DigitalGoods struct {
	Goods
	brand   string
	typeStr string
}

// Printer 方法重写
func (d *DigitalGoods) Printer() {
	fmt.Printf("%v的储量为：%d,爱来自digital\n", d.name, d.amount)
}

func (d *DigitalGoods) TypePrinter() {
	fmt.Printf("its type is \"%v\"\n", d.typeStr)
}

func (d *DigitalGoods) BrandPrinter() {
	fmt.Printf("its Brand is \"%v\"\n", d.brand)
}

func check(m Manager) int {
	return m.Checker()
}

func reNew(m Manager, data int) {
	m.ReNewer(data)
}

func printFunc(m Manager) {
	m.Printer()
}

// 看看digital manager满足接口吗
func printFuncDig(m DigitalManager) {
	m.Printer()
}

func main() {
	good := Goods{
		name:   "翻译器",
		price:  12,
		amount: 12,
	}
	var dm DigitalGoods
	dm = DigitalGoods{
		Goods: Goods{
			name:   "移动电话",
			price:  12,
			amount: 12,
		},
		brand:   "apple",
		typeStr: "14pro max",
	}
	printFuncDig(&dm)
	printFunc(&good)
	fmt.Printf("%d\n", check(&dm))
	fmt.Printf("%d\n", check(&good))
}
