package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := []int{}
	for i := 0; i < 11; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
	for num, _ := range x {
		if num%2 == 0 {
			fmt.Println(strconv.Itoa(num) + " even")
		} else {
			fmt.Println(strconv.Itoa(num) + " odd")
		}
	}
}
