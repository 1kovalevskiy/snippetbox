package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/1kovalevskiy/snippetbox/internal/entity"
	"github.com/1kovalevskiy/snippetbox/pkg/mysql"
)

type SnippetRepo struct {
	*mysql.MySQL
}

func New(mysql_ *mysql.MySQL) *SnippetRepo {
	return &SnippetRepo{mysql_}
}

func (r *SnippetRepo) Insert(ctx context.Context, t entity.SnippetCreate) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(ctx, r.Timeout)
	defer cancel()

	result, err := stmt.ExecContext(ctx, t.Title, t.Content, t.Expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *SnippetRepo) Get(ctx context.Context, id int) (*entity.Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, r.Timeout)
	defer cancel()

	row := stmt.QueryRowContext(ctx, id)

	var val entity.Snippet
	err = row.Scan(&val.ID, &val.Title, &val.Content, &val.Created, &val.Expires)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entity.ErrNoRecord
	}
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func (r *SnippetRepo) Latest(ctx context.Context) ([]*entity.Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, r.Timeout)
	defer cancel()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*entity.Snippet

	for rows.Next() {
		s := &entity.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
