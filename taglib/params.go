/* taglib allows for the generation of custom tag library functions to be used in templates */

package taglib

import "errors"

//template function to add arbitrary key/val data to a template's context
func Params(input ...interface{}) (map[interface{}]interface{}, error) {
	if len(input)%2 != 0 {
		return nil, errors.New("Input parameters must be in pairs!")
	}
	dict := make(map[interface{}]interface{}, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		key := input[i]
		dict[key] = input[i+1]
	}
	return dict, nil
}
