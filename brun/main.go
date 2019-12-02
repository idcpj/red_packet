package main

import (
	"fmt"
	"github.com/idcpj/red_packet/infra"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"

	_ "github.com/idcpj/red_packet"
)

func main() {

	fmt.Println("ini  ...")

	filePath := kvs.GetCurrentFilePath("config.ini", 1)

	conf := ini.NewIniFileConfigSource(filePath)

	app := infra.NewBootApplication(conf)
	app.Start()

	//做循环用
	ints := make(chan int)
	<-ints
}


