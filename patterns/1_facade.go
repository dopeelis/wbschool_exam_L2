// структурный паттерн
// идея: "скрыть" сложные реализации

// пояснение примера: создание банковских счетов и переводов через "менеджера"
// сама логика этого не реализована, только схема

package main

import "fmt"

// создаем менеджера, который объединяет клиента, его счет и транзакции
type ManagerFacade struct {
	customer    *Customer
	account     *PersonalAccount
	transaction *Transaction
}

func newManagerFacage() *ManagerFacade {
	return &ManagerFacade{&Customer{}, &PersonalAccount{}, &Transaction{}}
}

// создаем функцию для создания счета у определенного клиента
// объединяем две функции: создание клиента и создание счета
func (facade *ManagerFacade) createCustomerPersonalAccount(customerName string, persAccTitle string, id int) (*Customer, *PersonalAccount) {
	customer := facade.customer.create(customerName, id)
	personalAccount := facade.account.create(persAccTitle, id)
	return customer, personalAccount
}

// создаем функцию для проведения транзакции через "менеджера"
func (facade *ManagerFacade) createTransaction(amount float32, id int) *Transaction {
	transaction := facade.transaction.create(amount, id)
	return transaction
}

// объявляем структуру клиента
type Customer struct {
	name string
	id   int
}

// функция для создания клиента
func (customer *Customer) create(name string, id int) *Customer {
	fmt.Println("Creating customer")
	customer.name = name
	customer.id = id
	return customer
}

// объявляем структуру персонального счета
type PersonalAccount struct {
	title      string
	customerId int
}

// функция для создания персонального счета
func (persAcc *PersonalAccount) create(title string, customerId int) *PersonalAccount {
	fmt.Println("Creating personal account")
	persAcc.title = title
	persAcc.customerId = customerId
	return persAcc
}

// объявляем структуру транзакций
type Transaction struct {
	amount     float32
	customerId int
}

// функция для создания транзакции
func (transaction *Transaction) create(amount float32, customerId int) *Transaction {
	fmt.Println("Creating transaction")
	transaction.amount = amount
	transaction.customerId = customerId
	return transaction
}

func main() {
	facade := newManagerFacage()
	customer, account := facade.createCustomerPersonalAccount("Liza S", "Savings account", 1)
	fmt.Printf("Customer %v name:%s\n", customer.id, customer.name)
	fmt.Printf("Personal Account of customer %v: %s\n", account.customerId, account.title)
	transaction := facade.createTransaction(666, customer.id)
	fmt.Printf("Transaction for customer %v for %f successful\n", transaction.customerId, transaction.amount)
	fmt.Printf("Account balance for customer %v: %f\n ", transaction.customerId, transaction.amount)
}
