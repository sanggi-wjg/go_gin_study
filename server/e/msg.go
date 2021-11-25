package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token 실패",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token 만료",
	ERROR_AUTH_TOKEN:               "Token 생성오류",
	ERROR_AUTH:                     "Token 오류",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
