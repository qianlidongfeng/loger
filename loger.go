package loger

import (
	"fmt"
	"log"
	"os"
)

func init(){
	log.SetFlags(log.Llongfile|log.Ltime|log.Ldate)
}

func Warn(v ...interface{}){
	log.SetPrefix("[warning]")
	log.Output(2,fmt.Sprint(v...))
}

func Fatal(v ...interface{}){
	log.SetPrefix("[fatal]")
	log.Output(2,fmt.Sprint(v...))
	os.Exit(1)
}

func Panic(v ...interface{}){
	s:=fmt.Sprint(v...)
	log.SetPrefix("[panic]")
	log.Output(2,s)
	panic(s)
}

func Print(v ...interface{}){
	s:=fmt.Sprint(v...)
	log.SetPrefix("")
	log.Output(2,s)
}

func Debug(v ...interface{}){
	s:=fmt.Sprint(v...)
	log.SetPrefix("")
	log.Output(2,s)
}