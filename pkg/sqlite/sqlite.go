package sqlite

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	_defaultTimeout = time.Second
)

type SQLite struct {
	Timeout time.Duration
	DB      *sql.DB
}

func New(url string, timeout string) (*SQLite, error) {
	db, err := openDB(url)
	if err != nil {
		return nil, err
	}
	to, err := time.ParseDuration(timeout)
	if err != nil {
		to = _defaultTimeout
	}

	err = prepareDB(db)
	if err != nil {
		return nil, err
	}

	mysql := &SQLite{
		Timeout: to,
		DB:      db,
	}

	return mysql, nil
}

func (p *SQLite) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func prepareDB(db *sql.DB) error {
	query := `
		DROP TABLE IF EXISTS snippets;
		CREATE TABLE IF NOT EXISTS snippets(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title varchar(100) NOT NULL,
			content text NOT NULL,
			created datetime NOT NULL,
			expires datetime NOT NULL
		);
		INSERT INTO snippets (id, title, content, created, expires) VALUES
		(1, 'Не имей сто рублей', 'Не имей сто рублей,\nа имей сто друзей.', '2021-01-27 13:09:34', '2030-01-27 13:09:34'),
		(2, 'Лучше один раз увидеть', 'Лучше один раз увидеть,\nчем сто раз услышать.', '2021-01-27 13:09:40', '2030-01-27 13:09:40'),
		(3, 'Не откладывай на завтра', 'Не откладывай на завтра,\nчто можешь сделать сегодня.', '2021-01-27 13:09:44', '2021-02-03 13:09:44'),
		(4, 'История про улитку', 'Улитка выползла из раковины,\nвытянула рожки,\nи опять подобрала их.', '2021-04-21 14:39:12', '2021-04-28 14:39:12');
		`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
