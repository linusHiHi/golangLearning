package main

import (
	"fmt"
	"math"
	"math/rand"
)

// lv1
func sum(a, b int) int {
	result := a + b
	return result
}

// LV2
func area(r float64) float64 {
	return math.Pow(r, 2) * (math.Pi)
}

// judgePrimeNum 判断是否是素数，返回真值
func judgePrimeNum(num int) bool {
	var judge = true

	for i := 2; i < num && judge; i++ {
		judge = !(num%i == 0)
	}

	return judge
}

// LV x binSearch 用于二分法查找rand传来的数：low， high 是搜索范围， num是要查找的随机数， arr 是所查找的列表
// 尝试递归
func binSearch(low, high, num int, arr [100]int) int {
	mid := (high + low) / 2

	switch {
	case arr[mid] == num:
		return mid

	case arr[mid] < num:
		return binSearch(mid, high, num, arr)

	case arr[mid] > num:
		return binSearch(low, mid, num, arr)

	}
	return 0
}

func main() {
	randomNum := rand.Intn(100)
	var array [100]int
	for a := 0; a < 100; a++ {
		array[a] = a
	}

	fmt.Printf("Testing the func \"sum\"  233 + 123 = %d\n", sum(233, 123))
	fmt.Printf("Testing the func \"area\"  r = 12.0 => %f\n", area(12.0))
	fmt.Printf("Testing the func \"judgePrimeNum\"  233 %v\n", judgePrimeNum(233))
	fmt.Printf("Testing the func \"binSearch\"  %d %d\n", randomNum, binSearch(0, len(array)-1, randomNum, array))
}
