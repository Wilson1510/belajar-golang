package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (w WebResponse) GetCode() int {
	return w.Code
}