package main

import "fmt"

func main() {
	t := 55

	for t < 60 {
		fmt.Println(t)
		t++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("this is i = ", i)
	}
}
