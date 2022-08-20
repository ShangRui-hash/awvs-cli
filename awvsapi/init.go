package awvsapi

import (
	"crypto/tls"
	"errors"
	"net/http"
	"os"
)

var (
	AWVS_SERVER    string
	AWVS_API_KEY   string
	API_ADD_TARGET string
	API_ADD_SCAN   string

	HttpClient *http.Client
)

func Init() error {
	var (
		ok bool
	)
	//身份校验
	AWVS_SERVER, ok = os.LookupEnv("AWVS_SERVER")
	if (!ok) || AWVS_SERVER == "" {
		return errors.New("AWVS_SERVER is not set")
	}
	AWVS_API_KEY, ok = os.LookupEnv("AWVS_API_KEY")
	if (!ok) || AWVS_API_KEY == "" {
		return errors.New("AWVS_API_KEY is not set")
	}

	//路由
	API_ADD_TARGET = AWVS_SERVER + "/api/v1/targets"
	API_ADD_SCAN = AWVS_SERVER + "/api/v1/scans"

	//http客户端
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	HttpClient = &http.Client{Transport: tr}
	return nil
}
