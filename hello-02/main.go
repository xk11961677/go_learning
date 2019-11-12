/**
函数
*/
package main

import (
	"fmt"
	"strconv"
)

// more return value
func swap(a int, b int) (int, int) {
	return b, a + b
}

// slice type params
func slice(a int, more ...int) {
	fmt.Print(a)
	for _, i := range more {
		fmt.Println(i)
	}
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func Find(m map[int]string, key int) (value string, ok bool) {
	value, ok = m[key]
	return
}

//defer 延迟调用,先执行①后执行②
func Inc(v int) int {
	defer func() {
		v++
		fmt.Println("===②" + strconv.Itoa(v))
	}()
	fmt.Println("======①" + strconv.Itoa(v))
	return 42
}

// go语言可动态变更栈大小， 所以几乎没有栈溢出
// go语言  & * 符号 对于栈和堆不太能确定某个变量到底在栈或堆中
// & 取址符号, 返回 x 变量的地址
func f(x int) *int {
	return &x
}

// * 指针符号，返回 x 指针(引用)
func g() int {
	var x = new(int)
	return *x
}

func main() {
	//func
	var a, b = swap(1, 2)
	fmt.Println(strconv.Itoa(a) + "\t" + strconv.Itoa(b))
	var more = []int{1, 2, 3, 4}
	// more... unpack
	slice(1, more...)

	var inter = []interface{}{123, "abc"}
	Print(inter...)
	Print(inter)
	var m = map[int]string{1: "123", 2: "456"}
	var value, ok = Find(m, 1)
	fmt.Println(value + "\t" + strconv.FormatBool(ok))

	var f = 40
	var k = Inc(f)
	fmt.Println(k)

}
