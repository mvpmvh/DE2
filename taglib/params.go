/* taglib allows for the generation of custom tag library functions to be used in templates */

package taglib

import "errors"

func Params(input ...interface{}) (map[string]interface{}, error) {
	if len(input)%2 != 0 {
		return nil, errors.New("input must be in pairs")
	}
	dict := make(map[string]interface{}, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		key, ok := input[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = input[i+1]
	}
	return dict, nil
}
