package main

import (
	jason "PO1/internal/pkg/JSON"
	xml "PO1/internal/pkg/XML"
	"PO1/internal/pkg/files"
	"PO1/internal/pkg/filesystem"
	zipp "PO1/internal/pkg/zip"
	"fmt"
	"os"
)

func main() {
	err := filesystem.PrintDiskInfo()
	if err != nil {
		fmt.Println("Ошибка при PrintDiskInfo:", err)
		os.Exit(1)
	}

	err = files.CreateTextFile("myTestFile.txt")
	if err != nil {
		fmt.Println("Ошибка при CreateTextFile:", err)
		os.Exit(1)
	}
	err = files.ReadTextFromFile("myTestFile.txt")
	if err != nil {
		fmt.Println("Ошибка при ReadTextFromFile:", err)
		os.Exit(1)
	}
	err = files.WriteTextToFile("myTestFile.txt")
	if err != nil {
		fmt.Println("Ошибка при WriteTextToFile:", err)
		os.Exit(1)
	}
	err = files.ReadTextFromFile("myTestFile.txt")
	if err != nil {
		fmt.Println("Ошибка при ReadTextFromFile:", err)
		os.Exit(1)
	}
	err = files.DeleteFile("myTestFile.txt")
	if err != nil {
		fmt.Println("Ошибка при DeleteFile:", err)
		os.Exit(1)
	}

	err = jason.CreateJSONFile("myTestFile.json")
	if err != nil {
		fmt.Println("Ошибка при CreateJSONFile:", err)
		os.Exit(1)
	}
	err = jason.ReadJSONFromFile("myTestFile.json")
	if err != nil {
		fmt.Println("Ошибка при ReadJSONFromFile:", err)
		os.Exit(1)
	}
	err = jason.WriteJSONToFile("myTestFile.json")
	if err != nil {
		fmt.Println("Ошибка при WriteJSONToFile:", err)
		os.Exit(1)
	}
	err = jason.ReadJSONFromFile("myTestFile.json")
	if err != nil {
		fmt.Println("Ошибка при ReadJSONFromFile:", err)
		os.Exit(1)
	}

	err = xml.CreateXMLFile("myTestFile.xml")
	if err != nil {
		fmt.Println("Ошибка при CreateXMLFile:", err)
		os.Exit(1)
	}
	err = xml.ReadXMLFromFile("myTestFile.xml")
	if err != nil {
		fmt.Println("Ошибка при ReadXMLFromFile:", err)
		os.Exit(1)
	}
	err = xml.WriteXMLToFile("myTestFile.xml")
	if err != nil {
		fmt.Println("Ошибка при WriteXMLToFile:", err)
		os.Exit(1)
	}
	err = xml.ReadXMLFromFile("myTestFile.xml")
	if err != nil {
		fmt.Println("Ошибка при ReadXMLFromFile:", err)
		os.Exit(1)
	}

	err = zipp.AddFileToZip("myArchive.zip", "to_zip.txt")
	if err != nil {
		fmt.Println("Ошибка при AddFileToZip:", err)
		os.Exit(1)
	}
	err = zipp.UnzipFile("myArchive.zip")
	if err != nil {
		fmt.Println("Ошибка при UnzipFile:", err)
		os.Exit(1)
	}

}
