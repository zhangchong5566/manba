package remote

import (
	"context"
	"fmt"

	"code.yunzhanghu.com/be/yos"
)

// InitYosClient 初始化 yos 客户端
func InitYosClient(addr string) (err error) {
	// 初始化yos
	if addr == "" {
		addr = fmt.Sprintf("etcd:///yos_plain")
	}
	err = yos.InitClientV2ByAddressWithInsecure(context.Background(), addr)

	return
}

func UploadFile(ctx context.Context, filePath, fileId, fileName string, isEncrypt bool) (fID string, err error) {
	opt := &yos.WriterOptions{
		ACL:        yos.ACLPrivate,
		TargetName: fileName,
		IsEncrypt:  isEncrypt,
	}

	if len(fileId) > 0 {
		opt.FileID = fileId
	}

	fID, err = yos.UploadFileV2(filePath, opt, nil)
	if err != nil {
		return
	}
	return
}
