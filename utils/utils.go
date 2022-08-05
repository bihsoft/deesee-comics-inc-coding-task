package utils

func FieldSet(fields ...string) map[string]bool {
	set := make(map[string]bool, len(fields))
	for _, s := range fields {
		set[s] = true
	}
	return set
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
