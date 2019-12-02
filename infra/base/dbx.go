package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/idcpj/red_packet/infra"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
)

var database *dbx.Database

func DbxDatabas() *dbx.Database {
	return database
}

type DBxDatabaseStater struct {
	infra.BaseStarter
}

func (s *DBxDatabaseStater) Setup(ctx infra.StarterContext) {
	conf := ctx.Props()
	settings := dbx.Settings{}

	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	db, err := dbx.Open(settings)
	if err != nil {
		panic("open mysql:"+err.Error())
	}
	database = db

}
