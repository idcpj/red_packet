package base

import (
	"github.com/idcpj/red_packet/infra"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisrecover "github.com/kataras/iris/middleware/recover"
	"github.com/sirupsen/logrus"
	"time"
)

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisStatus struct {
	infra.BaseStarter
}

func (i *IrisStatus) Init(ctx infra.StarterContext) {
	irisApplication = initIris()

	//日志组织配置和扩展
	logger := irisApplication.Logger()
	logger.Install(logrus.StandardLogger())

}
func (i *IrisStatus) Start(ctx infra.StarterContext) {
	//把路由器打印到控制太
	routes := Iris().GetRoutes()

	for _, v := range routes {
		logrus.Info(v.Trace())
	}

	port := ctx.Props().GetDefault("app.server.port", "1800")
	//启动 iris
	Iris().Run(iris.Addr(":" + port))

}

func (i *IrisStatus) StartBlocking() bool {
	return true
}

func initIris() *iris.Application {
	app := iris.New()
	//异常回复 irisrecover "github.com/kataras/iris/middleware/recover"
	app.Use(irisrecover.New())
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			app.Logger().Info("| %s | %s | %s | %s | %s | %s | %s | %s |",
				now.Format("2006-01-2.15:04:05"),
				latency.String(),
				status,
				ip, method,
				path,
				headerMessage)
		},
	}
	app.Use(logger.New(cfg))
	return app

}
