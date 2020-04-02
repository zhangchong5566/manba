package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/fagongzi/log"
	"github.com/labstack/echo"
	"github.com/zhangchong5566/manba/grpcx"
	"github.com/zhangchong5566/manba/pkg/pb/metapb"
	"github.com/zhangchong5566/manba/pkg/remote"
	"github.com/zhangchong5566/manba/pkg/util"
	"io"
	"os"
	"strings"
	"time"
)

func initProtoSetFileRouter(server *echo.Group) {
	server.GET("/protosets/:id",
		grpcx.NewGetHTTPHandle(idParamFactory, getProtoSetFileHandler))
	server.DELETE("/protosets/:id",
		grpcx.NewGetHTTPHandle(idParamFactory, deleteProtoSetFileHandler))
	server.PUT("/protosets",
		NewFormBodyHTTPHandle(putProtoSetFileFactory, postProtoSetFileHandler))
	server.GET("/protosets",
		grpcx.NewGetHTTPHandle(limitQueryFactory, listProtoSetFileHandler))
}

func postProtoSetFileHandler(value interface{}) (*grpcx.JSONResult, error) {
	id, err := Store.PutProtoSetFile(value.(*metapb.ProtoSetFile))
	if err != nil {
		log.Errorf("api-protoset-put: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{Data: id}, nil
}

func deleteProtoSetFileHandler(value interface{}) (*grpcx.JSONResult, error) {
	err := Store.RemoveProtoSetFile(value.(uint64))
	if err != nil {
		log.Errorf("api-protoset-delete: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{}, nil
}

func getProtoSetFileHandler(value interface{}) (*grpcx.JSONResult, error) {
	value, err := Store.GetProtoSetFile(value.(uint64))
	if err != nil {
		log.Errorf("api-protoset-get: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{Data: value}, nil
}

func listProtoSetFileHandler(value interface{}) (*grpcx.JSONResult, error) {
	query := value.(*limitQuery)
	var values []*metapb.ProtoSetFile

	err := Store.GetProtoSetFiles(limit, func(data interface{}) error {
		v := data.(*metapb.ProtoSetFile)
		if int64(len(values)) < query.limit && v.ID > query.afterID {
			values = append(values, v)
		}
		return nil
	})
	if err != nil {
		log.Errorf("api-protoset-list-get: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{Data: values}, nil
}

func putProtoSetFileFactory(c echo.Context) interface{} {

	name := c.FormValue("name")
	version := c.FormValue("version")

	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("上传文件时，获取文件错误 error, errors:%+v", err)
		return err
	}

	src, err := file.Open()
	if err != nil {
		log.Errorf("打开文件错误 error, errors:%+v", err)
		return err
	}
	defer src.Close()

	tempCatalog := "/tmp/protoset/"
	if strings.LastIndex(tempCatalog, "/") != len(tempCatalog)-1 {
		tempCatalog = tempCatalog + "/"
	}
	// 如果目录不存在，先创建目录
	if !util.FileIsExist(tempCatalog) {
		err := os.MkdirAll(tempCatalog, os.ModePerm)
		if err != nil {
			log.Errorf("putProtoSetFileFactory os.MkdirAll error, errors:%+v", err)
			return nil
		}
	}
	fileId := time.Now().UnixNano()
	filePath := tempCatalog + fmt.Sprintf("%x", fileId) + ".protoset"

	dst, err := os.Create(filePath)
	if err != nil {
		log.Errorf("putProtoSetFileFactory os.Create error, errors:%+v", err)
		return nil
	}
	defer dst.Close()

	// 文件保存到本地
	if _, err = io.Copy(dst, src); err != nil {
		log.Errorf("putProtoSetFileFactory io.Copy error, errors:%+v", err)
		return nil
	}

	// 上传到yos
	yosFileId, err := remote.UploadFile(context.Background(), filePath, "", file.Filename, false)
	if err != nil {
		log.Errorf("putProtoSetFileFactory  remote.UploadFile error, err=%v, fileName=%v", err, file.Filename)
		return nil
	}

	// 计算文件Md5值
	md5 := md5.New()
	io.Copy(md5, src)
	MD5Str := hex.EncodeToString(md5.Sum(nil))

	protoSetFile := &metapb.ProtoSetFile{
		Name:     name,
		Version:  version,
		FileName: file.Filename,
		FileId:   yosFileId,
		FileMd5:  MD5Str,
		CreateAt: util.NowWithMillisecond(),
	}

	return protoSetFile
}
