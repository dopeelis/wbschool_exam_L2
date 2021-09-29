// поведенческий паттерн
// позволяет менять поведение
// в зависимости от внутреннего состояния

// пояснение примера: есть светофор, который, в зависимости от своего состояниях сейчас,
// меняет цвет на идущий именно за ним
// красный - красный с желтым - зеленый - желтый - снова красный и т.д

package main

import "fmt"

type state interface {
	handle(*trafficLights)
}

// создаем стуктуру светофор, у которой есть состояние
type trafficLights struct {
	state state
}

// запрашиваем цвет и меняем его на следующий
func (tL *trafficLights) request() {
	tL.state.handle(tL)
}

// функция для смены состояния
func (tL *trafficLights) changeState(state state) {
	tL.state = state
}

// определяем виды состояний
type redLight struct{}
type redAndYellowLight struct{}
type greenLight struct{}
type yellowLight struct{}

// реализуем функции смены для каждого состояния
func (r *redLight) handle(t *trafficLights) {
	fmt.Println("The red light is on. Next: red and yellow.")
	t.changeState(new(redAndYellowLight))
}

func (rY *redAndYellowLight) handle(t *trafficLights) {
	fmt.Println("The red and yellow light is on. Next: green")
	t.changeState(new(greenLight))
}

func (g *greenLight) handle(t *trafficLights) {
	fmt.Println("The green light is on. Next: yellow")
	t.changeState(new(yellowLight))
}

func (y *yellowLight) handle(t *trafficLights) {
	fmt.Println("The yellow light is on. Next: red")
	t.changeState(new(redLight))
}

func main() {
	// создаем новый светофор. первое состояние - красный
	trafficLights1 := trafficLights{new(redLight)}
	// далее делаем запрос и меняем состояние на следующее
	trafficLights1.request()
	trafficLights1.request()
	trafficLights1.request()
	trafficLights1.request()
	trafficLights1.request()
}
