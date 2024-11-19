package main

import (
	"fmt"
	"strconv"
)

func main() {
	numChannel := make(chan int)
	strChannel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			numChannel <- i
		}
		close(numChannel)
	}()

	for i := 0; i < 10; i++ {
		go func() {
			num := <-numChannel
			strChannel <- strconv.Itoa(num)
		}()
	}

	fmt.Println("Итоговый канал строк: ")
	for i := 0; i < 10; i++ {
		fmt.Println("Число в строке:", <-strChannel)
	}
}
