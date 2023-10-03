package files

import (
	"fmt"
	"os"
)

func CreateTextFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return err
	}
	defer file.Close()
	fmt.Println("Создан текстовый файл:", filename)
	return nil
}

func WriteTextToFile(filename string) error {
	fmt.Print("Введите текст для записи в файл: ")
	text := ""
	fmt.Scanln(&text)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Ошибка при открытии файла для записи:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text + "\n")
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return err
	}
	fmt.Println("Текст успешно записан в файл.")
	return nil
}

func ReadTextFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return err
	}
	fmt.Println("Содержимое файла:")
	fmt.Println(string(data))
	return nil
}

func DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Ошибка при удалении файла:", err)
		return err
	}
	fmt.Printf("Файл '%s' успешно удален\n", filename)
	return nil
}
