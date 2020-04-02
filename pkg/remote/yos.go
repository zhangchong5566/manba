package remote

import (
	"context"
	"fmt"
	"github.com/fagongzi/log"
	"io/ioutil"

	"code.yunzhanghu.com/be/yos"
)

// InitYosClient 初始化 yos 客户端
func InitYosClient(addr string) (err error) {
	// 初始化yos
	if addr == "" {
		addr = fmt.Sprintf("etcd:///yos_plain")
	}
	err = yos.InitClientV2ByAddress(context.Background(), addr)

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
	log.Info(ctx, "yos.UploadFileV2 start", "err", err)
	fID, err = yos.UploadFileV2(filePath, opt, nil)
	if err != nil {
		log.Info(ctx, "yos.UploadFileV2 error", "err", err)
		return
	}
	log.Info(ctx, "yos.UploadFileV2 end", "file_id", fileId)
	return
}

// DownloadFile 下载文件
func DownloadFile(ctx context.Context, fileID string) (b []byte, err error) {
	log.Info(ctx, "yos.DownloadFileV2 start", "file_id", fileID)
	r, size, err := yos.DownloadFileV2(fileID)
	if err != nil {
		log.Error(ctx, "yos.DownloadFileV2 failed", "file_id", fileID, "err", err)
		return
	}
	log.Info(ctx, "yos.DownloadFileV2 end", "file_id", fileID, "size", size)
	b, err = ioutil.ReadAll(r)
	return
}
