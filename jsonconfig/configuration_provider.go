package jsonconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type t interface {
}

func GetSection(section string) (string, error) {
	jsonMap, err := getJsonMap()
	if err != nil {
		return "", err
	}

	arr := strings.Split(section, ":")

	var t t
	var m map[string]interface{}
	flag := false

	for _, v := range arr {
		if flag == false {
			t = jsonMap[v]
			flag = true
		} else {
			t = m[v]
		}
		typeStr := fmt.Sprintf("%T", t)
		if typeStr == "map[string]interface {}" {
			m = t.(map[string]interface{})
		} else if t != nil {
			return t.(string), nil
		} else {
			return "", nil
		}
	}
	return "", nil
}

func Bind(i interface{}, section string) error {
	jsonMap, err := getJsonMap()
	if err != nil {
		return err
	}

	var t t
	var m map[string]interface{}
	flag := false

	arr := strings.Split(section, ":")
	for _, v := range arr {
		if flag == false {
			t = jsonMap[v]
			flag = true

		} else {
			t = m[v]
		}

		typeStr := fmt.Sprintf("%T", t)
		if typeStr == "map[string]interface {}" {
			m = t.(map[string]interface{})
		}
	}

	st, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = json.Unmarshal(st, i)
	if err != nil {
		return err
	}
	return nil
}

func getJsonMap() (map[string]interface{}, error) {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}
