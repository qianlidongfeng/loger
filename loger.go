package loger

type Loger interface{
	Warn(e ...interface{})
	Fatal(e ...interface{})
	Msg(label string,msg ...interface{})
	Close()
}