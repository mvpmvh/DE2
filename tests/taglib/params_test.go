package taglib_test

import (
	"fmt"
	"source.discoveryeducation.com/users/mhatch/repos/de2/taglib"
	"testing"
)

func Example_Params_Should_Return_An_Error_When_Odd_Number_of_Args() {
	_, err := taglib.Params("key1", "val1", "ke2")
	fmt.Println(err)
	//Output: Input parameters must be in pairs!
}

func TestParams_Should_Return_Nil_Map_When_Odd_Number_of_Args(t *testing.T) {
	if m, _ := taglib.Params("key1", "val1", "key2"); m != nil {
		t.Error("Nil should be returned for map when odd arguments provided")
	}
}
