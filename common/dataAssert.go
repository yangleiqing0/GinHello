package common

func IsStringSliceHas(target interface{}, slice []string) (bool, int) {
	for index, key := range slice {
		if key == target {
			return true, index
		}
	}
	return false, -1
}
