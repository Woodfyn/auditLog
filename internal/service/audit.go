package service

import (
	"context"

	audit "github.com/Woodfyn/auditLog/pkg/core"
)

type Repository interface {
	Insert(ctx context.Context, req audit.LogItem) error
}

type AuditRepo struct {
	repo Repository
}

func NewAuditRepo(repo Repository) *AuditRepo {
	return &AuditRepo{repo: repo}
}

func (s *AuditRepo) Insert(ctx context.Context, req *audit.LogRequest) error {
	item := audit.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
