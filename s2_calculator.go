package main

import "fmt"

/*
this is a simple calculator,which allow you to type in your simple expression in.

*/

type Calculator func(int, int) int

//it is using to process information from keyboard
//option is for recognizing + _ * /

func ProcessIncome(option string) Calculator {
	funcMap := map[string]Calculator{
		"add":     Add,
		"minus":   Minus,
		"product": Product,
		"divide":  Divide,
	}
	//former function
	/*
		switch option {
		case "add":
			return computer(num1, num2, Add)
		case "minus":
			return computer(num1, num2, Minus)
		case "product":
			return computer(num1, num2, Product)
		case "divide":
			return computer(num1, num2, Divide)
		}
	*/
	return funcMap[option]
}

// 用于调用具体计算函数
func computer(a int, b int, function Calculator) int {
	return function(a, b)
}

func Add(a int, b int) int {
	return b + a
}

func Minus(a, b int) int {
	return a - b
}

func Product(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}

func main() {
	var sig string
	var num1, num2 int
	var result int

	fmt.Printf("Good morning,Good evening,and good night!\n" +
		"you computor is running!\n" +
		"type in you expression,plz!\n" +
		"e.g. 12 , 15 , Add(minus,product,divide)\n")

	scan, err := fmt.Scanf(`%d , %d , %s`, &num1, &num2, &sig)

	if err != nil {
		funcUsing := ProcessIncome(sig)
		result = computer(num1, num2, funcUsing)
		fmt.Printf("\nanswer is %d\n%v", result, scan)
	} else {
		print("Error!")
	}

}
