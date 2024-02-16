package service

import (
	"context"

	"github.com/Woodfyn/auditLog/pkg/core"
	audit "github.com/Woodfyn/auditLog/pkg/proto"
)

type Repository interface {
	Insert(ctx context.Context, req core.LogItem) (*audit.Empty, error)
}

type Audit struct {
	repo Repository
}

func NewAuditRepo(repo Repository) *Audit {
	return &Audit{repo: repo}
}

func (s *Audit) Insert(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	item := core.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
