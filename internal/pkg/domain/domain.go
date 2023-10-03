package domain

// Структура для данных в JSON
type Data struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Структура для данных в XML
type XMLData struct {
	Name  string `xml:"Name"`
	Value int    `xml:"Value"`
}
