package service

import (
	"fmt"
	"github.com/zhangchong5566/manba/grpcx"
	"net/http"

	"github.com/fagongzi/util/format"
	"github.com/labstack/echo"
)

const (
	apiVersion = "/v1"
)

// InitHTTPRouter init http router
func InitHTTPRouter(server *echo.Echo, ui, uiPrefix string) {
	versionGroup := server.Group(apiVersion)
	initClusterRouter(versionGroup)
	initServerRouter(versionGroup)
	initBindRouter(versionGroup)
	initRoutingRouter(versionGroup)
	initAPIRouter(versionGroup)
	initPluginRouter(versionGroup)
	initSystemRouter(versionGroup)
	initProtoSetFileRouter(versionGroup)
	initStatic(server, ui, uiPrefix)

}

type limitQuery struct {
	limit   int64
	afterID uint64
}

func idParamFactory(ctx echo.Context) (interface{}, error) {
	value := ctx.Param("id")
	if value == "" {
		return nil, fmt.Errorf("missing id path value")
	}

	id, err := format.ParseStrUInt64(value)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func limitQueryFactory(ctx echo.Context) (interface{}, error) {
	query := &limitQuery{
		limit: limit,
	}

	value := ctx.QueryParam("limit")
	if value != "" {
		l, err := format.ParseStrInt64(value)
		if err != nil {
			return nil, err
		}
		query.limit = l
	}

	value = ctx.QueryParam("after")
	if value != "" {
		l, err := format.ParseStrUInt64(value)
		if err != nil {
			return nil, err
		}
		query.afterID = l
	}

	return query, nil
}

func emptyParamFactory(ctx echo.Context) (interface{}, error) {
	return nil, nil
}

// NewFormBodyHTTPHandle returns a http handle JSON body
func NewFormBodyHTTPHandle(factory func(ctx echo.Context) interface{}, handler func(interface{}) (*grpcx.JSONResult, error)) func(echo.Context) error {
	return func(ctx echo.Context) error {
		value := factory(ctx)
		err := grpcx.ReadJSONFromBody(ctx, value)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, &grpcx.JSONResult{
				Data: err.Error(),
			})
		}

		result, err := handler(value)
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		return ctx.JSON(http.StatusOK, result)
	}
}
