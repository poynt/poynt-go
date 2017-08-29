package poynt

import (
	"encoding/json"
	"fmt"
)

// pretty print any poynt data model
func PrettyPrint(model interface{}) {
	bytes, _ := json.MarshalIndent(model, "", "    ")
	fmt.Println(string(bytes))
}

func Stringify(structInstance interface{}) (string, error) {
	jsonByteArr, err := json.Marshal(structInstance)
	if err != nil {
		fmt.Println("err: ", err)
	}

	return string(jsonByteArr), err
}
