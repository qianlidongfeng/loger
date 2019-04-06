package loger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)
var mu sync.Mutex
func init(){
	log.SetFlags(log.Llongfile|log.Ltime|log.Ldate)
}


type LocalLoger struct{
}

func NewLocalLoger() *LocalLoger{
	return &LocalLoger{
	}
}

func (this *LocalLoger)Warn(v ...interface{}){
	var buf [1024]byte
	n:=runtime.Stack(buf[:],false)
	stack := string(buf[:n-1])
	mu.Lock()
	log.SetPrefix("[warning]")
	s:=fmt.Sprint(v...)
	log.Output(2,fmt.Sprint(s,"\n[stack]:\n",stack,"\n--------\n"))
	mu.Unlock()
}

func (this *LocalLoger)Fatal(v ...interface{}){
	var buf [1024]byte
	n:=runtime.Stack(buf[:],true)
	stack := string(buf[:n-1])
	mu.Lock()
	log.SetPrefix("[fatal]")
	s:=fmt.Sprint(v...)
	log.Output(2,fmt.Sprint(s,"\n[stack]:\n",stack,"\n--------\n"))
	os.Exit(1)
	mu.Unlock()
}

func (this *LocalLoger)Panic(v ...interface{}){
	s:=fmt.Sprint(v...)
	mu.Lock()
	log.SetPrefix("[panic]")
	log.Output(2,s)
	panic(s)
	mu.Unlock()
}

func (this *LocalLoger)Print(v ...interface{}){
	s:=fmt.Sprint(v...)
	mu.Lock()
	log.SetPrefix("[log]")
	log.Output(2,s)
	mu.Unlock()
}

func (this *LocalLoger)Debug(v ...interface{}){
	s:=fmt.Sprint(v...)
	mu.Lock()
	log.SetPrefix("[debug]")
	log.Output(2,s)
	mu.Unlock()
}

func (this *LocalLoger)Msg(label string,msg ...interface{}){
	s:=fmt.Sprint(msg...)
	mu.Lock()
	log.SetPrefix(label)
	log.Output(2,s)
	mu.Unlock()

}

func (this *LocalLoger)SetOutPut(w io.Writer){
	mu.Lock()
	log.SetOutput(w)
	mu.Unlock()
}