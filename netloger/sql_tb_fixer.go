package netloger

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewSqlTbFixer() *SqlTbFixer{
	return &SqlTbFixer{}
}

type SqlTbFixer struct{

}

func (this *SqlTbFixer) FixFields(db *sql.DB,tb string,fields []string) error{
	for _,v:=range fields{
		switch v {
		case "hostname","file","log":
			_,err:=db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s VARCHAR(255)",tb,v))
			if err != nil{
				return err
			}
		case "label","process":
			_,err:=db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s VARCHAR(64)",tb,v))
			if err != nil{
				return err
			}
		case "stack":
			_,err:=db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s TEXT",tb,v))
			if err != nil{
				return err
			}
		case "line","pid":
			_,err:=db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s INT",tb,v))
			if err != nil{
				return err
			}
		case "time":
			_,err:=db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s DATETIME",tb,v))
			if err != nil{
				return err
			}
		}
	}
	return nil
}