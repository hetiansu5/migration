package mysql

import "database/sql"

type Driver struct {
	db *sql.DB
}

func NewDriver(db *sql.DB) *Driver {
	return &Driver{
		db: db,
	}
}

func (d *Driver) Run(statement string) error {
	_, err := d.db.Exec(statement)
	return err
}
