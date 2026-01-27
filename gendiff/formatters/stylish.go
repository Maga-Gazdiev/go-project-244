package formatters

import (
	"code/gendiff/model"
	"fmt"
	"sort"
	"strings"
)

type StylishFormatter struct{}

func (f *StylishFormatter) Format(diffTree []model.DiffNode) string {
	lines := []string{"{"}
	lines = append(lines, formatDiff(diffTree, 1)...)
	lines = append(lines, "}")
	return strings.Join(lines, "\n")
}

func formatDiff(diffTree []model.DiffNode, depth int) []string {
	var result []string

	for _, v := range diffTree {
		indent := makeIndent(depth)

		switch v.Status {
		case model.StatusAdded:
			result = append(result,
				fmt.Sprintf("%s+ %s: %s", indent, v.Key, formatValue(v.NewValue, depth+1)))
		case model.StatusRemoved:
			result = append(result,
				fmt.Sprintf("%s- %s: %s", indent, v.Key, formatValue(v.OldValue, depth+1)))
		case model.StatusUnchanged:
			result = append(result,
				fmt.Sprintf("%s  %s: %s", indent, v.Key, formatValue(v.OldValue, depth+1)))
		case model.StatusChanged:
			result = append(result,
				fmt.Sprintf("%s- %s: %s", indent, v.Key, formatValue(v.OldValue, depth+1)),
				fmt.Sprintf("%s+ %s: %s", indent, v.Key, formatValue(v.NewValue, depth+1)),
			)
		case model.StatusNested:
			result = append(result, fmt.Sprintf("%s  %s: {", indent, v.Key))
			result = append(result, formatDiff(v.Children, depth+1)...)
			result = append(result, fmt.Sprintf("%s  }", indent))
		}
	}

	return result
}

func formatValue(value any, depth int) string {
	indent := strings.Repeat(" ", depth*4)

	if value == nil {
		return "null"
	} else if m, ok := value.(map[string]any); ok {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		lines := []string{"{"}
		for _, k := range keys {
			lines = append(lines,
				fmt.Sprintf("%s%s: %s", indent, k, formatValue(m[k], depth+1)),
			)
		}
		lines = append(lines, strings.Repeat(" ", (depth-1)*4)+"}")
		return strings.Join(lines, "\n")
	} else {
		return fmt.Sprintf("%v", value)
	}
}

func makeIndent(depth int) string {
	return strings.Repeat(" ", depth*4-2)
}
