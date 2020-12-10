package Result

const (
	Ok               = iota
	UserNotExit      = iota
	UserExit         = iota
	PassWordNotRight = iota
	UnKnow           = iota
)

func code2Msg(code int) string {
	m := make(map[int]string)
	m[Ok] = "success"
	m[UserExit] = "user has exit"
	m[UserNotExit] = "user not exit"
	m[PassWordNotRight] = "passWord is not right"
	m[UnKnow] = "unknown"
	return m[code]
}
