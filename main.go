package main

import (
	"os"

	"github.com/ShangRui-hash/awvs-cli/awvsapi"
	"github.com/ShangRui-hash/awvs-cli/controller"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

var configFile string

func main() {
	author := cli.Author{
		Name:  "无在无不在",
		Email: "2227627947@qq.com",
	}
	app := &cli.App{
		Name:      "awvs-cli",
		Usage:     "awvs-cli",
		UsageText: "awvs-cli",
		Version:   "v0.1",
		Authors:   []*cli.Author{&author},
		Flags:     []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:   "scan",
				Usage:  "send url to scanner",
				Action: controller.Scan,
			},
			{
				Name:  "vuln",
				Usage: "get vuln list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}
	//初始化awvs-api
	if err := awvsapi.Init(); err != nil {
		logrus.Error("awvsapi.Init failed,err:", err)
		return
	}
	//启动app
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
	}
}
