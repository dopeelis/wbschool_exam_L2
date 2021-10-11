// структурный паттерн
// идея: "скрыть" сложные реализации

// пояснение примера: создание банковских счетов через "менеджера"
// он вызывает две службы - службу безопасности и службу регистрации
// при создании счета запускает работу обеих служб, сначала проверяя на наличие такого счета
// и такого клиента, если нет, то вызывает вторую службу для создания

package main

import (
	"fmt"
)

// база всех клиентов
var customerBase = make(map[int]Customer)

// Find функция для поиска элемента в слайсе
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		} else {
			continue
		}
	}
	return false
}

// ManagerFacade - "менеджер", который объединяет две службы
type ManagerFacade struct {
	SecurityService     *SecurityService
	RegistrationService *RegistrationService
}

func newManagerFacage() *ManagerFacade {
	return &ManagerFacade{&SecurityService{}, &RegistrationService{}}
}

// createCustomer создаем нового клиента через менеджера
func (facade *ManagerFacade) createCustomer(name string, id int) *Customer {
	// служба безопасности проверяет наличие клиента
	if facade.SecurityService.checkCustomer(id) {
		fmt.Println("This customer already exist")
		return nil
	}
	// если ок, то служба регистрации создает клиента
	customer := facade.RegistrationService.createCustomer(name, id)
	return customer
}

// createPersonalAccount создаем нового счета клиента через менеджера
func (facade *ManagerFacade) createPersonalAccount(id int, title string) *PersonalAccount {
	// служба безопасности проверяет наличие клиента
	if !facade.SecurityService.checkCustomer(id) {
		fmt.Printf("Customer with id %v doesn't exist\n", id)
		return nil
	}
	// служба безопасности проверяет аккаунт
	if facade.SecurityService.checkPersonalAccount(id, title) {
		fmt.Printf("Customer %v already have account %s\n", id, title)
		return nil
	}
	customer := customerBase[id]
	// если ок, служба регистрации создает аккаунт
	personalAccount := facade.RegistrationService.createPersonalAccount(&customer, title)
	return personalAccount
}

// getInfo предоставление информации о клиенте
func (facade *ManagerFacade) getInfo(customer *Customer) {
	if facade.SecurityService.checkCustomer(customer.id) {
		fmt.Printf("Name: %s\nID: %v\nAccounts: %v\n", customer.name, customer.id, customer.personalAccounts)
		return
	}
	fmt.Println("Customer doesn't exist")
}

// SecurityService структура службы безопасности
type SecurityService struct {
	customer *Customer
	account  *PersonalAccount
}

// checkCustomer проверяет наличие клиента в базе
func (secServ *SecurityService) checkCustomer(id int) bool {
	_, ok := customerBase[id]
	return ok
}

// checkPersonalAccount проверяет наличие аккаунта у данного клиента
func (secServ *SecurityService) checkPersonalAccount(id int, title string) bool {
	customer := customerBase[id]
	return Find(customer.personalAccounts, title)
}

// RegistrationService структура службы регистрации
type RegistrationService struct {
	customer *Customer
	account  *PersonalAccount
}

// createCustomer функция создания клиента
func (regServ *RegistrationService) createCustomer(name string, id int) *Customer {
	fmt.Println("Creating customer")
	customer := newCustomer()
	customer.name = name
	customer.id = id
	customerBase[customer.id] = *customer
	return customer
}

// createPersonalAccount функция создания персонального счета у клиента
func (regServ *RegistrationService) createPersonalAccount(customer *Customer, title string) *PersonalAccount {
	fmt.Println("Creating personal account")
	personalAccount := newPersonalAccount()
	personalAccount.title = title
	customer.personalAccounts = append(customer.personalAccounts, title)
	return personalAccount
}

// Customer струкрура клиента
type Customer struct {
	name             string
	id               int
	personalAccounts []string
}

func newCustomer() *Customer {
	customer := &Customer{}
	return customer
}

// PersonalAccount структура счета
type PersonalAccount struct {
	title string
}

func newPersonalAccount() *PersonalAccount {
	personalAccount := &PersonalAccount{}
	return personalAccount
}

func main() {
	// создаем нового менеджера
	facade := newManagerFacage()
	// создаем первого клиента через менеджера
	Bob := facade.createCustomer("Bob", 1)

	// получаем о нем информацию
	facade.getInfo(Bob)

	// создаем счет через менеджера
	facade.createPersonalAccount(1, "Saving account")
}
