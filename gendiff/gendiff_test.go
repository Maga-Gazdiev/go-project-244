package gendiff

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func getFixturePath(filename string) string {
	_, testFile, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(testFile)
	projectRoot := filepath.Clean(filepath.Join(testDir, "..", ".."))
	if filepath.Base(projectRoot) != "go-project-lvl2" {
		projectRoot = filepath.Clean(filepath.Join(projectRoot, "go-project-lvl2"))
	}
	return filepath.Join(projectRoot, "testdata", "fixture", filename)
}

func readExpectedResult(filename string) (string, error) {
	path := getFixturePath(filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func normalizeString(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")

	lines := strings.Split(s, "\n")
	normalized := make([]string, 0, len(lines))
	for _, line := range lines {
		normalized = append(normalized, strings.TrimRight(line, " \t\r"))
	}
	for len(normalized) > 0 && normalized[len(normalized)-1] == "" {
		normalized = normalized[:len(normalized)-1]
	}
	return strings.Join(normalized, "\n")
}

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name           string
		file1          string
		file2          string
		format         string
		expectedResult string
	}{
		{
			name:           "JSON with stylish format",
			file1:          "file1.json",
			file2:          "file2.json",
			format:         "stylish",
			expectedResult: "result_stylish.txt",
		},
		{
			name:           "JSON with plain format",
			file1:          "file1.json",
			file2:          "file2.json",
			format:         "plain",
			expectedResult: "result_plain.txt",
		},
		{
			name:           "JSON with json format",
			file1:          "file1.json",
			file2:          "file2.json",
			format:         "json",
			expectedResult: "result_json.json",
		},
		{
			name:           "YAML with stylish format",
			file1:          "file1.yml",
			file2:          "file2.yml",
			format:         "stylish",
			expectedResult: "result_stylish.txt",
		},
		{
			name:           "YAML with plain format",
			file1:          "file1.yml",
			file2:          "file2.yml",
			format:         "plain",
			expectedResult: "result_plain.txt",
		},
		{
			name:           "YAML with json format",
			file1:          "file1.yml",
			file2:          "file2.yml",
			format:         "json",
			expectedResult: "result_json.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file1 := getFixturePath(tt.file1)
			file2 := getFixturePath(tt.file2)

			result, err := GenDiff(file1, file2, tt.format)
			if err != nil {
				t.Fatalf("GenDiff вернул ошибку: %v", err)
			}

			expected, err := readExpectedResult(tt.expectedResult)
			if err != nil {
				t.Fatalf("Не удалось прочитать ожидаемый результат: %v", err)
			}

			if tt.format == "json" {
				var resultJSON, expectedJSON interface{}
				if err := json.Unmarshal([]byte(result), &resultJSON); err != nil {
					t.Fatalf("Не удалось распарсить результат как JSON: %v", err)
				}
				if err := json.Unmarshal([]byte(expected), &expectedJSON); err != nil {
					t.Fatalf("Не удалось распарсить ожидаемый результат как JSON: %v", err)
				}

				resultBytes, _ := json.Marshal(resultJSON)
				expectedBytes, _ := json.Marshal(expectedJSON)

				if string(resultBytes) != string(expectedBytes) {
					t.Errorf("Результат не совпадает с ожидаемым.\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
				}
			} else {
				normalizedResult := normalizeString(result)
				normalizedExpected := normalizeString(expected)

				if normalizedResult != normalizedExpected {
					t.Errorf("Результат не совпадает с ожидаемым.\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
				}
			}
		})
	}
}
