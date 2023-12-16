package usecase

import (
	"context"

	"github.com/1kovalevskiy/snippetbox/internal/entity"
)

type (
	Snippet interface {
		CreateSnippet(ctx context.Context, t entity.SnippetCreate) (int, error)
		GetSnippet(ctx context.Context, id int) (*entity.Snippet, error)
		GetTenLatestSnippet(ctx context.Context) ([]*entity.Snippet, error)
	}

	SnippetRepo interface {
		Insert(ctx context.Context, t entity.SnippetCreate) (int, error)
		Get(ctx context.Context, id int) (*entity.Snippet, error)
		Latest(ctx context.Context) ([]*entity.Snippet, error)
	}
)

type SnippetUseCase struct {
	repo SnippetRepo
}

func New(r SnippetRepo) *SnippetUseCase {
	return &SnippetUseCase{r}
}

func (uc *SnippetUseCase) CreateSnippet(ctx context.Context, t entity.SnippetCreate) (int, error) {
	id, err := uc.repo.Insert(ctx, t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *SnippetUseCase) GetSnippet(ctx context.Context, id int) (*entity.Snippet, error) {
	snippet, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return snippet, nil
}

func (uc *SnippetUseCase) GetTenLatestSnippet(ctx context.Context) ([]*entity.Snippet, error) {
	snippets, err := uc.repo.Latest(ctx)
	if err != nil {
		return nil, err
	}
	return snippets, nil
}
