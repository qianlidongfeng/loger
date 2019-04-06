package netloger

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewSqlTbChecker() *SqlTbChecker{
	return &SqlTbChecker{}
}

type SqlTbChecker struct{

}

func (this *SqlTbChecker) CheckFields(db *sql.DB,tb string) (subFields[] string,err error){
	fields:= make([]string,len(TABLE_FIELDS))
	copy(fields,TABLE_FIELDS[:])
	rows, err := db.Query("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME='"+tb+"'")
	if err != nil{
		return []string{},err
	}
	var field string
	for rows.Next(){
		err:=rows.Scan(&field)
		if err != nil{
			return []string{},err
		}
		for k,v := range fields{
			if v == field{
				fields=append(fields[:k],fields[k+1:]...)
				break
			}
		}
	}
	return fields,err
}

func (this *SqlTbChecker) CheckTable(db *sql.DB,tb string)  error{
	s:=fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id BIGINT NOT NULL auto_increment,
		PRIMARY KEY (id)
	)ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci;`,tb)
	_,err:=db.Exec(s)
	return err
}
