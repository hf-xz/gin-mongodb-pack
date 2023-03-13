package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// ID mongodb主键
type ID struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

// SoftDelete 软删除
// 参考 https://arielorozco.com/how-to-soft-delete-with-mongodb-and-go
type SoftDelete struct {
	DeletedAt time.Time `json:"-" bson:"deleted_at,omitempty"`
}
