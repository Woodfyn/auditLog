package server

import (
	"context"

	audit "github.com/Woodfyn/auditLog/pkg/proto"
)

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error)
}

type AuditServer struct {
	service AuditService
	audit.UnimplementedAuditServer
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service:                  service,
		UnimplementedAuditServer: audit.UnimplementedAuditServer{},
	}
}

func (s *AuditServer) Log(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	empty, err := s.service.Insert(ctx, req)

	return empty, err
}
