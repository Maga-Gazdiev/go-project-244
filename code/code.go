package code

import "hexlet-project-lvl2/gendiff"

func GenDiff(filePath1, filePath2, format string) (string, error) {
	return gendiff.GenDiff(filePath1, filePath2, format)
}
