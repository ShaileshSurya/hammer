package requesto

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint prettifies printing of a struct, map, array, slice using MarshalIndent function in json package
// This function is very handy while debugging using Print statements
func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
