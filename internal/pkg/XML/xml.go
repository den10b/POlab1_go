package XML

import (
	"PO1/internal/pkg/domain"
	"encoding/xml"
	"fmt"
	"os"
)

func CreateXMLFile(filename string) error {
	xmlData, _ := xml.MarshalIndent(domain.XMLData{Name: "Example", Value: 42}, "", "  ")
	err := os.WriteFile(filename, []byte(xml.Header+string(xmlData)), 0644)
	if err != nil {
		fmt.Println("Ошибка при создании XML файла:", err)
		return err
	}
	fmt.Println("Создан XML файл:", filename)
	return nil
}

func WriteXMLToFile(filename string) error {
	var data domain.XMLData
	fmt.Println("Введите данные для записи в XML файл:")
	fmt.Print("Имя: ")
	fmt.Scanln(&data.Name)
	fmt.Print("Значение: ")
	fmt.Scanln(&data.Value)
	xmlData, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при маршализации данных в XML:", err)
		return err
	}
	err = os.WriteFile(filename, []byte(xml.Header+string(xmlData)), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в XML файл:", err)
		return err
	}
	fmt.Println("Данные успешно записаны в XML файл.")
	return nil
}

func ReadXMLFromFile(filename string) error {
	xmlData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении XML файла:", err)
		return err
	}
	var data domain.XMLData
	err = xml.Unmarshal(xmlData, &data)
	if err != nil {
		fmt.Println("Ошибка при десериализации данных из XML:", err)
		return err
	}
	fmt.Println("Данные из XML файла:")
	fmt.Printf("Имя: %s, Значение: %d\n", data.Name, data.Value)
	return nil
}
