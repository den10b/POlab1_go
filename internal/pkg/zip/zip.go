package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func AddFileToZip(zipFilename string, filenameToAdd string) error {
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		fmt.Println("Ошибка при создании ZIP файла:", err)
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileToAdd, err := os.Open(filenameToAdd)
	if err != nil {
		fmt.Println("Ошибка при открытии файла для добавления в ZIP архив:", err)
		return err
	}
	defer fileToAdd.Close()

	info, err := fileToAdd.Stat()
	if err != nil {
		fmt.Println("Ошибка при получении информации о файле:", err)
		return err
	}

	header := &zip.FileHeader{
		Name:   filenameToAdd,
		Method: zip.Deflate,
	}
	header.SetModTime(info.ModTime())

	zipFileWriter, err := zipWriter.CreateHeader(header)
	if err != nil {
		fmt.Println("Ошибка при создании файла в ZIP архиве:", err)
		return err
	}

	_, err = io.Copy(zipFileWriter, fileToAdd)
	if err != nil {
		fmt.Println("Ошибка при копировании файла в ZIP архив:", err)
		return err
	}

	fmt.Printf("Файл '%s' добавлен в ZIP архив '%s'\n", filenameToAdd, zipFilename)

	return nil
}

func UnzipFile(zipFilename string) error {
	reader, err := zip.OpenReader(zipFilename)
	if err != nil {
		fmt.Println("Ошибка при открытии ZIP архива:", err)
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		fmt.Printf("Разархивирован файл: %s\n", file.Name)

		// Чтение содержимого файла
		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("Ошибка при открытии файла в ZIP архиве:", err)
			return err
		}
		defer fileReader.Close()

		data, err := io.ReadAll(fileReader)
		if err != nil {
			fmt.Println("Ошибка при чтении файла из ZIP архива:", err)
			return err
		}

		fmt.Printf("Содержимое файла:\n%s\n", data)
	}
	return nil
}
