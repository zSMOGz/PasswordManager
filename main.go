package main

import (
	"fmt"

	"github.com/fatih/color"
	"main.go/account"
	"main.go/files"
	"main.go/output"
)

func main() {
	fmt.Println("Менеджер паролей")
	vault := account.NewVault(files.NewJsonDb("data.json"))

	isContinue := true
	for isContinue {
		isContinue = getMenu(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

	newAccount, error := account.NewAccount(login,
		password,
		url)
	if error != nil {
		fmt.Println(error)
		return
	}

	vault.AddAccount(*newAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	accounts := vault.FindAccountsByURL(url)

	if len(accounts) == 0 {
		output.PrintError("По указанному URL аккаунтов не нашлось")
	} else {
		for _, account := range accounts {
			account.Output()
		}
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	isDeleted := vault.DeleteAccountsByURL(url)
	if isDeleted {
		color.Green("Найденные аккаунты были удалены")
	} else {
		output.PrintError("По указанному URL аккаунтов не нашлось")
	}
}

func promptData[T any](prompt []T) string {
	for index, element := range prompt {
		if index == (len(prompt) - 1) {
			fmt.Printf("%v: ", element)
		} else {
			fmt.Println(element)
		}
	}
	var res string
	fmt.Scan(&res)
	return res
}

func getMenu(vault *account.VaultWithDb) bool {
	command := promptData([]string{
		"1. Создать аккаунт",
		"2. Найти аккаунт",
		"3. Удалить аккаунт",
		"4. Выход",
		"Выберете вариант",
	})

	switch command {
	case "1":
		createAccount(vault)
		return true
	case "2":
		findAccount(vault)
		return true
	case "3":
		deleteAccount(vault)
		return true
	default:
		return false
	}
}
