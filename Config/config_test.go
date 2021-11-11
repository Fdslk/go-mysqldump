package config

import (
	"reflect"
	"testing"
)

func TestGetConfigSuccess(t *testing.T) {
	config, _ := NewConfig("config.test.yaml")
	expexctConfig := &Config{
		DataBase: struct {
			UserName string `yaml:"username"`
			PassWord string `yaml:"password"`
			HostName string `yaml:"hostname"`
			Port     int16  `yaml:"port"`
			Dbname   string `yaml:"dbname"`
		}{
			UserName: "root",
			PassWord: "1qaz2wsx",
			HostName: "localhost",
			Port:     3306,
			Dbname:   "localEmenu",
		},
	}
	if !reflect.DeepEqual(expexctConfig, config) {
		t.Fatalf("expect %#v, got %#v ", expexctConfig, config)
	}
}

func TestGetConfigFailWithErrorPath(t *testing.T) {
	_, err := NewConfig("config.uat.yaml")
	expectedErrMessege := "open config.uat.yaml: no such file or directory"
	if !reflect.DeepEqual(expectedErrMessege, err.Error()) {
		t.Fatalf("expect %s, got %s", expectedErrMessege, err.Error())
	}
}

// func TestGetConfigFailWithMisMatchedContent(t *testing.T) {
// 	config, err := NewConfig("config.error.yaml")
// 	expectedErrMessege := "open config.uat.ymal: no such file or directory"
// 	if !reflect.DeepEqual(expectedErrMessege, err.Error()) {
// 		t.Fatalf("expect %s, got %s", expectedErrMessege, err.Error())
// 		fmt.Println(config)
// 	}
// }
