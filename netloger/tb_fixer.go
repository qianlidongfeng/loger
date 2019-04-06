package netloger

import "database/sql"

type TbFixer interface{
	FixFields(db *sql.DB,tb string,fields []string) error
}