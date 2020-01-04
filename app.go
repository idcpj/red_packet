package resk

import (
	"github.com/idcpj/red_packet/infra"
	"github.com/idcpj/red_packet/infra/base"
)

func init() {
	infra.Register(&base.PropsStatus{})
	infra.Register(&base.DBxDatabaseStater{})
	infra.Register(&base.LogStatus{})
	infra.Register(&base.ValidatorStater{})

	//最后一个阻塞的程序 web
	infra.Register(&base.IrisStatus{})
}
