package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_AUTH_INACTIVE            = 20005
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "Invalid Params",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Unauthorized",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Timeout",
	ERROR_AUTH_TOKEN:               "Token Auth Error",
	ERROR_AUTH:                     "Unauthorized",
	ERROR_AUTH_INACTIVE:            "Inactive",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
