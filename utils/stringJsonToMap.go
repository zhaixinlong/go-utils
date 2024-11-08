package util

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// 字符串json 转mapInterface
func StringJsonToMap(str string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// json string to map
func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)

	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		logrus.Errorf("JsonToMap  error:%#v", err)
		return nil, err
	}
	return m, nil
}

func MapToJsonIngoreErr(m map[string]string) string {
	b, err := json.Marshal(m)
	if err != nil {
		logrus.WithError(err).Errorf("MapToJsonIngoreErr error, m: %+v", m)
		return ""
	}
	return string(b)
}

func MapToJson(m map[string]string) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		logrus.WithError(err).Errorf("MapToJson error, m: %+v", m)
		return "", err
	}
	return string(b), nil
}

func MapInterToJson(m map[string]interface{}) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		logrus.WithError(err).Errorf("MapToJson error, m: %+v", m)
		return "", err
	}
	return string(b), nil
}

func MergeMap(mObj ...map[string]string) map[string]string {
	newObj := map[string]string{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}
