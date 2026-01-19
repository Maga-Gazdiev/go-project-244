package formatters

import (
	"fmt"
	"hexlet-project-lvl2/gendiff/model"
)

type Formatter interface {
	Format(diffTree []model.DiffNode) string
}

func GetFormatter(formatName string) (Formatter, error) {
	switch formatName {
	case "stylish":
		return &StylishFormatter{}, nil
	case "plain":
		return &PlainFormatter{}, nil
	case "json":
		return &JsonFormatter{}, nil
	default:
		return nil, fmt.Errorf("неподдерживаемый формат: %s", formatName)
	}
}

