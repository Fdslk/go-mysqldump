package common

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetColorFulString(t *testing.T) {
	expectStr := "{\n  \"str\": \"foo\"\n}"
	var obj map[string]interface{}
	json.Unmarshal([]byte(expectStr), &obj)
	result, _ := GetColorfulString(obj)
	if !reflect.DeepEqual(expectStr, result) {
		t.Fatalf("expect %s, got %s", expectStr, result)
	}
}
