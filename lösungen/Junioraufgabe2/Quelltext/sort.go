package main

type Term struct {
	a string
	b string
}

func Sort(terms []Term) map[string]int {
	var valueMap = make(map[string]int)
	var oldMap = CopyMap(valueMap)
	for _, term := range terms {
		if valueMap[term.a] == 0 {
			valueMap[term.a] = 1
		}
		if valueMap[term.b] == 0 {
			valueMap[term.b] = 1
		}
		if valueMap[term.a] <= valueMap[term.b] {
			valueMap[term.a] = valueMap[term.b] + 1
		}
		if CompareMap(valueMap, oldMap) {
			break
		}
	}
	return valueMap
}

func Greatest(sortedMap map[string]int) string {
	var greatest string
	for k, v := range sortedMap {
		if greatest == "" {
			greatest = k
			continue
		}
		if sortedMap[greatest] < v {
			greatest = k
		}
	}
	for k, v := range sortedMap {
		if v == sortedMap[greatest] && k != greatest {
			return "-1"
		}
	}
	return greatest
}
