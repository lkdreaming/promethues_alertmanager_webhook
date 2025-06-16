package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	http2 "promethues_alertmanager_webhook/common/constant/http"
	"reflect"
	"strings"
)

func DoHttpRequestPostJson(url string, method string, headers map[string]interface{}, data map[string]interface{},
	responseType interface{}, dataType interface{}) (interface{}, error) {
	method = strings.ToUpper(method)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating request: %s", err))
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, errors.New(fmt.Sprintf("Error creating request: %s", err))
	}
	for k, v := range headers {
		value, ok := v.(string)
		if ok {
			request.Header.Add(k, value)
		}
	}
	request.Header.Add(http2.CONTENT_TYPE, "application/json")
	log.Println(request.Header)
	log.Println(url)
	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error making request: %s", err))
	}
	log.Println("resp: ", resp)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading response: %s", err))
	}
	log.Println("body: ", string(body))
	if responseType == nil {
		return (string)(body), nil
	}
	// 创建 responseType 类型的实例
	responseInstance := reflect.New(reflect.TypeOf(responseType)).Interface()
	if err := json.Unmarshal(body, responseInstance); err == nil {
		// 使用反射获取 responseInstance 中的 Data 字段
		responseValue := reflect.ValueOf(responseInstance).Elem()
		dataField := responseValue.FieldByName("Data")
		if dataField.IsValid() && !dataField.IsNil() {
			dataBytes, err := json.Marshal(dataField.Interface())
			if err == nil {
				// 创建 dataType 类型的实例
				dataInstance := reflect.New(reflect.TypeOf(dataType)).Interface()
				if err := json.Unmarshal(dataBytes, dataInstance); err == nil {
					// 设置 Data 字段的值
					dataField.Set(reflect.ValueOf(dataInstance).Elem())
				} else {
					fmt.Println("dataType err: ", err)
				}
			}
		}
		return responseInstance, nil
	}

	// 如果响应无法解析为 responseType，则返回原始 JSON 字符串
	return string(body), nil
}
