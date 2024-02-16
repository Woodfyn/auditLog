package repository

import (
	"context"

	"github.com/Woodfyn/auditLog/pkg/core"
	audit "github.com/Woodfyn/auditLog/pkg/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuditMongo struct {
	db *mongo.Database
}

func NewAuditMongo(db *mongo.Database) *AuditMongo {
	return &AuditMongo{
		db: db,
	}
}

func (s *AuditMongo) Insert(ctx context.Context, item core.LogItem) (*audit.Empty, error) {
	_, err := s.db.Collection("audit").InsertOne(ctx, item)

	return &audit.Empty{}, err
}
