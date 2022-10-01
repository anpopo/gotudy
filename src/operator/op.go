package main

import "fmt"

func main() {
	a := 4
	b := 5

	fmt.Printf("%v\n", a&b)
	fmt.Printf("%v\n", a|b)

	c := 5
	d := 7
	fmt.Printf("%v\n", c<<1)
	fmt.Printf("%v\n", d>>1)

	fmt.Printf("%v\n", d>>2)

	w := 3
	if w == 3 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
