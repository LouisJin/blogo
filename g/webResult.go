package g

const Ok = 0
const Failure = 1001

type WebResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetDefaultWebResult() *WebResult {
	return &WebResult{
		Code: Failure,
		Msg:  "",
		Data: nil,
	}
}
