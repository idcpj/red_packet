package base

import (
	"github.com/idcpj/red_packet/infra"
	"github.com/tietang/props/kvs"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource{
	return props
}


type PropsStatus struct {
	infra.BaseStarter
}


func (p *PropsStatus) Init(ctx infra.StarterContext){
	//c:= ini.NewIniFileConfigSource("config.ini")



}
