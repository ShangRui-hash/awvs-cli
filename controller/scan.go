package controller

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ShangRui-hash/awvs-cli/awvsapi"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func Scan(c *cli.Context) error {
	//从标准输入中读
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//获取输入的内容
		siteURL := scanner.Text()
		//发送到 awvs
		r, err := awvsapi.AddTarget(siteURL)
		if err != nil {
			logrus.Error("awvsapi.AddTarget failed,err:", err)
			continue
		}
		if err := awvsapi.AddScan(r.TargetId, awvsapi.FullScan); err != nil {
			logrus.Error("awvsapi.AddScan failed,err:", err)
			continue
		}
		logrus.Info(fmt.Sprintf("url:%s add scan success", siteURL))
	}
	return nil
}
