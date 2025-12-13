package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	//перебираем все значения
	for _, data := range dataset {
		//очищаем пустые пробелы
		data = strings.TrimSpace(data)
		if data == "" {
			continue
		}
		//парсим данные
		err := dp.Parse(data)
		//обработка ошибки парсинга
		if err != nil {
			log.Printf("Ошибка парсинга '%s': %v", data, err)
			continue
		}

		//формируем и выводим строку с информацией
		info, err := dp.ActionInfo()
		//обработка ошибка формирования
		if err != nil {
			log.Printf("Ошибка формирования'%s': %v", data, err)
			continue
		}

		fmt.Println(info)
	}
}
