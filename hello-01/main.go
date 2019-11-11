package main

import (
	"fmt"
	"strconv"
)

func main() {
	// step 1
	fmt.Println("hello world!")
	fmt.Println("===============")

	// step 2
	var s1 = [...]string{"hello", "world"}
	var s2 = [...]string{1: "world", 0: "hello"}
	var s3 = [...]string{"world", "hello", 2: "", "golang"}
	var s4 = [...]int{1, 2, 4: 5, 6}
	for _, i := range s1 {
		fmt.Print(i + "\t")
	}
	fmt.Println()
	fmt.Println("===============")
	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i] + "\t")
	}
	fmt.Println()
	fmt.Println("===============")
	for k, v := range s3 {
		fmt.Print(strconv.Itoa(k) + ":" + v + "\t")
	}
	fmt.Println()
	fmt.Println("===============")
	for k, v := range s4 {
		fmt.Print(strconv.Itoa(k) + ":" + strconv.Itoa(v) + "\t")
	}
	fmt.Println()
	fmt.Println("===============")

	//step 3
	var slice1 = []int{1, 2, 3}
	var slice2 = [4]int{1, 2, 3}
	slice1[3] = 4
	for k, v := range slice1 {
		fmt.Print(strconv.Itoa(k) + ":" + strconv.Itoa(v) + "\t")
	}
	slice2[3] = 4
	for k, v := range slice2 {
		fmt.Print(strconv.Itoa(k) + ":" + strconv.Itoa(v) + "\t")
	}
}
