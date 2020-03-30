package service

import (
	"github.com/zhangchong5566/manba/pkg/pb/rpcpb"
	"github.com/zhangchong5566/manba/pkg/store"
)

var (
	// MetaService global service
	MetaService rpcpb.MetaServiceServer
	// Store global store db
	Store store.Store
)

// Init init service package
func Init(db store.Store) {
	Store = db
	MetaService = newMetaService(db)
}
