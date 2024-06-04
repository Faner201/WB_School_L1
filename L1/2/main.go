package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}

	chanel(numbers)
}

func chanel(numbers [5]int) {
	// создаем канал для успешной синхронизации и блокировки горутины
	ch := make(chan int)
	// итерируемся по нашему массиву и добавляем в канал квадраты наших значений, за счет горутин
	for _, number := range numbers {
		go func(i int) {
			ch <- i * i
		}(number)
	}
	// первый способ вывести информацию из канала

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%v ", <-ch)
	}

	// второй способ

	// var i int
	// for v := range ch {
	// 	i += 1
	// 	fmt.Printf("%d ", v)
	// 	if i == 5 {
	// 		close(ch)
	// 	}
	// }

	// способы конфликтуют между собой
}

func waitGroup(numbers [5]int) {
	var wg sync.WaitGroup

	// данные подход может получить ошибку при неожиданном создание n+1 горутин

	// for _, number := range numbers {
	// 	wg.Add(1)
	// 	go func (i int)  {
	// 		defer wg.Done()
	// 		fmt.Printf("%v ", i * i)
	// 	}(number)
	// }

	// wg.Wait()

	wg.Add(len(numbers))
	for _, number := range numbers {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%v ", i*i)
		}(number)

		wg.Wait()
	}
}
