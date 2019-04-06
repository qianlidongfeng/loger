package loger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

var mu sync.Mutex
func init(){
	log.SetFlags(log.Llongfile|log.Ltime|log.Ldate)
}

func Warn(v ...interface{}){
	var buf [1024]byte
	n:=runtime.Stack(buf[:],false)
	stack := string(buf[:n-1])
	mu.Lock()
	log.SetPrefix("[warning]")
	s:=fmt.Sprint(v...)
	log.Output(2,fmt.Sprint(s,"\n[stack]:\n",stack))
	mu.Unlock()
}

func Fatal(v ...interface{}){
	var buf [1024]byte
	n:=runtime.Stack(buf[:],true)
	stack := string(buf[:n-1])
	mu.Lock()
	log.SetPrefix("[fatal]")
	s:=fmt.Sprint(v...)
	log.Output(2,fmt.Sprint(s,"\n[stack]:\n",stack))
	mu.Unlock()
	os.Exit(1)
}

func Panic(v ...interface{}){
	s:=fmt.Sprint(v...)
	mu.Lock()
	log.SetPrefix("[panic]")
	log.Output(2,s)
	panic(s)
	mu.Unlock()
}

func Print(v ...interface{}){
	s:=fmt.Sprint(v...)
	mu.Lock()
	log.SetPrefix("")
	log.Output(2,s)
	mu.Unlock()
}

func Debug(v ...interface{}){
	s:=fmt.Sprint(v...)
	mu.Lock()
	log.SetPrefix("[debug]")
	log.Output(2,s)
	mu.Unlock()
}