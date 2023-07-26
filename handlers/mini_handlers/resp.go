package minihandlers

type Resp struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp(w Wrong, data interface{}) Resp {
	if data == nil {
		data = struct{}{}
	}
	return Resp{
		Code: w.Code(),
		Msg:  w.Error(),
		Data: data,
	}
}
