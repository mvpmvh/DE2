//package for testing all templace functions

package taglib_test

import (
	"fmt"
	"source.discoveryeducation.com/users/mhatch/repos/de2/taglib"
)

func Example_Append_Should_Append_Keys_From_Map_To_Map() {
	var map1 = map[interface{}]interface{}{"a": "A", "b": "B"}
	var map2 = map[interface{}]interface{}{"c": "C", "d": "D"}
	var override = false

	var outputMap, _ = taglib.Append(map1, map2, override)
	fmt.Println(outputMap)
	//Output: map[c:C d:D a:A b:B]

}

func Example_Append_Should_Override_Map_2_Keys_When_Override_True() {
	var map1 = map[interface{}]interface{}{"a": "A", "b": "B"}
	var map2 = map[interface{}]interface{}{"a": "a", "c": "C"}
	var override = true

	var outputMap, _ = taglib.Append(map1, map2, override)
	fmt.Println(outputMap)
	//Output: map[a:A c:C b:B]
}

func Example_Append_Should_Append_Mixed_Types() {
	var map1 = map[interface{}]interface{}{"a": "A", "b": "B"}
	var map2 = map[interface{}]interface{}{1: 1, 2: 2}
	var override = true

	var outputMap, _ = taglib.Append(map1, map2, override)
	fmt.Println(outputMap)
	//Output: map[1:1 2:2 a:A b:B]
}
