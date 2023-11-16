package main

import "fmt"

// Manager 管理接口
type Manager interface {
	Checker() int
	ReNewer(int)
	Printer()
}

// DigitalManager 电子产品的接口
type DigitalManager interface {
	Manager
	TypePrinter()
	BrandPrinter()
}

// Goods 商品结构体和方法
type Goods struct {
	name   string
	price  int
	amount int
}

// Checker 使用指针接收者是为了同reNew（刷新）保持一致
func (g *Goods) Checker() int {
	return g.amount
}

// ReNewer 刷新用的函数
func (g *Goods) ReNewer(newAmount int) {
	g.amount = newAmount
}

// Printer 打印名称和储量
func (g *Goods) Printer() {
	fmt.Printf("%v的储量为：%d\n", g.name, g.amount)
}

// DigitalGoods 电子产品子结构体
type DigitalGoods struct {
	Goods
	brand   string
	typeStr string
}

// Printer 方法重写，子类覆盖父类
func (d *DigitalGoods) Printer() {
	fmt.Printf("%v的储量为：%d,爱来自digital\n", d.name, d.amount)
}

func (d *DigitalGoods) TypePrinter() {
	fmt.Printf("its type is \"%v\"\n", d.typeStr)
}

func (d *DigitalGoods) BrandPrinter() {
	fmt.Printf("its Brand is \"%v\"\n", d.brand)
}

//调用接口的函数们
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

func typePrint(obj DigitalManager) {
	obj.TypePrinter()
}

func main() {
	//测试初始化
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
	//测试各函数
	printFuncDig(&dm)
	printFunc(&good)
	fmt.Printf("%d\n", check(&dm))
	fmt.Printf("%d\n", check(&good))
	reNew(&dm, 233)
	//测试digital manager的有效性
	printFuncDig(&dm)
	typePrint(&dm)
}
