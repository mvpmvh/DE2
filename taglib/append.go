package taglib

//Appends the key/val pairs of the first map to the second map. The override flag is optional and specifies whether
//a duplicate key in map 2 should use its value or the value from map 1. Use map 1 value if override is true. Override is false if not set.
func Append(input1 map[interface{}]interface{}, input2 map[interface{}]interface{}, override ...bool) (output map[interface{}]interface{}, err error) {
	output = input2

	for key, val := range input1 {
		if _, ok := input2[key]; ok {
			if len(override) > 0 && override[0] {
				output[key] = val
			}
		} else {
			output[key] = val
		}
	}

	return
}
