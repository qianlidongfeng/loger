package loger

import (
	"github.com/qianlidongfeng/loger/netloger"
	"github.com/qianlidongfeng/toolbox"
)

type Loger interface{
	Warn(e ...interface{})
	Fatal(e ...interface{})
	Msg(label string,msg ...interface{})
	Close()
}

type Config struct{
	LogType string
	DB toolbox.MySqlConfig
}

func NewLoger(cfg Config) (loger Loger,err error){
	if cfg.LogType == "netlog"{
		lg:=netloger.NewSqloger()
		err=lg.Init(cfg.DB)
		if err != nil{
			return
		}
		loger = lg
	}else{
		loger=NewLocalLoger()
	}
	return
}