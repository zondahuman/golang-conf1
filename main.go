package main

import (
	"flag"
	"fmt"
	"golang-conf1/constant"
	"github.com/widuu/goini"
	"runtime"
	"strings"
)

//引入数据模型
func init() {
	confName := flag.String("confName", "local", "this is conf name")

	flag.Parse()
	fmt.Println("confName:", *confName)
	fmt.Println("current operatesystem....:", runtime.GOOS)
	constant.GlobalConfName = "golang-"+*confName+".ini"
	fmt.Println("GlobalConfName....:", constant.GlobalConfName)
	var confPath string = ""
	operateSystem := runtime.GOOS
	if strings.EqualFold(operateSystem, constant.LINUX) {
		confPath = "../conf/" + constant.GlobalConfName
	}else {
		confPath = "./conf/" + constant.GlobalConfName
	}
	conf := goini.SetConfig(confPath) //goini.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置
	constant.GlobalConf = *conf
	fmt.Println("GlobalConf....:", constant.GlobalConf.ReadList())
}

func main() {
	username := constant.GlobalConf.GetValue("nihao", "username")
	fmt.Println("username---:", username)
}








