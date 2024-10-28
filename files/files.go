package files

import (
	"fmt"
	"os"
)

func ReadFromFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

func WriteToFile(contetn []byte,
	name string) {

	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(contetn)

	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись в файл прошла успешно")
}
