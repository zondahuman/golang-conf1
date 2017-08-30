package main

import (
	"flag"
	"fmt"
	"golang-conf1/constant"
	"github.com/widuu/goini"
)

//引入数据模型
func init() {
	confName := flag.String("confName", "local", "this is conf name")

	flag.Parse()
	fmt.Println("confName:", *confName)

	constant.GlobalConfName = "golang-"+*confName+".ini"
	fmt.Println("GlobalConfName....:", constant.GlobalConfName)
	conf := goini.SetConfig("./conf/"+constant.GlobalConfName) //goini.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置
	constant.GlobalConf = *conf
	fmt.Println("GlobalConf....:", constant.GlobalConf)

}

func main() {
	username := constant.GlobalConf.GetValue("nihao", "username")
	fmt.Println("username---:", username)
}








