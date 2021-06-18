// Package base 提供自定义错误的方法
package base

// 错误号code信息
const (
	// 通用错误号，错误号范围 [0, 10000)
	ErrCodeSuccess            = 0  // 请求成功
	ErrCodeUnknownError       = 1  // 请求成功
	ErrCodeParamsError        = 2  // 参数错误
	ErrCodeGetConfError       = 3  // 获取配置失败
	ErrCodeJSONMarshalError   = 4  // json解析失败
	ErrCodeJSONUnMarshalError = 5  // json解析失败
	ErrCodeEncodeError        = 6  // 编码错误
	ErrCodeDecodeError        = 7  // 编码错误
	ErrCodeCallRedisError     = 8  // 调用 redis 失败
	ErrCodeCallMysqlError     = 9  // 调用 mysql 失败
	ErrCodeCallServiceError   = 10 // 服务错误
	ErrCodeSqlError           = 11 // sql错误

	// 业务错误号，错误号范围 [10000, 999999)；前两位表示属于哪个业务，后四位表示该业务的细分错误号
	// todo...
)

// 错误号msg信息
const (
	ErrMsgSuccess            = "success"
	ErrMsgUnknownError       = "Unknown err"
	ErrMsgParamsError        = "params err"
	ErrMsgGetConfError       = "get conf err"
	ErrMsgJSONMarshalError   = "json marshal err"
	ErrMsgJSONUnMarshalError = "json unmarshal err"
	ErrMsgEncodeError        = "encode err"
	ErrMsgDecodeError        = "decode err"
	ErrMsgCallRedisError     = "call redis err"
	ErrMsgCallMysqlError     = "call mysql err"
	ErrMsgCallServiceError   = "call service err"
	ErrMsgSqlError           = "sql err"
)

// 错误号和错误信息映射关系
var (
	ErrSuccess            = NewErrno(ErrCodeSuccess, ErrMsgSuccess)
	ErrUnknownError       = NewErrno(ErrCodeUnknownError, ErrMsgUnknownError)
	ErrParamsError        = NewErrno(ErrCodeParamsError, ErrMsgParamsError)
	ErrGetConfError       = NewErrno(ErrCodeGetConfError, ErrMsgGetConfError)
	ErrJSONMarshalError   = NewErrno(ErrCodeJSONMarshalError, ErrMsgJSONMarshalError)
	ErrJSONUnMarshalError = NewErrno(ErrCodeJSONUnMarshalError, ErrMsgJSONUnMarshalError)
	ErrEncodeError        = NewErrno(ErrCodeEncodeError, ErrMsgEncodeError)
	ErrDecodeError        = NewErrno(ErrCodeDecodeError, ErrMsgDecodeError)
	ErrCallRedisError     = NewErrno(ErrCodeCallRedisError, ErrMsgCallRedisError)
	ErrCallMysqlError     = NewErrno(ErrCodeCallMysqlError, ErrMsgCallMysqlError)
	ErrCallServiceError   = NewErrno(ErrCodeCallServiceError, ErrMsgCallServiceError)
	ErrSqlError           = NewErrno(ErrCodeSqlError, ErrMsgSqlError)
)

// ErrCode 错误号类型定义
type ErrCode uint32

// Errno 自定义错误号结构
type myErrno struct {
	code ErrCode
	msg  string
}

// Error 自定义错误号结构
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

// Errno 获取错误号方法
func (err *myErrno) Errno() ErrCode {
	return err.code
}
