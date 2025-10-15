package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка парсинга: %v", err)
			continue
		}
		infoStr, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка получения информации об активности: %v", err)
			continue
		}
		fmt.Println(infoStr)
	}
}
