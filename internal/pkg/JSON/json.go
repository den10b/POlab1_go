package JSON

import (
	"PO1/internal/pkg/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func CreateJSONFile(filename string) error {
	data := domain.Data{Name: "Example", Value: 42}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка при маршализации данных в JSON:", err)
		return err
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при создании JSON файла:", err)
		return err
	}
	fmt.Println("Создан JSON файл:", filename)

	return nil
}

func WriteJSONToFile(filename string) error {
	var data domain.Data
	fmt.Println("Введите данные для записи в JSON файл:")
	fmt.Print("Имя: ")
	fmt.Scanln(&data.Name)
	fmt.Print("Значение: ")
	fmt.Scanln(&data.Value)

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка при маршализации данных в JSON:", err)
		return err
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в JSON файл:", err)
		return err
	}
	fmt.Println("Данные успешно записаны в JSON файл.")

	return nil
}

func ReadJSONFromFile(filename string) error {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении JSON файла:", err)
		return err
	}
	var data domain.Data
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Ошибка при десериализации данных из JSON:", err)
		return err
	}
	fmt.Println("Данные из JSON файла:")
	fmt.Printf("Имя: %s, Значение: %d\n", data.Name, data.Value)

	return nil
}
