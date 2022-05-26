package main

import (
	"fmt"
	"github.com/phobiadev/goitertools"
)

func test[T any](c []T, br int) {
	p := 0
	for n := range goitertools.Cycle(c) {
		if p == br {
			break;
		} 
		fmt.Println(n)
		p++
	}
}

func main() {
	a := []int{1,2,3}
	s := []string{"t","e","s","t"}
	test(a,6)
	fmt.Println()
	test(s,8)
}