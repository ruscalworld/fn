package fn

func Map[I, O comparable](mapper func(I) O, values []I) []O {
	result := make([]O, len(values))

	for i, v := range values {
		result[i] = mapper(v)
	}

	return result
}

func MapErr[I, O comparable](mapper func(I) (O, error), values []I) ([]O, error) {
	result := make([]O, len(values))

	var err error
	for i, v := range values {
		result[i], err = mapper(v)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
