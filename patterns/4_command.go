// поведенческий паттерн
// превращает запросы в объекты,
// позволяя передавать их как аргументы при вызове методов

// пояснение примера: имеется счет и несколько операций с ним
// поплнить, списать, проверить баланс
// также имеются команды с методом ВЫПОЛНИТЬ (execute) для каждой операции

package main

import "fmt"

// создаем интерфейс счета с методами
type account interface {
	deposit(float32)  // поплнить
	withdraw(float32) // списать
	checkbalance()    // проверить баланс
}

type command interface {
	execute(float32)
}
type depositCommand struct {
	account account
}

func (c *depositCommand) execute(sum float32) {
	c.account.deposit(sum)
}

type withdrawCommand struct {
	account account
}

func (c *withdrawCommand) execute(sum float32) {
	c.account.withdraw(sum)
}

type checkbalanceCommand struct {
	account account
}

func (c *checkbalanceCommand) execute(float32) {
	c.account.checkbalance()
}

type personalAccount struct {
	ownerName string
	balance   float32
}

func (pAcc *personalAccount) deposit(sum float32) {
	pAcc.balance += sum
	fmt.Println("Balance replenishment operation completed successfully")
}

func (pAcc *personalAccount) withdraw(sum float32) {
	if pAcc.balance >= sum {
		pAcc.balance -= sum
		fmt.Println("Debit operation completed successfully")
	}
	fmt.Println("There are not enough funds on the account")
}

func (pAcc *personalAccount) checkbalance() {
	fmt.Printf("Account %s balance: %f\n", pAcc.ownerName, pAcc.balance)
}

type operation struct {
	command command
}

// запускаем выполнение операций (метод запуска команд)
func (o *operation) do(sum float32) {
	o.command.execute(sum)
}

func main() {
	myAcc := &personalAccount{ownerName: "Lisa S", balance: 0}

	depositCommand := &depositCommand{account: myAcc}
	checkbalanceCommand := &checkbalanceCommand{account: myAcc}

	depositOperation := &operation{command: depositCommand}
	checkbalanceOperation := &operation{command: checkbalanceCommand}

	depositOperation.do(10)
	checkbalanceOperation.do(0)

	withdrawCommand := &withdrawCommand{account: myAcc}
	withdrawOperation := &operation{command: withdrawCommand}

	withdrawOperation.do(50)
}
