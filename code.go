package code

import "code/code/gendiff"

func GenDiff(filePath1, filePath2, format string) (string, error) {
	return gendiff.GenDiff(filePath1, filePath2, format)
}
