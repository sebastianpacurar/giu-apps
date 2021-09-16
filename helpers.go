package main

// ConvertMapToInterfaceArray - This function takes a map[string]bool as a parameter, and spits out an interface.
///It simply splits the map parameter into {key:value} objects. Based on an index which keeps incrementing,
///the result variable will end up with content like this:  [map[geography:false] map[quiz:false]]
// TODO: need to fix this expression: result[index] = map[string]bool{k: v}
func ConvertMapToInterfaceArray(mapItems map[string]bool) []interface{} {
	result := make([]interface{}, len(mapItems))
	index := 0
	for k, v := range mapItems {
		result[index] = map[string]bool{k: v}
		index++
	}
	return result
}
