package main

import (
	"fmt"

	"github.com/fatih/color"
	"main.go/account"
)

func main() {
	fmt.Println("Менеджер паролей")
	vault := account.NewVault()

	isContinue := true
	for isContinue {
		isContinue = getMenu(vault)
	}
}

func createAccount() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	newAccount, error := account.NewAccount(login,
		password,
		url)
	if error != nil {
		fmt.Println(error)
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*newAccount)
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите URL для поиска: ")
	accounts := vault.FindAccountsByURL(url)

	if len(accounts) == 0 {
		color.Red("По указанному URL аккаунтов не нашлось")
	} else {
		for _, account := range accounts {
			account.Output()
		}
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите URL для поиска: ")
	isDeleted := vault.DeleteAccountsByURL(url)
	if isDeleted {
		color.Green("Найденные аккаунты были удалены")
	} else {
		color.Red("По указанному URL аккаунтов не нашлось")
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt, " ")
	var res string
	fmt.Scan(&res)
	return res
}

func printMenu() {
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
}

func getMenu(vault *account.Vault) bool {
	printMenu()

	var command int
	fmt.Scan(&command)

	switch command {
	case 1:
		createAccount()
		return true
	case 2:
		findAccount(vault)
		return true
	case 3:
		deleteAccount(vault)
		return true
	default:
		return false
	}
}
