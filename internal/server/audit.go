package server

import (
	"context"

	"github.com/Woodfyn/auditLog/pkg/core"
)

type AuditService interface {
	Insert(ctx context.Context, msg *core.LogItem) error
}

type AuditServer struct {
	service AuditService
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (s *AuditServer) Insert(ctx context.Context, msg *core.LogItem) error {
	err := s.service.Insert(ctx, msg)

	return err
}
