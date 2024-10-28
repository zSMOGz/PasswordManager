package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) Output() {
	color.Green(acc.Login)
	color.Red(acc.Password)
	color.Blue(acc.Url)
}

func (acc *Account) generatePssword(countRunes int) {
	validRunes := []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-_")
	password := make([]rune, countRunes)

	for index := range password {
		password[index] += validRunes[rand.IntN(len(validRunes))]
	}
	acc.Password = string(password)
}

// Конструктор
func NewAccount(login,
	password,
	urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("логин не задан")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("неверный формат URL")
	}

	acc := &Account{
		Url:       urlString,
		Login:     login,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		acc.generatePssword(10)
	} else {
		acc.Password = password
	}

	return acc, nil
}
