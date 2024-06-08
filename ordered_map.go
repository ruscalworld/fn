package fn

import "sort"

func OrderedMap[K ~string, V any](mapper func(K, V) error, m map[K]V) error {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, string(k))
	}

	sort.Strings(keys)

	for _, k := range keys {
		err := mapper(K(k), m[K(k)])
		if err != nil {
			return err
		}
	}

	return nil
}
