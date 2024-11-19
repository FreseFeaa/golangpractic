package main

import "fmt"

func main() {
	fmt.Println("Sigma channel Deadlock")
	Channel := make(chan int)

	<-Channel

}
