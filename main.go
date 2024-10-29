package main

import (
	"fmt"
	"strings"

	"main.go/account"
	"main.go/encryptor"
	"main.go/files"
	"main.go/output"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("Менеджер паролей")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти env файл")
	}

	vault := account.NewVault(files.NewJsonDb("data.vault"),
		*encryptor.NewEncryptor())

	isContinue := true
	for isContinue {
		isContinue = getMenu(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	newAccount, error := account.NewAccount(login,
		password,
		url)
	if error != nil {
		fmt.Println(error)
		return
	}

	vault.AddAccount(*newAccount)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, url string) bool {
		return strings.Contains(acc.Url, url)
	})

	outputResaultFind(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	url := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, login string) bool {
		return strings.Contains(acc.Login, login)
	})

	outputResaultFind(&accounts)
}

func outputResaultFind(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("По указанному параметру аккаунтов не нашлось")
	} else {
		for _, account := range *accounts {
			account.Output()
		}
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountsByURL(url)
	if isDeleted {
		color.Green("Найденные аккаунты были удалены")
	} else {
		output.PrintError("По указанному URL аккаунтов не нашлось")
	}
}

func promptData(prompt ...any) string {
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
	command := promptData("1. Создать аккаунт",
		"2. Найти аккаунт по URL",
		"3. Найти аккаунт по логину",
		"4. Удалить аккаунт",
		"5. Выход",
		"Выберете вариант",
	)

	menuFunc := menu[command]
	if menuFunc == nil {
		return false
	} else {
		menuFunc(vault)
		return true
	}
}
