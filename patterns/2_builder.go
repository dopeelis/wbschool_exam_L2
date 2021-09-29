// порождающий паттерн
// даёт возможность использовать один и тот же код строительства
// для получения разных представлений объектов

// пояснение примера: строитель выдает полный почтовый адрес,
// в зависимости от страны, куда отправляет письмо

package main

import (
	"fmt"
)

// создаем строителя
type MailAddressBuilder interface {
	setTitle(string) string
	setFirstLine(string, string, string) string
	setSecondLine(string, string, string) string
	getMailAddress()
}

// определяем его направления, в зависимости от поступивших данных
func getMailAddresBuilder(county string) MailAddressBuilder {
	if county == "Russia" {
		return &RusMailAddressBuilder{}
	}
	if county == "USA" {
		return &AmericanMailAddressBuilder{}
	}
	return nil
}

// создаем структуру письма в Россию
type RusMailAddressBuilder struct {
	title     string
	street    string
	houseNum  string
	apartment string
	сity      string
	region    string
	postcode  string
}

// создаем функцию для определения заголовка
func (rusAddress *RusMailAddressBuilder) setTitle(title string) string {
	rusAddress.title = title
	s := fmt.Sprint(rusAddress.title)
	return s
}

// создаем функцию для определения первой почтовой строки
func (rusAddress *RusMailAddressBuilder) setFirstLine(street string, houseNum string, apartment string) string {
	rusAddress.street = street
	rusAddress.houseNum = houseNum
	rusAddress.apartment = apartment
	s := fmt.Sprint("ул. ", rusAddress.street, ", д. ", rusAddress.houseNum, ", кВ. ", rusAddress.apartment)
	return s
}

// создаем функцию для определения второй почтовой строки
func (rusAddress *RusMailAddressBuilder) setSecondLine(сity string, region string, postcode string) string {
	rusAddress.сity = сity
	rusAddress.region = region
	rusAddress.postcode = postcode
	s := fmt.Sprint("г.", rusAddress.сity, ",", rusAddress.region, ",", rusAddress.postcode)
	return s
}

// создаем функцию для получения полного почтового адреса
func (rusAddress *RusMailAddressBuilder) getMailAddress() {
	fmt.Println(rusAddress.setTitle(rusAddress.title))
	fmt.Println(rusAddress.setFirstLine(rusAddress.street, rusAddress.houseNum, rusAddress.apartment))
	fmt.Println(rusAddress.setSecondLine(rusAddress.сity, rusAddress.region, rusAddress.postcode))
}

// создаем структуру письма в Америку
type AmericanMailAddressBuilder struct {
	title     string
	houseNum  string
	street    string
	apartment string
	сity      string
	state     string
	zipcode   string
}

// создаем функцию для определения заголовка
func (americanAddress *AmericanMailAddressBuilder) setTitle(title string) string {
	americanAddress.title = title
	s := fmt.Sprint(americanAddress.title)
	return s
}

// создаем функцию для определения первой почтовой строки
func (americanAddress *AmericanMailAddressBuilder) setFirstLine(houseNum string, street string, apartment string) string {
	americanAddress.houseNum = houseNum
	americanAddress.street = street
	americanAddress.apartment = apartment
	s := fmt.Sprint(americanAddress.houseNum, " ", americanAddress.street, " ", americanAddress.apartment)
	return s
}

// создаем функцию для определения второй почтовой строки
func (americanAddress *AmericanMailAddressBuilder) setSecondLine(сity string, state string, zipcode string) string {
	americanAddress.сity = сity
	americanAddress.state = state
	americanAddress.zipcode = zipcode
	s := fmt.Sprint(americanAddress.сity, ",", americanAddress.state, ",", americanAddress.zipcode)
	return s
}

// создаем функцию для получения полного почтового адреса
func (americanAddress *AmericanMailAddressBuilder) getMailAddress() {
	fmt.Println(americanAddress.setTitle(americanAddress.title))
	fmt.Println(americanAddress.setFirstLine(americanAddress.houseNum, americanAddress.street, americanAddress.apartment))
	fmt.Println(americanAddress.setSecondLine(americanAddress.сity, americanAddress.state, americanAddress.zipcode))
}

func main() {
	// объявляем для строителя тип адреса
	americanAddress := getMailAddresBuilder("USA")
	russianAddress := getMailAddresBuilder("Russia")

	// задаем данные для Американского адреса
	americanAddress.setTitle("Josh P.")
	americanAddress.setFirstLine("775", "Westminster Avenue", "APT D5")
	americanAddress.setSecondLine("Brooklyn", " NY", " 11230")
	americanAddress.getMailAddress()

	fmt.Println()

	// задаем данные для Российского адреса
	russianAddress.setTitle("Аля П.")
	russianAddress.setFirstLine("Луговая", "55", "38")
	russianAddress.setSecondLine("Бор", " обл. Нижегородская", " 603278")
	russianAddress.getMailAddress()

}
