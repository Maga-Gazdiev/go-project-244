package builder

import (
	"code/code/gendiff/model"
	"sort"
)

func BuildDiff(data1, data2 map[string]any) []model.DiffNode {
	keys := collectKeys(data1, data2)
	sort.Strings(keys)
	var result []model.DiffNode

	for _, key := range keys {
		val1, exist1 := data1[key]
		val2, exist2 := data2[key]

		if exist1 && !exist2 {
			result = append(result, model.DiffNode{
				Key:      key,
				Status:   model.StatusRemoved,
				OldValue: val1,
				NewValue: nil,
				Children: nil,
			})
			continue
		}

		if !exist1 && exist2 {
			result = append(result, model.DiffNode{
				Key:      key,
				Status:   model.StatusAdded,
				OldValue: nil,
				NewValue: val2,
				Children: nil,
			})
			continue
		}

		if isMap(val1) && isMap(val2) {
			result = append(result, model.DiffNode{
				Key:      key,
				Status:   model.StatusNested,
				Children: BuildDiff(val1.(map[string]any), val2.(map[string]any)),
			})
			continue
		}

		if val1 == val2 {
			result = append(result, model.DiffNode{
				Key:      key,
				Status:   model.StatusUnchanged,
				OldValue: val1,
				NewValue: nil,
			})
			continue
		}

		result = append(result, model.DiffNode{
			Key:      key,
			Status:   model.StatusChanged,
			OldValue: val1,
			NewValue: val2,
		})
	}

	return result
}

func isMap(v any) bool {
	_, ok := v.(map[string]any)
	return ok
}

func collectKeys(data1, data2 map[string]any) []string {
	keys := make(map[string]struct{})

	for key := range data1 {
		keys[key] = struct{}{}
	}
	for key := range data2 {
		keys[key] = struct{}{}
	}

	result := make([]string, 0, len(keys))
	for key := range keys {
		result = append(result, key)
	}

	return result
}

