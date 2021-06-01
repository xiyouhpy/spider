// Package base 提供自定义错误的方法
package base

type ErrCode int32

// MyError 自定义错误
type MyError struct {
	errCode ErrCode
	errMsg  string
}

const (
	// 通用错误号，错误号范围 [0, 10000)
	ErrSuccess          = 0 // 请求成功
	ErrErrnoUnregister  = 1 // 未注册错误号
	ErrParamsError      = 2 // 参数错误
	ErrJSONMarshalError = 3 // json解析失败
	ErrEncodeError      = 4 // 编码错误
	ErrCallRedisError   = 5 // 调用 redis 失败
	ErrServiceError     = 6 // 服务错误
	ErrGetConfError     = 7 // 获取配置失败

	// 业务错误号，错误号范围 [10000, 999999)；前两位表示属于哪个业务，后四位表示该业务的细分错误号
	// todo...
)

// errMap 错误号和错误描述map关系
var errMap = map[ErrCode]string{
	// 预留前 1000 作为系统错误号
	0: "success",
	1: "errno unregister",
	2: "params error",
	3: "json marshal error",
	4: "encode error",
	5: "call redis error",
	6: "service error",
	7: "get conf error",
}

// NewError 自定义错误构造方法
func NewError(errCode ErrCode, errMsg string) *MyError {
	myErrCode := &MyError{
		errCode: errCode,
		errMsg:  errMsg,
	}

	return myErrCode
}

// MyErrMsg 获取错误号的 msg
func (m *MyError) MyErrMsg() string {
	if errMap[m.errCode] == "" {
		return errMap[ErrErrnoUnregister]
	}

	return m.errMsg
}

// MyErrCode 获取错误号的 code
func (m *MyError) MyErrCode() ErrCode {
	if errMap[m.errCode] == "" {
		return ErrErrnoUnregister
	}

	return m.errCode
}
