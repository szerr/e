package e

// 基础异常，建议用在系统内部描述
type ErrBase struct {
	Msg string // 异常描述
}

func (self *ErrBase) Error() string {
	return self.Msg
}

func NewErr(msg string) *ErrBase {
	return &ErrBase{msg}
}

// api用异常，用于返回给其他系统
type ErrApi struct {
	Code int    // 面向api的错误码
	Msg  string // 异常描述
}

func (self *ErrApi) Error() string {
	return self.Msg
}

func NewApi(code int, msg string) *ErrApi {
	return &ErrApi{code, msg}
}

// 携带异常的异常，用于系统内部描述。可以对它使用 errors.Is 和 errors.As
type Erro struct {
	Msg string // 异常描述
	Err error  // 原始异常
}

func (self *Erro) Error() string {
	return self.Msg
}

func (self *Erro) Unwrap() error {
	return self.Err
}

func NewErro(msg string, err error) *Erro {
	return &Erro{msg, err}
}
