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
			log.Printf("парсинг не удался: %v", err)
			continue
		}
		infoStr, err := dp.ActionInfo()
		if err != nil {
			log.Printf("не удалось получить информацию об активности: %v", err)
			continue
		}
		fmt.Println(infoStr)
	}
}
