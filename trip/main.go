package main

import (
	"fmt"
	"trip/internal/domain"
)

func main() {
	x := domain.DefaultCapacity()
	fmt.Println(x)
	if err := domain.ValidateCapacity(x); err != nil {
		fmt.Println(err)
	}
}
