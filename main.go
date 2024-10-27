package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

// Пример композиции
type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login,
		acc.password,
		acc.url)
}

func (acc *account) generatePssword(countRunes int) {
	validRunes := []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-_")
	password := make([]rune, countRunes)

	for index := range password {
		password[index] += validRunes[rand.IntN(len(validRunes))]
	}
	acc.password = string(password)
}

// Конструктор
func newAccount(login,
	password,
	urlString string) (*account, error) {

	if login == "" {
		return nil, errors.New("Логин не задан")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Неверный формат URL")
	}

	acc := &account{
		url:   urlString,
		login: login,
	}

	if password == "" {
		acc.generatePssword(10)
	}

	return acc, nil
}

func newAccountWithTimeStamp(login,
	password,
	urlString string) (*accountWithTimeStamp, error) {

	acc, error := newAccount(login, password, urlString)

	if error != nil {
		fmt.Println(error)
		return nil, error
	}

	accWTS := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account:   *acc,
	}

	return accWTS, nil
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	account2, error := newAccount(login,
		password,
		url)
	if error != nil {
		fmt.Println(error)
		return
	}

	account2.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt, " ")
	var res string
	fmt.Scan(&res)
	return res
}
