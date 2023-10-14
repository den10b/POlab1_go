package main

import (
	jason "PO1/internal/pkg/JSON"
	xml "PO1/internal/pkg/XML"
	"PO1/internal/pkg/files"
	"PO1/internal/pkg/filesystem"
	zipp "PO1/internal/pkg/zip"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var fileName string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("------------------------------------------------------")
		fmt.Println("Меню:")
		fmt.Println("1. Вывести информацию о файловой системе")
		fmt.Println("2. Создать файл")
		fmt.Println("3. Записать в файл строку")
		fmt.Println("4. Прочитать файл в консоль")
		fmt.Println("5. Удалить файл")
		fmt.Println("6. Создать файл формате JSON")
		fmt.Println("7. Создать новый объект. Выполнить сериализацию объекта в формате JSON и \nзаписать в файл.")
		fmt.Println("8. Прочитать файл в консоль")
		fmt.Println("9. Создать файл формате XML  из редактора")
		fmt.Println("10. Записать в файл новые данные из консоли .")
		fmt.Println("11. Прочитать файл в консоль.")
		fmt.Println("12. Добавить файл в архив")
		fmt.Println("13. Разархивировать файл и вывести данные о нем")
		fmt.Println("0. Выход")
		fmt.Println("------------------------------------------------------")
		fmt.Print("Введите номер команды: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "0" {
			fmt.Println("Выход из программы.")
			break
		}

		switch input {
		case "1":
			err = filesystem.PrintDiskInfo()
			if err != nil {
				fmt.Println("Ошибка при PrintDiskInfo:", err)
				continue
			}
		case "2":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = files.CreateTextFile(fileName + ".txt")
			if err != nil {
				fmt.Println("Ошибка при CreateTextFile:", err)
				continue
			}
		case "3":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = files.WriteTextToFile(fileName + ".txt")
			if err != nil {
				fmt.Println("Ошибка при WriteTextToFile:", err)
				continue
			}

		case "4":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = files.ReadTextFromFile(fileName + ".txt")
			if err != nil {
				fmt.Println("Ошибка при ReadTextFromFile:", err)
				continue
			}
		case "5":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = files.DeleteFile(fileName)
			if err != nil {
				fmt.Println("Ошибка при DeleteFile:", err)
				continue
			}
		case "6":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = jason.CreateJSONFile(fileName + ".json")
			if err != nil {
				fmt.Println("Ошибка при CreateJSONFile:", err)
				continue
			}
		case "7":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = jason.WriteJSONToFile(fileName + ".json")
			if err != nil {
				fmt.Println("Ошибка при WriteJSONToFile:", err)
				continue
			}
		case "8":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = jason.ReadJSONFromFile(fileName + ".json")
			if err != nil {
				fmt.Println("Ошибка при ReadJSONFromFile:", err)
				continue
			}
		case "9":
			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = xml.CreateXMLFile(fileName + ".xml")
			if err != nil {
				fmt.Println("Ошибка при CreateXMLFile:", err)
				continue
			}
		case "10":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = xml.WriteXMLToFile(fileName + ".xml")
			if err != nil {
				fmt.Println("Ошибка при WriteXMLToFile:", err)
				continue
			}
		case "11":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = xml.ReadXMLFromFile(fileName + ".xml")
			if err != nil {
				fmt.Println("Ошибка при ReadXMLFromFile:", err)
				continue
			}
		case "12":

			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			archName := fileName

			fmt.Println("Введите название архива: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = zipp.AddFileToZip(archName+".zip", fileName)
			if err != nil {
				fmt.Println("Ошибка при AddFileToZip:", err)
				continue
			}
		case "13":
			fmt.Println("Введите название файла: ")
			_, err = fmt.Scanf("%s\n", &fileName)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("------------------------------------------------------")
			err = zipp.UnzipFile(fileName + ".zip")
			if err != nil {
				fmt.Println("Ошибка при UnzipFile:", err)
				continue
			}
		default:
			fmt.Println("Выбран некорректный вариант")
		}
	}

}
