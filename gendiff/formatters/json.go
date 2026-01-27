package formatters

import (
	"code/gendiff/model"
	"encoding/json"
	"strings"
)

type JsonFormatter struct{}

func (f *JsonFormatter) Format(diffTree []model.DiffNode) string {
	levelNode := model.DiffNode{
		Key:      "",
		Status:   "level",
		Children: diffTree,
	}
	data := buildJSONData(levelNode)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return formatJSON(string(jsonBytes))
}

func buildJSONData(node model.DiffNode) map[string]interface{} {
	result := make(map[string]interface{})

	if node.Status == "level" {
		result["key"] = ""
	} else if node.Key != "" {
		result["key"] = node.Key
	}

	var nodeType string
	switch node.Status {
	case model.StatusAdded:
		nodeType = model.StatusAdded
	case model.StatusRemoved:
		nodeType = model.StatusRemoved
	case model.StatusUnchanged:
		nodeType = model.StatusUnchanged
	case model.StatusChanged:
		nodeType = model.StatusChanged
	case model.StatusNested:
		nodeType = model.StatusNested
	case "level":
		nodeType = "level"
	}
	result["type"] = nodeType

	if node.Status == model.StatusNested || node.Status == "level" {
		children := make([]map[string]interface{}, 0, len(node.Children))
		for _, child := range node.Children {
			children = append(children, buildJSONData(child))
		}
		result["children"] = children
	} else {
		if node.OldValue != nil {
			result["value1"] = node.OldValue
		}
		if node.NewValue != nil {
			result["value2"] = node.NewValue
		}
	}

	return result
}

func formatJSON(jsonStr string) string {
	var result strings.Builder
	indent := 0
	inString := false
	escapeNext := false

	for i, char := range jsonStr {
		if escapeNext {
			result.WriteRune(char)
			escapeNext = false
			continue
		}

		switch char {
		case '"':
			if i > 0 && jsonStr[i-1] != '\\' {
				inString = !inString
			}
			result.WriteRune(char)
		case '\\':
			escapeNext = true
			result.WriteRune(char)
		case '{', '[':
			result.WriteRune(char)
			if !inString {
				indent++
				if i+1 < len(jsonStr) && (jsonStr[i+1] != '}' && jsonStr[i+1] != ']') {
					result.WriteString("\n" + strings.Repeat("  ", indent))
				}
			}
		case '}', ']':
			if !inString {
				indent--
				if i > 0 && jsonStr[i-1] != '{' && jsonStr[i-1] != '[' {
					result.WriteString("\n" + strings.Repeat("  ", indent))
				}
			}
			result.WriteRune(char)
		case ',':
			result.WriteRune(char)
			if !inString {
				if i+1 < len(jsonStr) && jsonStr[i+1] != ' ' {
					result.WriteString("\n" + strings.Repeat("  ", indent))
				}
			}
		case ':':
			result.WriteRune(char)
			if !inString && i+1 < len(jsonStr) && jsonStr[i+1] != ' ' {
				result.WriteRune(' ')
			}
		case ' ', '\n', '\t':
			if inString {
				result.WriteRune(char)
			}
		default:
			result.WriteRune(char)
		}
	}

	return result.String()
}
