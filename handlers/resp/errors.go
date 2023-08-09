package resp

type Wrong string

func (e Wrong) Error() string {
	return string(e)
}

func (e Wrong) Code() uint32 {
	return wrongCodeMap[e]
}

var wrongCodeMap = map[Wrong]uint32{
	ErrAuthFailed: 40001,
	Success:       20000,
}

const (
	ErrAuthFailed = "user auth failed"
	Success       = "success"
)
