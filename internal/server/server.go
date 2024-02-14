package server

import (
	"fmt"
	"net"

	audit "github.com/Woodfyn/auditLog/pkg/core"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv  *grpc.Server
	auditSrv audit.AuditServer
}

func New(auditSrv audit.AuditServer) *Server {
	return &Server{
		grpcSrv:  grpc.NewServer(),
		auditSrv: auditSrv,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	audit.RegisterAuditServer(s.grpcSrv, s.auditSrv)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
