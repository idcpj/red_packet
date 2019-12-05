package base

import (
	"github.com/idcpj/red_packet/infra"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/go-utils"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)


var formatter *prefixed.TextFormatter




type LogStatus struct {
	infra.BaseStarter
}

func (p *LogStatus) Init(ctx infra.StarterContext) {
	pinit()
}

func (p *LogStatus) Setup(ctx infra.StarterContext){


	conf := ctx.Props()

	InitLog(conf)

	log.Info(" log init success")
}


func pinit() {


	//formatter :=&log.TextFormatter{}// logrus 自带的文本格式
	formatter :=&prefixed.TextFormatter{} //引用第三方的文本格式
	formatter.FullTimestamp=true
	formatter.TimestampFormat="2006-01-02.15:04:05.000000"
	formatter.ForceFormatting=true

	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})

	log.SetFormatter(formatter)
	//日志级别
	level:=os.Getenv("log.debug")
	source := ini.NewIniFileConfigSource("config.ini")
	if level =="true" {
		log.SetLevel(log.DebugLevel)
	}else{
		b, e := source.GetBool("log.debug")
		if e != nil {
			panic(e)
		}
		if b {
			log.SetLevel(log.DebugLevel)

		}
	}
	log.SetReportCaller(true)


	//日志文件和滚动配置
	log.Info("日志系统启动...")

}



//初始化log配置，配置logrus日志文件滚动生成和
func InitLog(conf kvs.ConfigSource) {
	//设置日志输出级别
	level, err := log.ParseLevel(conf.GetDefault("log.level", "info"))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
	showfileLength()

	//配置日志输出目录
	logDir := conf.GetDefault("log.dir", "../logs")
	logTestDir, err := conf.Get("log.test.dir")
	if err == nil {
		logDir = logTestDir
	}
	logPath := logDir //+ "/logs"
	logFilePath, _ := filepath.Abs(logPath)
	log.Infof("log_path: %s", logFilePath)
	logFileName := conf.GetDefault("log.file.name", "log")
	maxAge := conf.GetDurationDefault("log.max.age", time.Hour*24)
	rotationTime := conf.GetDurationDefault("log.rotation.time", time.Hour*1)
	os.MkdirAll(logPath, os.ModePerm)

	baseLogPath := path.Join(logPath, logFileName)
	//设置滚动日志输出writer
	writer, err := rotatelogs.New(
		strings.TrimSuffix(baseLogPath, ".log")+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", err)
	}
	//设置日志文件输出的日志格式
	//formatter := &log.JSONFormatter{}

	formatter =&prefixed.TextFormatter{} //引用第三方的文本格式
	//控制台高亮显示
	formatter.DisableColors=true
	formatter.FullTimestamp=true
	formatter.TimestampFormat="2006-01-02.15:04:05.000000"
	formatter.ForceFormatting=true

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, formatter)

	log.AddHook(lfHook)

}

func showfileLength() {
	lfh := utils.NewLineNumLogrusHook()
	lfh.EnableFileNameLog = true
	lfh.EnableFuncNameLog = false
	log.AddHook(lfh)
}
