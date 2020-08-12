package requesto

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint function is very handy while debugging
func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
