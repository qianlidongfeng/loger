package netloger

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var TABLE_FIELDS = [9]string{"time","hostname","process","pid","label","log","file","line","stack"}

func NewSqloger() *Sqloger{
	return &Sqloger{
		tbChecker:NewSqlTbChecker(),
		tbFixer:NewSqlTbFixer(),
		mu:sync.Mutex{},
	}
}

type Sqloger struct {
	db *sql.DB
	stmt *sql.Stmt
	hostname string
	process string
	tbChecker TbChcker
	tbFixer TbFixer
	mu sync.Mutex
}


func (this *Sqloger) Init(cfg SqlConfig) error{
	var err error
	this.db,err = sql.Open(cfg.Type,cfg.User+":"+cfg.PassWord+"@tcp("+cfg.Address+")/"+cfg.DB)
	if err != nil{
		return err
	}
	err=this.db.Ping()
	if err != nil{
		return err
	}
	this.db.SetMaxOpenConns(2000)
	this.db.SetMaxIdleConns(1000)
	//checktable
	err=this.tbChecker.CheckTable(this.db,cfg.Table)
	if err != nil{
		return err
	}
	//checkfields
	fields,err:=this.tbChecker.CheckFields(this.db,cfg.Table)
	if err != nil{
		return err
	}
	if len(fields) != 0{
		err = this.tbFixer.FixFields(this.db,cfg.Table,fields)
		if err != nil{
			return err
		}
	}
	fls:=""
	values:=""
	for _,v:=range TABLE_FIELDS{
		fls+=v+","
		values+="?,"
	}
	fls=strings.TrimRight(fls,",")
	values=strings.TrimRight(values,",")
	this.stmt,err = this.db.Prepare(fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`,cfg.Table,fls,values))
	if err != nil{
		return err
	}
	this.hostname,_ = os.Hostname()
	if runtime.GOOS == "windows"{
		arr:=strings.Split(os.Args[0],"\\")
		this.process=arr[len(arr)-1]
	}else{
		arr:=strings.Split(os.Args[0],"/")
		this.process=arr[len(arr)-1]
	}

	return err
}


func (this *Sqloger) Release(){
	this.stmt.Close()
	this.db.Close()
}
//0 is not limit
func (this *Sqloger) SetMaxOpenConns(n int){
	this.db.SetMaxOpenConns(n)
}
//0 is not limit
func (this *Sqloger) SetMaxIdleConns(n int){
	this.db.SetMaxOpenConns(n)
	this.db.SetMaxIdleConns(n)

}

func (this *Sqloger) Warn(e ...interface{}){
	_,file,line,_ := runtime.Caller(1)
	var buf [1024]byte
	n:=runtime.Stack(buf[:],false)
	stack := string(buf[:n-1])
	_,err:=this.stmt.Exec(time.Now().Format("2006-01-02 15:04:05"),this.hostname,this.process,os.Getpid(),"warnning",fmt.Sprint(e...),file,line,stack)
	if err != nil{
		this.mu.Lock()
		log.Println("------------------")
		log.Println("time:",time.Now())
		log.Println("process:",this.process)
		log.Println("label:warnning")
		log.Println("error:",e)
		log.Println("file:",file)
		log.Println("line:",line)
		log.Println("stack:",stack)
		_,file,line,_ = runtime.Caller(0)
		log.Println("")
		log.Println("")
		log.Println("loger error:",err)
		log.Println("file:",file)
		log.Println("line:",line)
		this.mu.Unlock()
	}
}

func (this *Sqloger) Fatal(e ...interface{}){
	_,file,line,_ := runtime.Caller(1)
	var buf [1024]byte
	n:=runtime.Stack(buf[:],true)
	stack := string(buf[:n-1])

	_,err:=this.stmt.Exec(time.Now().Format("2006-01-02 15:04:05"),this.hostname,this.process,os.Getpid(),"fatal",fmt.Sprint(e...),file,line,stack)
	if err != nil{
		this.mu.Lock()
		log.Println("------------------")
		log.Println("time:",time.Now())
		log.Println("process:",this.process)
		log.Println("label:fatal")
		log.Println("error:",e)
		log.Println("file:",file)
		log.Println("line:",line)
		log.Println("stack:",stack)
		_,file,line,_ = runtime.Caller(0)
		log.Println("")
		log.Println("")
		log.Println("loger error:",err)
		log.Println("file:",file)
		log.Println("line:",line)
		this.mu.Unlock()
	}
	this.Release()
	os.Exit(1)
}

func (this *Sqloger) Msg(label string,msg ...interface{}){
	_,err:=this.stmt.Exec(time.Now().Format("2006-01-02 15:04:05"),this.hostname,this.process,os.Getpid(),label,fmt.Sprint(msg...),nil,nil,nil)
	if err != nil{
		this.mu.Lock()
		log.Println("------------------")
		log.Println("time:",time.Now())
		log.Println("process:",this.process)
		log.Println("label:fatal")
		log.Println("msg:",msg)
		log.Println("")
		log.Println("")
		log.Println("time:",time.Now())
		log.Println("loger error:",err)
		this.mu.Unlock()
	}
}

type SqlConfig struct{
	User string
	PassWord string
	Address string
	Type string
	DB string
	Table string
}

