package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func ParseFile(path string) (map[string]any, error) {
	absolPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(absolPath)
	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		return parseJSON(data)
	case ".yml", ".yaml":
		return parseYAML(data)
	default:
		return nil, fmt.Errorf("неподдерживаемый формат: %s", ext)
	}
}

func parseJSON(data []byte) (map[string]any, error) {
	var result map[string]any
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parseYAML(data []byte) (map[string]any, error) {
	var result map[string]any
	err := yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
