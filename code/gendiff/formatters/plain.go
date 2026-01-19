package formatters

import (
	"fmt"
	"code/gendiff/model"
	"strings"
)

type PlainFormatter struct{}

func (f *PlainFormatter) Format(diffTree []model.DiffNode) string {
	var result []string
	formatPlain(diffTree, "", &result)
	return strings.Join(result, "\n")
}

func formatPlain(diffTree []model.DiffNode, pathPrefix string, result *[]string) {
	for _, node := range diffTree {
		currentPath := pathPrefix
		if currentPath != "" {
			currentPath += "." + node.Key
		} else {
			currentPath = node.Key
		}

		switch node.Status {
		case model.StatusAdded:
			value := formatPlainValue(node.NewValue)
			*result = append(*result, fmt.Sprintf("Property '%s' was added with value: %s", currentPath, value))
		case model.StatusRemoved:
			*result = append(*result, fmt.Sprintf("Property '%s' was removed", currentPath))
		case model.StatusChanged:
			oldValue := formatPlainValue(node.OldValue)
			newValue := formatPlainValue(node.NewValue)
			*result = append(*result, fmt.Sprintf("Property '%s' was updated. From %s to %s", currentPath, oldValue, newValue))
		case model.StatusNested:
			formatPlain(node.Children, currentPath, result)
		case model.StatusUnchanged:
			// Не выводим неизмененные свойства в plain формате
		}
	}
}

func formatPlainValue(value any) string {
	if value == nil {
		return "null"
	}

	switch v := value.(type) {
	case map[string]any:
		return "[complex value]"
	case []any:
		return "[complex value]"
	case string:
		return fmt.Sprintf("'%s'", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

