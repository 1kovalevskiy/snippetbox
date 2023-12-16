package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	_defaultTimeout = time.Second
)

type MySQL struct {
	Timeout time.Duration
	DB      *sql.DB
}

func New(url string, timeout string) (*MySQL, error) {
	db, err := openDB(url)
	if err != nil {
		return nil, err
	}
	to, err := time.ParseDuration(timeout)
	if err != nil {
		to = _defaultTimeout
	}
	mysql := &MySQL{
		Timeout: to,
		DB:      db,
	}

	return mysql, nil
}

func (p *MySQL) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
