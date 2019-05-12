package main

import (
	"fmt"
)

func foo(c chan int, value int) {
	c <- value * 5
}

func main() {
	fooValue := make(chan int)
	for i := 0; i < 10; i++ {
		go foo(fooValue, i)
	}

	for item := range fooValue {
		fmt.Println(item)
	}

}
