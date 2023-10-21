package u

func Die[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
