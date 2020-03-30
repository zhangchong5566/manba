package service

import (
	"github.com/fagongzi/log"
	"github.com/labstack/echo"
	"github.com/zhangchong5566/manba/grpcx"
	"github.com/zhangchong5566/manba/pkg/pb/metapb"
)

func initPluginRouter(server *echo.Group) {
	server.GET("/plugins/:id",
		grpcx.NewGetHTTPHandle(idParamFactory, getPluginHandler))
	server.DELETE("/plugins/:id",
		grpcx.NewGetHTTPHandle(idParamFactory, deletePluginHandler))
	server.PUT("/plugins",
		grpcx.NewJSONBodyHTTPHandle(putPluginFactory, postPluginHandler))
	server.GET("/plugins",
		grpcx.NewGetHTTPHandle(limitQueryFactory, listPluginHandler))
	server.PUT("/plugins/apply",
		grpcx.NewJSONBodyHTTPHandle(putPluginAppliedFactory, putPluginAppliedHandler))
	server.GET("/plugins/apply",
		grpcx.NewGetHTTPHandle(emptyParamFactory, getPluginAppliedHandler))
}

func getPluginAppliedHandler(value interface{}) (*grpcx.JSONResult, error) {
	value, err := Store.GetAppliedPlugins()
	if err != nil {
		log.Errorf("api-plugin-get-applied: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{Data: value}, nil
}

func putPluginAppliedHandler(value interface{}) (*grpcx.JSONResult, error) {
	err := Store.ApplyPlugins(value.(*metapb.AppliedPlugins))
	if err != nil {
		log.Errorf("api-plugin-put-applied: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{}, nil
}

func postPluginHandler(value interface{}) (*grpcx.JSONResult, error) {
	id, err := Store.PutPlugin(value.(*metapb.Plugin))
	if err != nil {
		log.Errorf("api-plugin-put: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{Data: id}, nil
}

func deletePluginHandler(value interface{}) (*grpcx.JSONResult, error) {
	err := Store.RemovePlugin(value.(uint64))
	if err != nil {
		log.Errorf("api-plugin-delete: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{}, nil
}

func getPluginHandler(value interface{}) (*grpcx.JSONResult, error) {
	value, err := Store.GetPlugin(value.(uint64))
	if err != nil {
		log.Errorf("api-plugin-get: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{Data: value}, nil
}

func listPluginHandler(value interface{}) (*grpcx.JSONResult, error) {
	query := value.(*limitQuery)
	var values []*metapb.Plugin

	err := Store.GetPlugins(limit, func(data interface{}) error {
		v := data.(*metapb.Plugin)
		if int64(len(values)) < query.limit && v.ID > query.afterID {
			values = append(values, v)
		}
		return nil
	})
	if err != nil {
		log.Errorf("api-plugin-list-get: req %+v, errors:%+v", value, err)
		return &grpcx.JSONResult{Code: -1, Data: err.Error()}, nil
	}

	return &grpcx.JSONResult{Data: values}, nil
}

func putPluginFactory() interface{} {
	return &metapb.Plugin{}
}

func putPluginAppliedFactory() interface{} {
	return &metapb.AppliedPlugins{}
}
