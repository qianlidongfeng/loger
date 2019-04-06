package netloger

type loger interface{
	Warn(e ...interface{})
	Fatal(e ...interface{})
	Msg(label string,msg ...interface{})
}
