package resk

import (
	"github.com/idcpj/red_packet/infra"
	"github.com/idcpj/red_packet/infra/base"
)

func init() {
	infra.Register(&base.PropsStatus{})
	infra.Register(&base.DBxDatabaseStater{})
	infra.Register(&base.LogStatus{})
}