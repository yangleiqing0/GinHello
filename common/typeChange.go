package common

func SliceToMap(sList []string) map[string]string {
	mapList := make(map[string]string)
	for i := 0; i < len(sList); i++ {
		mapList[sList[i]] = ""
	}
	return mapList
}
