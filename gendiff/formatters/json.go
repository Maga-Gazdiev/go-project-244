package formatters

import (
	"encoding/json"
	"code/gendiff/model"
)

type JsonFormatter struct{}

type jsonNode struct {
	Key      string      `json:"key,omitempty"`
	Status   string      `json:"status,omitempty"`
	OldValue interface{} `json:"oldValue,omitempty"`
	NewValue interface{} `json:"newValue,omitempty"`
	Children []jsonNode  `json:"children,omitempty"`
}

func (f *JsonFormatter) Format(diffTree []model.DiffNode) string {
	jsonNodes := convertToJsonNodes(diffTree)
	jsonBytes, err := json.MarshalIndent(jsonNodes, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(jsonBytes)
}

func convertToJsonNodes(diffTree []model.DiffNode) []jsonNode {
	result := make([]jsonNode, 0, len(diffTree))
	for _, node := range diffTree {
		jsonNode := jsonNode{
			Key:    node.Key,
			Status: node.Status,
		}

		if node.Status == model.StatusNested {
			jsonNode.Children = convertToJsonNodes(node.Children)
		} else {
			if node.OldValue != nil {
				jsonNode.OldValue = node.OldValue
			}
			if node.NewValue != nil {
				jsonNode.NewValue = node.NewValue
			}
		}

		result = append(result, jsonNode)
	}
	return result
}


