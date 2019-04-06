package netloger_test

import (
	"github.com/qianlidongfeng/loger/netloger"
	"testing"
)

func TestSqloger_Fatal(t *testing.T) {
	loger := netloger.NewSqloger()
	cfg := netloger.SqlConfig{
		User:     "root",
		PassWord: "333221",
		Address:  "127.0.0.1:3306",
		DB:       "project_logs",
		Table:    "test",
		Type:     "mysql",
	}
	err := loger.Init(cfg)
	if err != nil {
		t.Error(err)
	}
	loger.Fatal("xixihaha")
	if err != nil {
		t.Error(err)
	}
}

func TestSqloger_Warn(t *testing.T) {
	loger := netloger.NewSqloger()
	cfg := netloger.SqlConfig{
		User:     "root",
		PassWord: "333221",
		Address:  "127.0.0.1:3306",
		DB:       "project_logs",
		Table:    "test",
		Type:     "mysql",
	}
	err := loger.Init(cfg)
	if err != nil {
		t.Error(err)
	}
	loger.Warn("xixihaha")
	loger.Release()
}

func TestSqloger_Msg(t *testing.T) {
	loger := netloger.NewSqloger()
	cfg := netloger.SqlConfig{
		User:     "root",
		PassWord: "333221",
		Address:  "127.0.0.1:3306",
		DB:       "project_logs",
		Table:    "test",
		Type:     "mysql",
	}
	err := loger.Init(cfg)
	if err != nil {
		t.Error(err)
	}
	loger.Msg("xixihaha","info")
	loger.Release()
}