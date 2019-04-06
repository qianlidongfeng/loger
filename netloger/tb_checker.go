package netloger

import "database/sql"

type TbChcker interface{
	CheckFields(*sql.DB,string) ([]string, error)
	CheckTable(*sql.DB,string) (error)
}