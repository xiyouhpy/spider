// Package base 提供自定义错误的方法
package base

// 错误号code信息
const (
	// 通用错误号，错误号范围 [0, 10000)
	ErrCodeSuccess          = 0 // 请求成功
	ErrCodeUnknownError     = 1 // 请求成功
	ErrCodeParamsError      = 2 // 参数错误
	ErrCodeJSONMarshalError = 3 // json解析失败
	ErrCodeEncodeError      = 4 // 编码错误
	ErrCodeCallRedisError   = 5 // 调用 redis 失败
	ErrCodeServiceError     = 6 // 服务错误
	ErrCodeGetConfError     = 7 // 获取配置失败

	// 业务错误号，错误号范围 [10000, 999999)；前两位表示属于哪个业务，后四位表示该业务的细分错误号
	// todo...
)

// 错误号msg信息
const (
	ErrMsgSuccess          = "success"
	ErrMsgUnknownError     = "Unknown err"
	ErrMsgParamsError      = "params err"
	ErrMsgJSONMarshalError = "json marshal err"
	ErrMsgEncodeError      = "encode err"
	ErrMsgCallRedisError   = "call redis err"
	ErrMsgServiceError     = "service err"
	ErrMsgGetConfError     = "get conf err"
)

// 错误号和错误信息映射关系
var (
	ErrSuccess          = NewErrno(ErrCodeSuccess, ErrMsgSuccess)
	ErrUnknownError     = NewErrno(ErrCodeUnknownError, ErrMsgUnknownError)
	ErrParamsError      = NewErrno(ErrCodeParamsError, ErrMsgParamsError)
	ErrJSONMarshalError = NewErrno(ErrCodeJSONMarshalError, ErrMsgJSONMarshalError)
	ErrEncodeError      = NewErrno(ErrCodeEncodeError, ErrMsgEncodeError)
	ErrCallRedisError   = NewErrno(ErrCodeCallRedisError, ErrMsgCallRedisError)
	ErrServiceError     = NewErrno(ErrCodeServiceError, ErrMsgServiceError)
	ErrGetConfError     = NewErrno(ErrCodeGetConfError, ErrMsgGetConfError)
)

// 错误号类型定义
type ErrCode uint32

// Errno 自定义错误号结构
type myErrno struct {
	code ErrCode
	msg  string
}

// Errno 自定义错误号结构
type Error interface {
	error
	Errno() ErrCode
}

// NewErrno 错误号对象
func NewErrno(code ErrCode, msg string) Error {
	return &myErrno{code, msg}
}

// ErrMsg 获取错误描述方法
func (err *myErrno) Error() string {
	return err.msg
}

// ErrNo 获取错误号方法
func (err *myErrno) Errno() ErrCode {
	return err.code
}
