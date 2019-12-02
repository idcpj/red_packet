package infra

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
)

//应用程序的管理器
type BootApplication struct {
	conf           kvs.ConfigSource
	StarterContext StarterContext
}

func NewBootApplication(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{conf: conf, StarterContext: StarterContext{}}
	b.StarterContext[KeyProps]=conf
	return b
}

func (b *BootApplication) Start() {
	// 1.初始化 starter
	logrus.Info("init() ...")
	b.init()
	// 2.安装 starter
	logrus.Info("setup() ...")
	b.setup()
	// 3.启动 starter
	logrus.Info("start() ...")
	b.start()
}

func (b *BootApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.StarterContext)
	}

}
func (b *BootApplication) setup() {

	if len(StarterRegister.AllStarters())==0 {
		panic("Starters len is 0")
	}

	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.StarterContext)
	}

}
func (b *BootApplication) start() {
	for i, starter := range StarterRegister.AllStarters() {
		starter.Init(b.StarterContext)

		//是否阻塞
		if starter.StartBlocking() {
			//最后一个让他正常启动
			if (i + 1) == len(StarterRegister.AllStarters()) {
				starter.Start(b.StarterContext)
			} else {
				go starter.Start(b.StarterContext)
			}

		} else {
			starter.Start(b.StarterContext)
		}
	}

}
