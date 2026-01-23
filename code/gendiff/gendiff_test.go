package gendiff

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func getFixturePath(filename string) string {
	_, testFile, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(testFile)
	return filepath.Join(testDir, "..", "fixture", filename)
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

func TestGenDiff_JSON_Stylish(t *testing.T) {
	file1 := getFixturePath("file1.json")
	file2 := getFixturePath("file2.json")

	result, err := GenDiff(file1, file2, "stylish")
	if err != nil {
		t.Fatalf("GenDiff вернул ошибку: %v", err)
	}

	expected, err := readExpectedResult("result_stylish.txt")
	if err != nil {
		t.Fatalf("Не удалось прочитать ожидаемый результат: %v", err)
	}

	normalizedResult := normalizeString(result)
	normalizedExpected := normalizeString(expected)

	if normalizedResult != normalizedExpected {
		t.Errorf("Результат не совпадает с ожидаемым.\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
	}
}

func TestGenDiff_JSON_Plain(t *testing.T) {
	file1 := getFixturePath("file1.json")
	file2 := getFixturePath("file2.json")

	result, err := GenDiff(file1, file2, "plain")
	if err != nil {
		t.Fatalf("GenDiff вернул ошибку: %v", err)
	}

	expected, err := readExpectedResult("result_plain.txt")
	if err != nil {
		t.Fatalf("Не удалось прочитать ожидаемый результат: %v", err)
	}

	normalizedResult := normalizeString(result)
	normalizedExpected := normalizeString(expected)

	if normalizedResult != normalizedExpected {
		t.Errorf("Результат не совпадает с ожидаемым.\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
	}
}

func TestGenDiff_JSON_JSON(t *testing.T) {
	file1 := getFixturePath("file1.json")
	file2 := getFixturePath("file2.json")

	result, err := GenDiff(file1, file2, "json")
	if err != nil {
		t.Fatalf("GenDiff вернул ошибку: %v", err)
	}

	expected, err := readExpectedResult("result_json.json")
	if err != nil {
		t.Fatalf("Не удалось прочитать ожидаемый результат: %v", err)
	}

	normalizedResult := normalizeString(result)
	normalizedExpected := normalizeString(expected)

	if normalizedResult != normalizedExpected {
		t.Errorf("Результат не совпадает с ожидаемым.\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
	}
}

func TestGenDiff_YAML_Stylish(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	result, err := GenDiff(file1, file2, "stylish")
	if err != nil {
		t.Fatalf("GenDiff вернул ошибку: %v", err)
	}

	expected, err := readExpectedResult("result_stylish.txt")
	if err != nil {
		t.Fatalf("Не удалось прочитать ожидаемый результат: %v", err)
	}

	normalizedResult := normalizeString(result)
	normalizedExpected := normalizeString(expected)

	if normalizedResult != normalizedExpected {
		t.Errorf("Результат не совпадает с ожидаемым.\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
	}
}

func TestGenDiff_YAML_Plain(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	result, err := GenDiff(file1, file2, "plain")
	if err != nil {
		t.Fatalf("GenDiff вернул ошибку: %v", err)
	}

	expected, err := readExpectedResult("result_plain.txt")
	if err != nil {
		t.Fatalf("Не удалось прочитать ожидаемый результат: %v", err)
	}

	normalizedResult := normalizeString(result)
	normalizedExpected := normalizeString(expected)

	if normalizedResult != normalizedExpected {
		t.Errorf("Результат не совпадает с ожидаемым.\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
	}
}

func TestGenDiff_YAML_JSON(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	result, err := GenDiff(file1, file2, "json")
	if err != nil {
		t.Fatalf("GenDiff вернул ошибку: %v", err)
	}

	expected, err := readExpectedResult("result_json.json")
	if err != nil {
		t.Fatalf("Не удалось прочитать ожидаемый результат: %v", err)
	}

	normalizedResult := normalizeString(result)
	normalizedExpected := normalizeString(expected)

	if normalizedResult != normalizedExpected {
		t.Errorf("Результат не совпадает с ожидаемым.\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
	}
}
