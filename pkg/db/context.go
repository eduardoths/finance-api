package db

import (
	"context"

	"gorm.io/gorm"
)

var db *gorm.DB

type db_ctx_key struct{}

var key = db_ctx_key{}

func GetFromContext(ctx context.Context) *gorm.DB {
	v := ctx.Value(key)
	if v == nil {
		return db
	}

	conn, valid := v.(*gorm.DB)
	if !valid {
		return db
	}

	return conn
}

func SetToContext(conn *gorm.DB, ctx context.Context) context.Context {
	return context.WithValue(ctx, key, conn)
}
