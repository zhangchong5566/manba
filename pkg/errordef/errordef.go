package errordef

import (
	"fmt"
)

const (
	ErrOK int32 = 0

	ErrMissingParameter    = 1000
	ErrSignExpired         = 1001
	ErrWrongParameter      = 1002
	ErrSign                = 1003
	Err3Des                = 1004
	ErrDataCountOverLimit  = 1013

	ErrInner          = 6000
	ErrOverRateLimit  = 60001 // 超过频率限制
	ErrSystemMaintain = 60002 // 系统维护中，稍后尝试

	// C端应用相关
	ErrAppIDIsEmpty      = 80001 // 请求头部应用ID错误
	ErrAppTypeIsEmpty    = 80002 // 请求头部应用类型错误
	ErrDeviceIDIsEmpty   = 80003 // 请求头部设备ID错误
	ErrDeviceTypeIsEmpty = 80004 // 请求头部设备类型错误
	ErrMemberIDIsEmpty   = 80005 // 请求头部会员ID错误
	ErrOsIsEmpty         = 80006 // 请求头部Os错误
	ErrOsVersionIsEmpty  = 80007 // 请求头部Os版本错误
)

var errorCodeToMessageMap = map[int32]string{
	ErrOK:                 "OK",
	ErrWrongParameter:     "请求参数格式不正确",
	ErrMissingParameter:   "缺少请求参数 %s",
	ErrSignExpired:        "签名已过期",
	ErrSign:               "签名错误",
	Err3Des:               "加密错误",
	ErrDataCountOverLimit: "请求数量超限",

	ErrInner:             "内部错误",
	ErrOverRateLimit:     "调用频次过高，请稍后尝试",
	ErrSystemMaintain:    "系统维护中",

	ErrAppIDIsEmpty:      "请求头部应用ID错误",
	ErrAppTypeIsEmpty:    "请求头部应用类型错误",
	ErrDeviceIDIsEmpty:   "请求头部设备ID错误",
	ErrDeviceTypeIsEmpty: "请求头部设备类型错误",
	ErrMemberIDIsEmpty:   "请求头部会员ID错误",
	ErrOsIsEmpty:         "请求头部Os错误",
	ErrOsVersionIsEmpty:  "请求头部Os版本错误",
}

// ErrSuccessOk用来判断从前置流程、规则引擎、流程引擎和组织机构的响应是否成功
const ErrSuccessOk = "0000"

func GetErrorMessage(errorCode int32) (errorMessage string) {
	var ok bool
	if errorMessage, ok = errorCodeToMessageMap[errorCode]; !ok {
		panic("unknown errors")
	}

	return
}

type Error struct {
	Code     int32
	Message  string
	RawError error
}

func (e *Error) Error() string {
	if e.RawError == nil {
		return fmt.Sprintf("%d-%s", e.Code, e.Message)
	}
	return fmt.Sprintf("%d-%s-%s", e.Code, e.Message, e.RawError.Error())
}
