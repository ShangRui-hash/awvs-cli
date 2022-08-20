package awvsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	//扫描模板
	FullScan                          = "11111111-1111-1111-1111-111111111111" //完全扫描
	HighRiskVulnerabilities           = "11111111-1111-1111-1111-111111111112" //高风险漏洞
	CrossSiteScriptingVulnerabilities = "11111111-1111-1111-1111-111111111116" //XSS 漏洞
	SQLInjectionVulnerabilities       = "11111111-1111-1111-1111-111111111113" //SQL 注入漏洞
	WeakPasswords                     = "11111111-1111-1111-1111-111111111115" //弱口令检测
	CrawlOnly                         = "11111111-1111-1111-1111-111111111117" //Crawl Only
	MalwareScan                       = "11111111-1111-1111-1111-111111111120" //恶意软件扫描
)

type Schedule struct {
	Disable bool `json:"disable"` //是否禁用
	// StartDate     string `json:"start_date"`     //开始时间
	TimeSensitive bool `json:"time_sensitive"` //是否是时间筛选
}

// data = {"target_id": target_id, "profile_id": profile_id, "incremental": False,
//                 "schedule": {"disable": False, "start_date": None, "time_sensitive": False}}
type AddScanParam struct {
	TargetId    string   `json:"target_id"`   //目标id
	ProfileId   string   `json:"profile_id"`  //扫描模板id
	Incremental bool     `json:"incremental"` //是否增量扫描
	Schedule    Schedule `json:"schedule"`
}

func AddScan(targetId string, scanType string) error {
	//构造请求
	request, err := http.NewRequest("POST", API_ADD_SCAN, nil)
	if err != nil {
		logrus.Error("http.NewRequest failed,err:", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf8")
	request.Header.Set("X-Auth", AWVS_API_KEY)
	param := AddScanParam{
		TargetId:    targetId,
		ProfileId:   scanType,
		Incremental: false,
		Schedule: Schedule{
			Disable:       false,
			TimeSensitive: false,
		},
	}
	requestBody, err := json.Marshal(param)
	if err != nil {
		logrus.Error("json.Marshal failed,err:", err)
		return err
	}
	request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	//发送请求
	resp, err := HttpClient.Do(request)
	if err != nil {
		logrus.Error("http.DefaultClient.Do failed,err:", err)
		return err
	}
	defer resp.Body.Close()
	//解析响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("ioutil.ReadAll failed,err:", err)
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("http.StatusOK != resp.StatusCreated,err:%s", string(body))
	}
	return nil
}
