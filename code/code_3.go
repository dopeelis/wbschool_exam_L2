// Объяснить внутреннее устройство интерфейсов
// их отличие от пустых интерфейсов

package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	// T = *os.PathError V = nil
	// т.е. значение интерфейса != 0
	return err
}

func main() {
	err := Foo()
	fmt.Println(err == nil)
}

// Когда указываем тип интерфейса, то его значение отлично от нуля
// Даже если значение указателя будет nil
