/**
1. 方法: 面向对象语言所称
2. go语言通过组合方式实现继承(编译器完成,非运行期完成,导致基类中this就是基类的this,而不会是多态类型)
*/
package main

import (
	"fmt"
	"image/color"
	"sync"
)

type Point struct {
	X, Y float64
}

// ColorPoint 继承 Point
type ColorPoint struct {
	Point
	Color color.RGBA
}

/***************************************/
// p.Lock 和 p.Unlock 在编译器会展开为 p.Mutex.lock() 和 p.Mutex.unlock()
type Cache struct {
	m map[string]string
	sync.Mutex
}

func (p *Cache) Lookup(key string) string {
	p.Lock()
	defer p.Unlock()

	return p.m[key]
}

func main() {
	var c ColorPoint
	c.Point.X = 12
	c.Point.Y = 20
	c.X = 13
	fmt.Println(c.X)

	var m = map[string]string{"key": "value"}
	var cache = new(Cache)
	cache.m = m
	lookup := cache.Lookup("key")
	fmt.Println(lookup)
}
