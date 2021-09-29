// поведенческий паттерн
// позволяет выбрать поведения алгоритма в ходе исполнения
// похоже на фабричный метод

// пояснение примера: имеется интерфейс данных iData с соответствующими методами
// в зависимости от типа данных, отправка осуществляется по-разному

package main

import "fmt"

// опеделяем методы для данных
type iData interface {
	send()
	read()
	delete()
	getDelails()
	setName(string)
	setContent(string)
}

// определяем параметры данных
type data struct {
	dataType string
	name     string
	size     int
	content  string
}

// реализуем методы для тип данных data
func (d *data) send() {
	fmt.Println("Data has been sent to storage")
}

func (d *data) read() {
	fmt.Println("Content:", d.content)
}

func (d *data) delete() {
	fmt.Println("The data has been deleted")
}

func (d *data) getDelails() {
	fmt.Println("\nDataType:", d.dataType)
	fmt.Println("Name:", d.name)
	fmt.Println("Size:", d.size)
	fmt.Println("Content:", d.content)
}

func (d *data) setName(name string) {
	d.name = name
}

func (d *data) setContent(c string) {
	d.content = c
	d.size += 1
}

// создаем подтип данных Российские даннные
type rusData struct {
	data
}

func newRusData() iData {
	return &rusData{data: data{dataType: "Russia", name: "Noname", size: 0, content: "Empty"}}
}

// переопределяем для него метод send
func (rusD *rusData) send() {
	fmt.Println("Data has been sent to storage 'Yandex'")
}

// создаем подтип данных Американские даннные
type americanData struct {
	data
}

func newAmericanData() iData {
	return &americanData{data: data{dataType: "USA", name: "Noname", size: 0, content: "Empty"}}
}

// переопределяем для него метод send
func (aData *americanData) send() {
	fmt.Println("Data has been sent to storage 'Amazon'")
}

// создаем новый экземпляр в зависимости от типа
func newData(dataType string) iData {
	if dataType == "Russia" {
		return newRusData()
	}
	if dataType == "USA" {
		return newAmericanData()
	}
	return nil
}

func main() {
	// создаем нужные экземпляры
	d1 := newData("Russia")
	d2 := newData("USA")

	// выводим данные до изменений
	d1.getDelails()
	d2.getDelails()

	// отправляем d2
	fmt.Print("American data: ")
	d2.send()

	// меняем дефолтные значения для d1
	d1.setName("Письмо")
	d1.setContent("Это - письмо")
	fmt.Print("Russian data: ")
	// отправляем d1
	d1.send()
	d1.getDelails()
}
