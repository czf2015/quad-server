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

	ERROR_UNMATCHED_PASSWORD = 20006
	ERROR_INVALID_USER       = 20007
	ERROR_RESET_PASSWORD     = 20008
	ERROR_EMAIL_REGISTERED   = 20009

	ERROR_GENERATE_CAPTCHA = 20010

	ERROR_ORM_CREATE = 20011
	ERROR_ORM_UPDATE = 20012
	ERROR_ORM_GET    = 20013
	ERROR_ORM_DELETE = 20014

	ERROR_FORM_FILE = 20015
	ERROR_FILE_TYPE = 20016
	ERROR_FILE_SIZE = 20017
	ERROR_FILE_SAVE = 20018

	ERROR_NOT_FOUND = 20019
)

var MsgFlags = map[int]string{
	SUCCESS:        "Ok",
	ERROR:          "Fail",
	INVALID_PARAMS: "Invalid Params",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Unauthorized",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Timeout",
	ERROR_AUTH_TOKEN:               "Token Auth Error",
	ERROR_AUTH:                     "Unauthorized",
	ERROR_AUTH_INACTIVE:            "Inactive",

	ERROR_UNMATCHED_PASSWORD: "unmatched password",
	ERROR_INVALID_USER:       "invalid user",
	ERROR_RESET_PASSWORD:     "You have reset your password",
	ERROR_EMAIL_REGISTERED:   "email already registered",

	ERROR_GENERATE_CAPTCHA: "generate captcha fail",

	ERROR_ORM_CREATE: "Create Fail",
	ERROR_ORM_UPDATE: "Update Fail",
	ERROR_ORM_GET:    "Get Fail",
	ERROR_ORM_DELETE: "Delete Fail",

	ERROR_FORM_FILE: "form file error",
	ERROR_FILE_TYPE: "Invalid File Type",
	ERROR_FILE_SIZE: "File size exceeds the limit",
	ERROR_FILE_SAVE: "Failed to save file",

	ERROR_NOT_FOUND: "Not Found",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
