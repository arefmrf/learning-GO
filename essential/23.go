package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a, len(a), cap(a))

	b := a[2:]
	fmt.Println(b, len(b), cap(b))

	c := a[:3]
	fmt.Println(c, len(c), cap(c))

	d := c[0:5]
	fmt.Println(d, len(d), cap(d))
}
