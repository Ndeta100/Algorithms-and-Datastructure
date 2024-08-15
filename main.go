package main

import "fmt"

func main() {
	println("Hello World")
	min, err := minVal(2, 2)
	if err != nil {
		println(err.Error())
	}
	println(min)
}

func minVal(a, b int) (int, error) {
	if a == b {
		fmt.Println("values are equal")
		return 0, nil
	}
	if a < b {
		return a, nil
	}
	return b, nil
}
