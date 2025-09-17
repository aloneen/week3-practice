package main

import (
	"fmt"
	"github.com/aloneen/week3-practice/tree/main/mod1/greeting"
	"github.com/aloneen/week3-practice/tree/main/mod2/mathutil"
)

func main() {
	fmt.Println(greeting.Hello("Dias"))
	fmt.Println("2 + 3 =", mathutil.Add(2, 3))
	fmt.Println("5 * 6 =", mathutil.Mul(5, 6))
}
