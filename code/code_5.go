package main

// Создаем свою структуру ошибок
type customError struct {
	msg string
}

// Добавляем ему функцию Error()
// для реализации интерфейса error
func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error

	err = test()
	// T = *customError, V=nil
	// т.е. значение интерфейса != 0
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
