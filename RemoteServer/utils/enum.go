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

const (
	// hardware
	DEF = iota
	REG
	HRB
	NOR
	END
	ASK
	SNY
	REQ
	OPR
)

const (
	// hardware
	REQ_START_SENDER = iota + 1000
	REQ_PAUSE_SENDER
)
