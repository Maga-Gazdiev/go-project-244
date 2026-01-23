package formatters

import (
	"code/code/gendiff/model"
	"encoding/json"
	"fmt"
)

type JsonFormatter struct{}

type jsonNode struct {
	Key      string      `json:"key,omitempty"`
	Type     string      `json:"type,omitempty"`
	Value1   interface{} `json:"value1,omitempty"`
	Value2   interface{} `json:"value2,omitempty"`
	Children []jsonNode  `json:"children,omitempty"`
}

type jsonRootNode struct {
	Key      string     `json:"key"`
	Type     string     `json:"type"`
	Children []jsonNode `json:"children"`
}

func (f *JsonFormatter) Format(diffTree []model.DiffNode) string {
	fmt.Println(diffTree)
	rootNode := jsonRootNode{
		Key:      "",
		Type:     "root",
		Children: convertToJsonNodes(diffTree),
	}

	jsonBytes, err := json.MarshalIndent(rootNode, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(jsonBytes)
}

func convertToJsonNodes(diffTree []model.DiffNode) []jsonNode {
	result := make([]jsonNode, 0, len(diffTree))
	for _, node := range diffTree {
		jsonNode := jsonNode{
			Key: node.Key,
		}

		switch node.Status {
		case model.StatusAdded:
			jsonNode.Type = "added"
		case model.StatusRemoved:
			jsonNode.Type = "deleted"
		case model.StatusUnchanged:
			jsonNode.Type = "unchanged"
		case model.StatusChanged:
			jsonNode.Type = "changed"
		case model.StatusNested:
			jsonNode.Type = "nested"
		}

		if node.Status == model.StatusNested {
			jsonNode.Children = convertToJsonNodes(node.Children)
		} else {
			if node.OldValue != nil {
				jsonNode.Value1 = node.OldValue
			}
			if node.NewValue != nil {
				jsonNode.Value2 = node.NewValue
			}
		}

		result = append(result, jsonNode)
	}
	return result
}
