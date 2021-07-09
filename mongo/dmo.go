package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Dmo database manage operation，不能用于业务操作
type Dmo struct {
	ClientName ClientName
	DBName     string
	database   *mongo.Database
}

// NewDmo 创建新的数据库操作对象
func NewDmo(clientName ClientName, dbName string) *Dmo {
	return &Dmo{
		ClientName: clientName,
		DBName:     dbName,
		database:   Client(clientName).Database(dbName),
	}
}

// ListCollectionNames 列出集群中所有数据库名
func (d *Dmo) ListCollectionNames(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]string, error) {
	return Client(d.ClientName).Database(d.DBName).ListCollectionNames(ctx, filter, opts...)
}
