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
	m[Ok] = "Success"
	m[UserExit] = "User has exit"
	m[UserNotExit] = "User not exit"
	m[PassWordNotRight] = "PassWord is not right"
	m[UnKnow] = "Unknow"
	return m[code]
}
