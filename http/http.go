package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Http struct {
	logrus        *logrus.Logger
	authorization string
}

func NewHttp(logrus *logrus.Logger, authorization string) *Http {
	return &Http{
		logrus:        logrus,
		authorization: authorization,
	}
}

func (ht *Http) SendHttpGet(sendUrl string) string {
	byteRes, err := ht.SendHttpApi("GET", sendUrl, "")
	if err != nil {
		logrus.Errorf("sendHttpGet.sendHttpApi err: %+v, sendUrl: %s", err, sendUrl)
		return ""
	}
	// logrus.Debugf("res: %s", string(byteRes))
	return string(byteRes)
}

func (ht *Http) SendHttpPost(sendUrl string, params map[string]interface{}) string {
	byteParams, err := json.Marshal(params)
	if err != nil {
		logrus.Errorf("sendHttpPost.json.Marshal err: %+v, sendUrl: %s, params: %+v", err, sendUrl, params)
		return ""
	}
	byteRes, err := ht.SendHttpApi("POST", sendUrl, string(byteParams))
	if err != nil {
		logrus.Errorf("sendHttpPost.sendHttpApi err: %+v, sendUrl: %s", err, sendUrl)
		return ""
	}
	// logrus.Debugf("res: %s", string(byteRes))
	return string(byteRes)
}

func (ht *Http) SendHttpPut(sendUrl string, params map[string]interface{}) string {
	byteParams, err := json.Marshal(params)
	if err != nil {
		logrus.Errorf("sendHttpPost.json.Marshal err: %+v, sendUrl: %s, params: %+v", err, sendUrl, params)
		return ""
	}
	byteRes, err := ht.SendHttpApi("PUT", sendUrl, string(byteParams))
	if err != nil {
		logrus.Errorf("sendHttpPost.sendHttpApi err: %+v, sendUrl: %s", err, sendUrl)
		return ""
	}
	// logrus.Debugf("res: %s", string(byteRes))
	return string(byteRes)
}

// 请求待表头参数的
func (ht *Http) SendHttpApi(method string, sendUrl, params string) ([]byte, error) {
	logrus.Infof("sendHttpApi.method: %s, sendUrl: %s, params: %s", method, sendUrl, params)
	request, err := http.NewRequest(method, sendUrl, strings.NewReader(params))
	if err != nil {
		logrus.Errorf("sendHttpApi.method: %s, sendUrl: %s, params: %s", method, sendUrl, params)
		return nil, errors.Wrapf(err, "http.NewRequest error, sendUrl: %s", sendUrl)
	}

	request.Header.Set("content-type", "application/json")
	request.Header.Set("Authorization", ht.authorization)
	var client = &http.Client{
		Timeout: 1 * time.Minute,
	}

	resp, err := client.Do(request)
	if err != nil {
		logrus.Errorf("sendHttpApi.method: %s, sendUrl: %s, params: %s", method, sendUrl, params)
		return nil, errors.Wrapf(err, "client.Do error, sendUrl: %s", sendUrl)
	}
	return ioutil.ReadAll(resp.Body)
}
