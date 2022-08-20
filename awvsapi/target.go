package awvsapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type AddTargetParam struct {
	Address     string `json:"address"`     //目标url
	Description string `json:"description"` //备注
	Criticality int    `json:"criticality"` //危险程度；范围:[30,20,10,0]; 默认为 10
}

type AddTargetResponse struct {
	TargetId string `json:"target_id"` //目标id
}

func AddTarget(url string) (*AddTargetResponse, error) {
	//构造请求
	request, err := http.NewRequest("POST", API_ADD_TARGET, nil)
	if err != nil {
		logrus.Error("http.NewRequest failed,err:", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf8")
	request.Header.Set("X-Auth", AWVS_API_KEY)
	param := AddTargetParam{
		Address:     url,
		Description: "awvs-cli",
		Criticality: 10,
	}
	requestBody, err := json.Marshal(param)
	if err != nil {
		logrus.Error("json.Marshal failed,err:", err)
		return nil, err
	}
	request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	//发送请求
	resp, err := HttpClient.Do(request)
	if err != nil {
		logrus.Error("http.DefaultClient.Do failed,err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	//解析响应
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("ioutil.ReadAll failed,err:", err)
		return nil, err
	}
	var response AddTargetResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		logrus.Error("json.Unmarshal failed,err:", err)
		return nil, err
	}
	return &response, nil
}
