package service

import (
	"github.com/fagongzi/log"
	"github.com/labstack/echo"
	"github.com/zhangchong5566/manba/grpcx"
)

type backup struct {
	ToAddr string `json:"toAddr"`
}

func initSystemRouter(server *echo.Group) {
	server.GET("/system",
		grpcx.NewGetHTTPHandle(emptyParamFactory, getSystemHandler))

	server.POST("/system/backup",
		grpcx.NewJSONBodyHTTPHandle(backupFactory, postBackupHandler))
}

func getSystemHandler(value interface{}) (*grpcx.JSONResult, error) {
	info, err := Store.System()
	if err != nil {
		log.Errorf("api-system-get: errors:%+v", err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{Data: info}, nil
}

func postBackupHandler(value interface{}) (*grpcx.JSONResult, error) {
	err := Store.BackupTo(value.(*backup).ToAddr)
	if err != nil {
		log.Errorf("api-system-backup: errors:%+v", err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{}, nil
}

func backupFactory() interface{} {
	return &backup{}
}
