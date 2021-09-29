// поведенчесткий паттерн
// передает запросы от компоненты  к компоненте
// определяет, нужно ли передавать дальше

// пояснение примера: имеется пространство, например флешка
// когда память на ней заканчивается, обработчики сообщают, на каком этапе
// в зависимости от отавшегося количества

package main

import (
	"fmt"
)

var maxSpace int = 64

type alarm interface {
	execute(*storage)
	setNext(alarm)
}

type storage struct {
	space int
}

// функция, имитирующая заполнение пространства
func (s *storage) takeSpace(gB int) {
	if gB > s.space {
		fmt.Println("Can't take space")
	} else {
		s.space -= gB
	}
}

// если пространство пустое (то есть не заполнено)
type fullStorage struct {
	next alarm
}

func (f *fullStorage) execute(s *storage) {
	if s.space == maxSpace {
		fmt.Println("Storage is empty")
		fmt.Println("Storage space:", s.space)
		return
	}
	f.next.execute(s) // если не пустое, то передает дальше
}

func (f *fullStorage) setNext(next alarm) {
	f.next = next
}

// если заполнено наполовину или больше
type halfStorage struct {
	next alarm
}

func (h *halfStorage) execute(s *storage) {
	if s.space >= maxSpace/2 {
		fmt.Println("Storage have more then half space")
		fmt.Println("Storage space:", s.space)
		return
	}
	if s.space < maxSpace/2 {
		if s.space >= maxSpace/4 {
			fmt.Println("Storage have less then half space!")
			fmt.Println("Storage space:", s.space)
			return
		} else {
			h.next.execute(s) // если меньше половины, то передает дальше
		}
	}
}

func (h *halfStorage) setNext(next alarm) {
	h.next = next
}

// если свободно четверть
type quarterStorage struct {
	next alarm
}

func (q *quarterStorage) execute(s *storage) {
	if s.space < maxSpace/4 {
		fmt.Println("Storage have less then quarter space!")
		fmt.Println("Storage space:", s.space)
		return
	}
	if s.space == 0 {
		fmt.Println("Storage is full!") // т.к. следующего обработчика нет, то передачи дальше не будет
		// просто объявляем, что заполнено
	}
}

func (q *quarterStorage) setNext(next alarm) {
	q.next = next
}

func main() {
	q := &quarterStorage{}

	h := &halfStorage{}
	h.setNext(q)

	f := &fullStorage{}
	f.setNext(h)

	flashDrive := &storage{space: maxSpace}
	// имитируем занимание пространсва
	flashDrive.takeSpace(38)
	f.execute(flashDrive)

	// и еще раз
	flashDrive.takeSpace(29)
	f.execute(flashDrive)

}
