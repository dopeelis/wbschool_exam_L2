// поведенческий паттерн
// позволяет добавлять поведение в структуру без ее изменения

// пояснение примера: есть набор животных, реализующих интерфейс "Animal"
// и у них есть метод move, через Visitor добавляем им еще метод SAY
// при этом не изменяя интерфейс "Animal"

package main

import "fmt"

// интерфейс животные добавляем доступ для Visitor
type Animal interface {
	move()
	accept(Visitor)
}

// создаем структуры разных животных и добавляем им метод move
// и добавляет метод для доступа Visitor
type Cat struct {
}

func (c *Cat) move() {
	fmt.Println("Sneaks quietly")
}

func (c *Cat) accept(v Visitor) {
	v.visitCat(c)
}

type Dog struct {
}

func (d *Dog) move() {
	fmt.Println("Runs madly")
}

func (d *Dog) accept(v Visitor) {
	v.visitDog(d)
}

type Bird struct {
}

func (b *Bird) move() {
	fmt.Println("Flies smoothly")
}

func (b *Bird) accept(v Visitor) {
	v.visitBird(b)
}

// создаем Visitor-а
type Visitor interface {
	visitCat(c *Cat)
	visitDog(d *Dog)
	visitBird(b *Bird)
}

// и создаем "метод" say для кждого животного
type say struct {
}

func (s *say) visitCat(c *Cat) {
	fmt.Println("Meeeeow")
}

func (s *say) visitDog(d *Dog) {
	fmt.Println("Woooof")
}

func (s *say) visitBird(b *Bird) {
	fmt.Println("Tweet-tweet")
}

func main() {
	// создаем конкретные экземпляры для каждого вида животных
	abyssinian := &Cat{}
	labrador := &Dog{}
	kiwi := &Bird{}

	fmt.Print("Moving:\n\n")

	// используем обычный метод для животных
	abyssinian.move()
	labrador.move()
	kiwi.move()

	// используем через Visitor-а "метод" say
	fmt.Print("\nSaying:\n\n")
	abyssinian.accept(&say{})
	labrador.accept(&say{})
	kiwi.accept(&say{})
}
