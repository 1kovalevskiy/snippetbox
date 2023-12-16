package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/1kovalevskiy/snippetbox/internal/entity"
	"github.com/1kovalevskiy/snippetbox/pkg/sqlite"
)

type SnippetRepo struct {
	*sqlite.SQLite
}

func New(mysql_ *sqlite.SQLite) *SnippetRepo {
	return &SnippetRepo{mysql_}
}

func (r *SnippetRepo) Insert(ctx context.Context, t entity.SnippetCreate) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, datetime('now'), datetime('now', ?))`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(ctx, r.Timeout)
	defer cancel()

	result, err := stmt.ExecContext(ctx, t.Title, t.Content, fmt.Sprintf("+%s day", t.Expires))
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
	WHERE expires > datetime('now') AND id = ?`
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
	WHERE expires > datetime('now') ORDER BY created DESC LIMIT 10`
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
