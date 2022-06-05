package myerr

type MyErr struct {
	Code    int
	Message string
}

var (
	ErrRouterEmpty = &MyErr{
		Code:    10001,
		Message: "路由为空",
	}
)
