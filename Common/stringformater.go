package common

import (
	"encoding/json"
	"log"

	"github.com/TylerBrock/colorjson"
)

func GetColorfulString(obj interface{}) (string, error) {
	var newMap map[string]interface{}
	data, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	json.Unmarshal(data, &newMap)

	f := colorjson.NewFormatter()
	f.Indent = 2
	result, err := f.Marshal(newMap)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(result), nil
}
