package gendiff

import (
	"code/code/gendiff/builder"
	"code/code/gendiff/formatters"
	"code/code/gendiff/parser"
)

func GenDiff(filePath1, filePath2, format string) (string, error) {
	data1, err := parser.ParseFile(filePath1)
	if err != nil {
		return "", err
	}

	data2, err := parser.ParseFile(filePath2)
	if err != nil {
		return "", err
	}

	diffTree := builder.BuildDiff(data1, data2)

	formatter, err := formatters.GetFormatter(format)
	if err != nil {
		return "", err
	}

	return formatter.Format(diffTree), nil
}
