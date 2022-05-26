package main

import (
	"fmt"
	"github.com/phobiadev/goitertools"
)

func main() {
	for n := range goitertools.Count(1,10) {
		if n == 31 {
			break
		}
		fmt.Println(n)
	}

}