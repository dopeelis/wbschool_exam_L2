// порождающий метод
// позволяет подклассам изменять тип создаваемых объектов
// интерфейс которых определен в суперклассе
// в Go можно реализовать "простую фабрику"

// пояснение примера: имеется интерфейс iToy, описывающий "поведение" игрушек разных видов
// также имеется два вида ирушек: кукла и лего
// функция getToy при создании сразу определяет, к какому виду относится и создаем именно этот экземпляр

package main

import "fmt"

// общий интерфейс для всех игрушек
type iToy interface {
	setName(string)
	setColor(string)
	setStartAge(int)
	getName() string
	getColor() string
	getStartAge() int
}

// описание набора параметров для всех игрушек
type toy struct {
	name     string
	color    string
	startAge int
}

// реализация функций для любого типа игрушек
func (t *toy) setName(name string) {
	t.name = name
}

func (t *toy) setColor(color string) {
	t.color = color
}

func (t *toy) setStartAge(age int) {
	t.startAge = age
}

func (t *toy) getName() string {
	return t.name
}

func (t *toy) getColor() string {
	return t.color
}

func (t *toy) getStartAge() int {
	return t.startAge
}

// подвид игрушек: кукла
type doll struct {
	toy
}

// инициализация новой куклы со стандартными параметрами
func newDoll() iToy {
	return &doll{toy: toy{name: "Doll", color: "Multi", startAge: 3}}
}

// подвид игрушек: лег
type lego struct {
	toy
}

// инициализация новой игры лего со стандартными параметрами
func newLego() iToy {
	return &lego{toy: toy{name: "Lego", color: "Yellow", startAge: 6}}
}

// функция создания конкретного вида игрушек, в зависимости от типа toyType
func getToy(toyType string) iToy {
	if toyType == "doll" {
		return newDoll()
	}
	if toyType == "lego" {
		return newLego()
	}
	fmt.Println("Wrong toy type")
	return nil
}

func main() {
	// создаем экземпляры игрушек с указанием типа
	doll := getToy("doll")
	lego := getToy("lego")

	// выводим дефолтные параметры для каждого экземпляра
	fmt.Println("Before changes:")
	details(doll)
	fmt.Println()
	details(lego)

	// меняем параметры на другие
	doll.setName("Nano")
	lego.setColor("Red")

	// выводим значения после изменениц
	fmt.Println("\nAfter changes:")
	details(doll)
	fmt.Println()
	details(lego)
}

// функция для более понятного представления информации об экземпляре
func details(i iToy) {
	fmt.Println("Toy:", i.getName())
	fmt.Println("Color:", i.getColor())
	fmt.Println("Start age:", i.getStartAge())
}
