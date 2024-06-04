package main

import (
	"fmt"
	"sync"
)

/*
	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

*/

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}

	channel(numbers)
	waitGroup(numbers)
}

func channel(numbers [5]int) {
	// 1 способ
	// ch := make(chan int)
	// var sum int
	// for _, number := range numbers {
	// 	go func(i int) {
	// 		ch <- i * i
	// 	}(number)
	// }
	// for i := range ch {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// 2 способ
	ch := make(chan int)
	var sum int

	// создаем дополнительный канал, для отслеживание всех горутин

	done := make(chan bool)

	for _, number := range numbers {
		go func(i int) {
			ch <- i * i
			done <- true
		}(number)
	}

	go func() {
		for range numbers {
			<-done // Ожидаем сигналы от всех горутин
		}
		close(ch)
	}()

	for i := range ch {
		sum += i
	}
	fmt.Println(sum)

}

func waitGroup(numbers [5]int) {
	// создаем группу для отслеживание горутин
	var wg sync.WaitGroup
	// говорим за сколькими горутинами следить
	wg.Add(len(numbers))

	ch := make(chan int)

	var sum int

	for _, number := range numbers {
		go func(i int) {
			defer wg.Done()
			ch <- i * i
		}(number)
	}

	go func() {
		wg.Wait() // ожидаем все горутины и закрываем канал
		close(ch)
	}()

	for i := range ch {
		sum += i
	}
	fmt.Println(sum)
}
