package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Cmo cluster manage operation，不能用于业务操作
type Cmo struct {
	ClientName ClientName
	client     *mongo.Client
}

// NewCmo 创建新的集群管理操作对象
func NewCmo(clientName ClientName) *Cmo {
	return &Cmo{
		ClientName: clientName,
		client:     Client(clientName),
	}
}

// ListDatabaseNames 列出集群中所有数据库名
func (cmo *Cmo) ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error) {
	return cmo.client.ListDatabaseNames(ctx, filter, opts...)
}
