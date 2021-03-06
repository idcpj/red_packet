package main

import (
	"github.com/idcpj/red_packet/infra"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"

	_ "github.com/idcpj/red_packet"
)

func main() {

	filePath := kvs.GetCurrentFilePath("config.ini", 1)

	conf := ini.NewIniFileConfigSource(filePath)

	app := infra.NewBootApplication(conf)
	app.Start()

}
