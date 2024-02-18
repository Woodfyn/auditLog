package repository

import (
	"context"

	"github.com/Woodfyn/auditLog/pkg/core"
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

func (s *AuditMongo) Insert(ctx context.Context, msg *core.LogItem) error {
	_, err := s.db.Collection("audit").InsertOne(ctx, msg)

	return err
}
