package common

import (
	"encoding/json"
	"fmt"
)

func GetArgs(args []string) string {

	argBytes, err := json.Marshal(&ArgStruct{
		Func: args[0],
		Args: args[1:],
	})
	if err != nil {

		panic(fmt.Errorf("error marshaling empty args struct: %v", err))
	}
	return string(argBytes)
}

func getDefaultValueAndDescription(defaultValue string, defaultDescription string, overrides ...string) (value, description string) {
	if len(overrides) > 0 && len(overrides[0]) > 0 {
		value = overrides[0]
	} else {
		value = defaultValue
	}
	if len(overrides) > 1 && len(overrides[1]) > 0 {
		description = overrides[1]
	} else {
		description = defaultDescription
	}
	return value, description
}
