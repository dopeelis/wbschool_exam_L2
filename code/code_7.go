package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Создаем канал из поступивших числе
func asChan(vs ...int) chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

// Пытаемся объединить каналы
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		// Закрываем канал после всех операций
		defer close(c)
		// Пока в каналх что-то есть
		for a != nil || b != nil {
			select {
			// Пробуем читать из канала
			case v, ok := <-a:
				// Если не выходит, то ставим значение в 0, идем дальше
				if !ok {
					a = nil
					continue
				}
				// Если выходит, записываем в канал с
				c <- v
				// Тоже самое со вторым каналом
				// Проверяем, что там что-то есть
			case v, ok := <-b:
				// Если нет, ставим в 0, идем дальше
				if !ok {
					b = nil
					continue
				}
				// Если да, записываем в канал с
				c <- v
			}
		}
	}()

	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
