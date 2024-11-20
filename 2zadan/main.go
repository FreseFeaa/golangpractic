// ## Задание 2
// Записать в канал 10 числовых значений в потоке основной горутины.
// Запустить чтение из канала внутри других 10 горутин.
// Каждая горутина должна преобразовывать числовое значение в строку и складывать во второй канал.
// Основная горутина должна читать из второго канала строки и выводить их на экран
// в порядке поступления из канала.
// Строчка с запуском горутины в коде должна быть одна (go ...)
// В результате в консоль должно распечататься 10 чисел в строковом представлении.

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	channlNum := make(chan int, 10)
	channlStr := make(chan string, 10)

	var waitgrp sync.WaitGroup

	for i := 1; i <= 10; i++ {
		channlNum <- i
	}
	close(channlNum)

	for i := 0; i < 10; i++ {
		waitgrp.Add(1)
		go func() {
			defer waitgrp.Done()
			num := <-channlNum
			str := strconv.Itoa(num)
			channlStr <- str
		}()
	}

	waitgrp.Wait()
	close(channlStr)

	fmt.Println("Итоговый канал строк: ")
	for stroka := range channlStr {
		fmt.Println("Число в строке: ", stroka)
	}
}
