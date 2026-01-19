package gendiff

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenDiffStylish(t *testing.T) {
	tests := []struct {
		name     string
		file1    string
		file2    string
		expected string
	}{
		{
			name:  "test1",
			file1: filepath.Join("..", "testdata", "file1.json"),
			file2: filepath.Join("..", "testdata", "file2.json"),
			expected: `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`,
		},
		{
			name:  "test2",
			file1: filepath.Join("..", "testdata", "file1.yml"),
			file2: filepath.Join("..", "testdata", "file2.yml"),
			expected: `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			diff, err := GenDiff(test.file1, test.file2, "stylish")
			if err != nil {
				t.Fatalf("ошибка: %v", err)
			}
			if diff != test.expected {
				t.Errorf("получено:\n%s\nожидается:\n%s", diff, test.expected)
			}
		})
	}
}

func TestGenDiffPlain(t *testing.T) {
	tests := []struct {
		name     string
		file1    string
		file2    string
		expected string
	}{
		{
			name:  "plain test1",
			file1: filepath.Join("..", "testdata", "file3.json"),
			file2: filepath.Join("..", "testdata", "file4.json"),
			expected: `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			diff, err := GenDiff(test.file1, test.file2, "plain")
			if err != nil {
				t.Fatalf("ошибка: %v", err)
			}
			if diff != test.expected {
				t.Errorf("получено:\n%s\nожидается:\n%s", diff, test.expected)
			}
		})
	}
}

func TestGenDiffJson(t *testing.T) {
	tests := []struct {
		name     string
		file1    string
		file2    string
		validate func(t *testing.T, output string)
	}{
		{
			name:  "json test1",
			file1: filepath.Join("..", "testdata", "file1.json"),
			file2: filepath.Join("..", "testdata", "file2.json"),
			validate: func(t *testing.T, output string) {
				var result []map[string]interface{}
				err := json.Unmarshal([]byte(output), &result)
				if err != nil {
					t.Fatalf("JSON невалиден: %v", err)
				}

				// Проверяем наличие ожидаемых ключей
				expectedKeys := map[string]bool{
					"follow":  false,
					"host":    false,
					"proxy":   false,
					"timeout": false,
					"verbose": false,
				}

				for _, node := range result {
					key, ok := node["key"].(string)
					if !ok {
						t.Errorf("узел не содержит ключ 'key' или он не строка")
						continue
					}
					if _, exists := expectedKeys[key]; exists {
						expectedKeys[key] = true
					}

					// Проверяем наличие обязательных полей
					if _, ok := node["status"]; !ok {
						t.Errorf("узел '%s' не содержит поле 'status'", key)
					}
				}

				// Проверяем, что все ожидаемые ключи найдены
				for key, found := range expectedKeys {
					if !found {
						t.Errorf("ожидаемый ключ '%s' не найден в результате", key)
					}
				}

				// Проверяем конкретные статусы
				foundStatuses := make(map[string]bool)
				for _, node := range result {
					status, ok := node["status"].(string)
					if ok {
						foundStatuses[status] = true
					}
				}

				expectedStatuses := []string{"remove", "unchanged", "changed", "added"}
				for _, status := range expectedStatuses {
					if !foundStatuses[status] {
						t.Errorf("ожидаемый статус '%s' не найден", status)
					}
				}
			},
		},
		{
			name:  "json test2 - nested structures",
			file1: filepath.Join("..", "testdata", "file3.json"),
			file2: filepath.Join("..", "testdata", "file4.json"),
			validate: func(t *testing.T, output string) {
				var result []map[string]interface{}
				err := json.Unmarshal([]byte(output), &result)
				if err != nil {
					t.Fatalf("JSON невалиден: %v", err)
				}

				// Проверяем наличие вложенных структур
				hasNested := false
				for _, node := range result {
					if children, ok := node["children"]; ok && children != nil {
						hasNested = true
						break
					}
				}

				if !hasNested {
					t.Error("ожидались вложенные структуры, но они не найдены")
				}

				// Проверяем, что JSON содержит ожидаемые ключи верхнего уровня
				topLevelKeys := []string{"common", "group1", "group2", "group3"}
				foundKeys := make(map[string]bool)
				for _, node := range result {
					if key, ok := node["key"].(string); ok {
						foundKeys[key] = true
					}
				}

				for _, key := range topLevelKeys {
					if !foundKeys[key] && (key == "common" || key == "group1") {
						t.Errorf("ожидаемый ключ верхнего уровня '%s' не найден", key)
					}
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			diff, err := GenDiff(test.file1, test.file2, "json")
			if err != nil {
				t.Fatalf("ошибка: %v", err)
			}

			// Проверяем, что вывод не пустой
			if strings.TrimSpace(diff) == "" {
				t.Error("вывод пустой")
			}

			// Выполняем валидацию
			test.validate(t, diff)
		})
	}
}
