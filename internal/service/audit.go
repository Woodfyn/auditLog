package service

import (
	"context"

	"github.com/Woodfyn/auditLog/pkg/core"
)

type Repository interface {
	Insert(ctx context.Context, msg *core.LogItem) error
}

type Audit struct {
	repo Repository
}

func NewAuditRepo(repo Repository) *Audit {
	return &Audit{repo: repo}
}

func (s *Audit) Insert(ctx context.Context, msg *core.LogItem) error {
	return s.repo.Insert(ctx, msg)
}
