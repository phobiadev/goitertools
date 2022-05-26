package main

import (
	"fmt"
	"github.com/phobiadev/goitertools"
)

func main() {
	a1 := []float64{1.1,1.1,1.1}
	a2 := []int{1,2,3}

	fmt.Println(goitertools.AccumulateAdd(a1,0))
	fmt.Println(goitertools.AccumulateAdd(a2,0))

	a3 := []string{"a","b","c","d"}
	fmt.Println(goitertools.Accumulate(a3,func(a,b string) string {return fmt.Sprintf("%v%v",a,b)},""))
}