package ch3

func SliceToMap(sl []string) map[string]int {
	m := map[string]int{}

	for i, v := range sl {
		m[v] = i
	}

	return m
}

func MapToTwoSlices[K any, V any](m map[any]any) ([]any, []any) {
	keys := []any{}
	value := []any{}

	for k, v := range m {
		keys = append(keys, k)
		value = append(value, v)
	}

	return keys, value

}
