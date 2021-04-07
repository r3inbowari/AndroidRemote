package bilicoin

// 全局枚举量

// EXPs status
const (
	// server
	Normal = iota + 2000
	KnownError
	DatabaseError
	StructValidateError
	InternalServerError

	// user
	UserLoginSucceed
	UserExistError
	UserLoginValidateError
	UserTokenCreateError
	UserUnauthorized
	// android

	// admin

	// grpc

	// known
	NO = iota + 8086
	OK = iota + 1<<6
)
